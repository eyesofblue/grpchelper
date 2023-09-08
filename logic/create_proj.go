package logic

import (
	"github.com/eyesofblue/grpchelper/comm"
	"github.com/eyesofblue/grpchelper/model"
	"os"
	"os/exec"
	"strings"
	tpl "text/template"
	"time"
)

func NewTmplObj(tplPath string) *tpl.Template {
	tplName := tplPath[(strings.LastIndex(tplPath, "/") + 1):]
	funcMap := tpl.FuncMap{"Time2Date": comm.Time2Date}
	tmpl := tpl.Must(tpl.New(tplName).Funcs(funcMap).ParseFiles(tplPath))
	return tmpl
}

func CreateTpl(tplPath string, tplData *model.TplModel, outPath string) {
	file, err := os.OpenFile(outPath, os.O_EXCL|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tmpl := NewTmplObj(tplPath)
	err = tmpl.Execute(file, tplData)
	if err != nil {
		panic(err)
	}
}

func Create(rawName string, ip string, port uint) {
	//rawName可以是一个path
	//rawName整体作为go mod init的参数，用于指定go mod init的根目录
	//rawName路径的最后一个dir作为项目的dir
	lastDir := comm.GetLastDirFromPath(rawName)
	// 创建相关文件夹
	mainDir := comm.GetMainDir(lastDir)
	if comm.PathExist(mainDir) {
		panic("Module Exists")
	}

	pbDir := comm.GetPbDir(mainDir)
	svrDir := comm.GetSvrDir(mainDir)
	cliToolDir := comm.GetCliToolDir(mainDir)
	handlerDir := comm.GetHandlerDir(mainDir)
	stubDir := comm.GetStubDir(mainDir)

	err := comm.MakeDir(mainDir)
	if err != nil {
		panic(err)
	}

	// 执行go mod init
	cmdStr := "cd " + mainDir + "; go mod init " + rawName
	cmd := exec.Command("/bin/bash", "-c", cmdStr)
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	err = comm.MakeDir(pbDir)
	if err != nil {
		panic(err)
	}

	err = comm.MakeDir(svrDir)
	if err != nil {
		panic(err)
	}

	err = comm.MakeDir(cliToolDir)
	if err != nil {
		panic(err)
	}

	err = comm.MakeDir(handlerDir)
	if err != nil {
		panic(err)
	}

	err = comm.MakeDir(stubDir)
	if err != nil {
		panic(err)
	}

	// 模版数据组装
	projName, err := comm.RawName2ProjName(lastDir)
	if err != nil {
		panic(err)
	}
	/*
		prefixFromGoSrcPath, err := comm.GetPrefixFromGoSrcPath()
		if err != nil {
			panic(err)
		}
	*/
	tplData := model.NewTplModel()
	tplData.GoModulePath = rawName
	tplData.Date = time.Now()
	tplData.RawName = rawName
	dirName, _ := comm.RawName2DirName(lastDir)
	tplData.DirName = dirName
	tplData.ProjName = projName
	// tplData.PrefixFromGoSrcPath = prefixFromGoSrcPath
	tplData.SvrIp = ip
	tplData.SvrPort = uint32(port)

	// 创建pb/service.proto文件
	pbFilePath := comm.GetPbFilePath(pbDir)
	pbTplPath := comm.GetTplPath4Pb()
	CreateTpl(pbTplPath, tplData, pbFilePath)

	// 创建svr/svr_main.go文件
	svrMainFilePath := comm.GetSvrMainFilePath(svrDir)
	svrTplPath := comm.GetTplPath4Svr()
	CreateTpl(svrTplPath, tplData, svrMainFilePath)

	// 创建svr/handler/handler.go文件
	handlerFilePath := comm.GetHandlerFilePath(handlerDir)
	handlerTplPath := comm.GetTplPath4Handler()
	CreateTpl(handlerTplPath, tplData, handlerFilePath)

	// 创建cli_tool/cli_tool_main.go文件
	cliToolMainFilePath := comm.GetCliToolMainFilePath(cliToolDir)
	cliTplPath := comm.GetTplPath4Cli()
	CreateTpl(cliTplPath, tplData, cliToolMainFilePath)

	// 创建cli_tool/stub/stub.go文件
	stubFilePath := comm.GetStubFilePath(stubDir)
	stubTplPath := comm.GetTplPath4Stub()
	CreateTpl(stubTplPath, tplData, stubFilePath)

	// build.sh文件
	buildFilePath := comm.GetBuildFilePath(mainDir)
	buildTplPath := comm.GetTplPath4Build()
	comm.CopyFile(buildFilePath, buildTplPath)
}
