package logic

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/eyesofblue/grpchelper/comm"
	"os"
	"reflect"
)

// rpcRequest注册生成函数
var reqNewMap = make(map[string]func() interface{})

func RegisterReqNew(methodName string, implFunc func() interface{}) {
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

func CallByMethodName(ctx context.Context, st interface{}, methodName string, strJsonReq string) ([]reflect.Value, error) {
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

	req := reqImplFunc()
	if len(strJsonReq) > 0 {
		err := json.Unmarshal([]byte(strJsonReq), req)
		if err != nil {
			panic(errors.New("json req invalid:" + err.Error()))
		}
	}

	// todo 目前rpc接口一共只有两个入参 ctx和req  gRpc的一些CallOption参数后续要带进来
	params := make([]reflect.Value, 2)
	params[0] = reflect.ValueOf(ctx)
	params[1] = reflect.ValueOf(req)

	return v.Call(params), nil
}

func ClientStub(ctx context.Context, st interface{}) {
	var methodName string
	flag.StringVar(&methodName, "f", "", "methodName")

	var jsonReq string
	flag.StringVar(&jsonReq, "d", "", "jsonReq")

	flag.Parse()

	retList, err := CallByMethodName(ctx, st, methodName, jsonReq)

	if err != nil {
		fmt.Println(err)
		fmt.Println(ShowFuncList(st))
		return
	}

	// 所有rpc接口一共只有两个返回值 rsp和error
	if retList[1].IsNil() {
		var rsp []byte
		if retList[0].IsNil() {
			rsp, err = json.MarshalIndent("{}", "", "    ")
		} else {
			rsp, err = json.MarshalIndent(retList[0].Interface(), "", "    ")
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", string(rsp))
	} else {
		err = retList[1].Interface().(error)
		fmt.Printf("%+v\n", err)
	}
}
