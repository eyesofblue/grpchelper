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

func Create(rawName string) {
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
	// FIXME
	tplData.SvrIp = "127.0.0.1"
	tplData.SvrPort = 9999

	// 创建pb/service.proto文件
	pbPath := comm.GetPbFilePath(pbDir)
	pbTplPath := comm.GetTplPath4Pb()
	CreateTpl(pbTplPath, tplData, pbPath)

	// 创建svr/svr_main.go文件
	svrMainPath := comm.GetSvrMainFilePath(svrDir)
	svrTplPath := comm.GetTplPath4Svr()
	CreateTpl(svrTplPath, tplData, svrMainPath)

	// 创建svr/handler/handler.go文件
	handlerPath := comm.GetHandlerFilePath(handlerDir)
	handlerTplPath := comm.GetTplPath4Handler()
	CreateTpl(handlerTplPath, tplData, handlerPath)

	// 创建cli_tool/cli_tool_main.go文件
	cliToolMainPath := comm.GetCliToolMainFilePath(cliToolDir)
	cliTplPath := comm.GetTplPath4Cli()
	CreateTpl(cliTplPath, tplData, cliToolMainPath)

	// 创建cli_tool/stub/stub.go文件
	stubPath := comm.GetStubFilePath(stubDir)
	stubTplPath := comm.GetTplPath4Stub()
	CreateTpl(stubTplPath, tplData, stubPath)
}
