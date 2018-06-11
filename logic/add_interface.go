package logic

import (
	"bufio"
	"fmt"
	"github.com/eyesofblue/grpchelper/comm"
	"io"
	"os"
	"regexp"
)

func IsRpcExist(rpcName string) bool {
	isExist := false

	pbDir := comm.GetPbDir(".")
	if !comm.PathExist(pbDir) {
		panic("pb Dir Not Found")
	}

	pbFile := comm.GetPbFilePath(pbDir)
	if !comm.PathExist(pbFile) {
		tmpErr := fmt.Sprintf("%s File Not Found", pbFile)
		panic(tmpErr)
	}

	pattern := fmt.Sprintf(comm.SERVICE_EXIST_FLAG, rpcName, comm.GetRpcReqName(rpcName), comm.GetRpcRspName(rpcName))

	f, err := os.Open(pbFile)
	if err != nil {
		panic(err)
	}

	fileReader := bufio.NewReader(f)
	for {
		line, err := fileReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		if regexp.MustCompile(pattern).MatchString(line) {
			isExist = true
			break
		}
	}

	return isExist
}

func AddProtoFile(rpcName string) {
	pbDir := comm.GetPbDir(".")
	if !comm.PathExist(pbDir) {
		panic("pb Dir Not Found")
	}

	pbFile := comm.GetPbFilePath(pbDir)
	if !comm.PathExist(pbFile) {
		tmpErr := fmt.Sprintf("%s File Not Found", pbFile)
		panic(tmpErr)
	}

	insertList := make([]*comm.InsertItem, 0)

	// add msg define seg
	insertList = append(insertList, &comm.InsertItem{TargetLine: comm.GetTagSegEnd4PbMsg(), Content: comm.GetContentTmpl4PbMsg(rpcName), InsertBefore: true})
	// add service define seg
	insertList = append(insertList, &comm.InsertItem{TargetLine: comm.GetTagSegEnd4PbService(), Content: comm.GetContentTmpl4PbService(rpcName), InsertBefore: true})

	comm.Insert2File(pbFile, insertList)
}

/*
func CreateHandlerFile(handlerFile string) {
	mainDirFromGoSrcPath, err := comm.GetPrefixFromGoSrcPath()
	if err != nil {
		panic(err)
	}

	pbDirFromGoSrcPath := comm.GetPbDir(mainDirFromGoSrcPath)
	content := fmt.Sprintf(comm.CONTENT_TMPL_HANDLER_HEADER, pbDirFromGoSrcPath)
	content += "\n" + fmt.Sprintf("%s\n\n%s\n", comm.GetTagSegBegin4HandlerImpl(), comm.GetTagSegEnd4HandlerImpl())

	file, err := os.OpenFile(handlerFile, os.O_EXCL|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	fileWriter := bufio.NewWriter(file)
	fileWriter.WriteString(content)
	fileWriter.Flush()

	file.Close()
}
*/

func AddHandlerFile(rpcName string) {
	handlerDir := comm.GetHandlerDir(".")
	if !comm.PathExist(handlerDir) {
		panic("svr/handler Dir Not Found")
	}

	handlerFile := comm.GetHandlerFilePath(handlerDir)
	if !comm.PathExist(handlerFile) {
		tmpErr := fmt.Sprintf("%s File Not Found", handlerFile)
		panic(tmpErr)
	}

	insertList := make([]*comm.InsertItem, 0)

	// add handler impl
	insertList = append(insertList, &comm.InsertItem{TargetLine: comm.GetTagSegEnd4HandlerImpl(), Content: fmt.Sprintf(comm.CONTENT_TMPL_HANDLER_IMPL, rpcName, comm.GetRpcReqName(rpcName), comm.GetRpcRspName(rpcName)), InsertBefore: true})

	comm.Insert2File(handlerFile, insertList)
}

func AddStubFile(rpcName string) {
	stubDir := comm.GetStubDir(".")
	if !comm.PathExist(stubDir) {
		panic("cli_tool/stub Dir Not Found")
	}

	stubFile := comm.GetStubFilePath(stubDir)
	if !comm.PathExist(stubFile) {
		tmpErr := fmt.Sprintf("%s File Not Found", stubFile)
		panic(tmpErr)
	}

	rpcReqName := comm.GetRpcReqName(rpcName)

	insertList := make([]*comm.InsertItem, 0)

	// add newreq seg
	insertList = append(insertList, &comm.InsertItem{TargetLine: comm.GetTagSegEnd4StubNewReq(), Content: fmt.Sprintf(comm.CONTENT_TMPL_STUB_NEWREQ, rpcReqName, rpcReqName), InsertBefore: true})
	// add register seg
	insertList = append(insertList, &comm.InsertItem{TargetLine: comm.GetTagSegEnd4StubRegister(), Content: fmt.Sprintf(comm.CONTENT_TMPL_STUB_REGISTER, rpcName, rpcReqName), InsertBefore: true})

	comm.Insert2File(stubFile, insertList)
}

func Add(interfaceName string) {
	rpcName := comm.CapitalizeStr(interfaceName)
	if IsRpcExist(rpcName) {
		tmpErr := fmt.Sprintf("rpc %s exists", rpcName)
		panic(tmpErr)
	}

	/*
		AddProtoFile(rpcName)
		AddHandlerFile(rpcName)
		AddStubFile(rpcName)
	*/
}
