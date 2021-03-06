package logic

import (
	"github.com/eyesofblue/grpchelper/comm"
	"github.com/eyesofblue/grpchelper/model"
	"os"
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
	// 创建相关文件夹
	mainDir := comm.GetMainDir(rawName)
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
	projName, err := comm.RawName2ProjName(rawName)
	if err != nil {
		panic(err)
	}
	prefixFromGoSrcPath, err := comm.GetPrefixFromGoSrcPath()
	if err != nil {
		panic(err)
	}
	tplData := model.NewTplModel()
	tplData.Date = time.Now()
	tplData.RawName = rawName
	dirName, _ := comm.RawName2DirName(rawName)
	tplData.DirName = dirName
	tplData.ProjName = projName
	tplData.PrefixFromGoSrcPath = prefixFromGoSrcPath
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
