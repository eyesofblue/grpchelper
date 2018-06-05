package logic

import (
	"fmt"
	"github.com/eyesofblue/grpchelper/comm"
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

func AddHandlerFile() {
	pbDir := comm.GetPbDir(".")
	if !comm.PathExist(pbDir) {
		panic("pb Dir Not Found")
	}

	pbFile := comm.GetPbFilePath(pbDir)
	if !comm.PathExist(pbFile) {
		tmpErr := fmt.Sprintf("%s File Not Found", pbFile)
		panic(tmpErr)
	}
}

func AddSvrMainFile() {

}

func Add(interfaceName string) {
	rpcName := comm.CapitalizeStr(interfaceName)
	AddProtoFile(rpcName)
}
