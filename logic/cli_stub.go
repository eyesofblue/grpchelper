package logic

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/eyesofblue/grpchelper/comm"
	"golang.org/x/net/context"
	"os"
	"reflect"
	"time"
)

// rpcRequest注册生成函数
var reqNewMap = make(map[string]func([]byte) interface{})

func RegisterReqNew(methodName string, implFunc func([]byte) interface{}) {
	reqNewMap[methodName] = implFunc
}

func getAllMethods(st interface{}) []string {
	funcList := make([]string, 0)

	t := reflect.TypeOf(st)
	for i := 0; i != t.NumMethod(); i++ {
		funcList = append(funcList, t.Method(i).Name)
	}

	return funcList
}

func ShowFuncList(st interface{}) string {
	funcList := getAllMethods(st)
	var ret string

	ret += os.Args[0] + ":\n"
	for _, funcName := range funcList {
		ret += "\t-f " + funcName + " -d json_req\n"
	}

	return ret
}

func CallByMethodName(st interface{}, methodName string, strJsonReq string) ([]reflect.Value, error) {
	v := reflect.ValueOf(st).MethodByName(methodName)

	if !v.IsValid() {
		tmpMsg := fmt.Sprintf("[Method Not Found]:%s", methodName)
		return nil, errors.New(tmpMsg)
	}

	// 生成req
	reqImplFunc, isExist := reqNewMap[methodName]
	if !isExist {
		tmpMsg := fmt.Sprintf("[Request Not Registered]:%s", comm.GetRpcReqName(methodName))
		return nil, errors.New(tmpMsg)
	}

	req := reqImplFunc([]byte(strJsonReq))
	if req == nil {
		tmpMsg := fmt.Sprintf("[JsonUnmarshal Failed]:%s", strJsonReq)
		return nil, errors.New(tmpMsg)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 所有rpc接口一共只有两个入参 ctx和req
	params := make([]reflect.Value, 2)
	params[0] = reflect.ValueOf(ctx)
	params[1] = reflect.ValueOf(req)

	return v.Call(params), nil
}

func ClientStub(st interface{}) {
	var methodName string
	flag.StringVar(&methodName, "f", "", "methodName")

	var jsonReq string
	flag.StringVar(&jsonReq, "d", "", "jsonReq")

	flag.Parse()

	retList, err := CallByMethodName(st, methodName, jsonReq)

	if err != nil {
		fmt.Println(err)
		fmt.Println(ShowFuncList(st))
		return
	}

	// 所有rpc接口一共只有两个返回值 rsp和error
	if retList[1].IsNil() {
		rsp, err := json.MarshalIndent(retList[0].Interface(), "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", string(rsp))
	} else {
		err = retList[1].Interface().(error)
		fmt.Printf("%+v\n", err)
	}
}
