package logic

import (
	// "bufio"
	"fmt"
	"github.com/eyesofblue/grpchelper/comm"
	// "os"
)

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

	// add msg define seg
	msgTargetLine := comm.GetTagSegEnd4PbMsg()
	msgContent := comm.GetContentTmpl4PbMsg(rpcName)
	comm.Insert2File(pbFile, msgContent, msgTargetLine, true)

	// add service define seg
	serviceTargetLine := comm.GetTagSegEnd4PbService()
	serviceContent := comm.GetContentTmpl4PbService(rpcName)
	comm.Insert2File(pbFile, serviceContent, serviceTargetLine, true)
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

	// add handler impl
	handlerImplTargetLine := comm.GetTagSegEnd4HandlerImpl()
	handlerImplContent := fmt.Sprintf(comm.CONTENT_TMPL_HANDLER_IMPL, rpcName, comm.GetRpcReqName(rpcName), comm.GetRpcRspName(rpcName))
	comm.Insert2File(handlerFile, handlerImplContent, handlerImplTargetLine, true)
}

func AddSvrMainFile() {

}

func Add(interfaceName string) {
	rpcName := comm.CapitalizeStr(interfaceName)
	AddProtoFile(rpcName)
	AddHandlerFile(rpcName)
}
