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

func Create(moduleName string) {
	// 创建相关文件夹
	dirName, err := comm.ModuleName2DirName(moduleName)
	if err != nil {
		panic(err)
	}
	currentDir := comm.GetCurrentDirectory()
	mainDir := currentDir + "/" + dirName
	pbDir := mainDir + "/pb"
	publicDir := mainDir + "/public"
	svrDir := mainDir + "/svr"
	cliToolDir := mainDir + "/cli_tool"
	handlerDir := svrDir + "/handler"
	stubDir := cliToolDir + "/stub"

	err = comm.MakeDir(mainDir)
	if err != nil {
		panic(err)
	}

	err = comm.MakeDir(pbDir)
	if err != nil {
		panic(err)
	}

	err = comm.MakeDir(publicDir)
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
	projName, err := comm.ModuleName2ProjName(moduleName)
	if err != nil {
		panic(err)
	}
	prefixFromGoSrcPath, err := comm.GetPrefixFromGoSrcPath()
	if err != nil {
		panic(err)
	}
	tplData := model.NewTplModel()
	tplData.Date = time.Now()
	tplData.ModuleName = moduleName
	tplData.DirName = dirName
	tplData.ProjName = projName
	tplData.PrefixFromGoSrcPath = prefixFromGoSrcPath
	// FIXME
	tplData.SvrIp = "127.0.0.1"
	tplData.SvrPort = 9999

	// 创建PB文件
	pbFilePath := pbDir + "/" + moduleName + ".proto"
	pbTplPath := comm.GetTplPath4Pb()
	CreateTpl(pbTplPath, tplData, pbFilePath)

	// 创建svr_main文件
	svrMainPath := svrDir + "/" + moduleName + "_svr.go"
	svrTplPath := comm.GetTplPath4Svr()
	CreateTpl(svrTplPath, tplData, svrMainPath)

	// 创建const文件
	constFilePath := publicDir + "/const.go"
	constTplPath := comm.GetTplPath4Const()
	CreateTpl(constTplPath, tplData, constFilePath)

	// 创建cli_tool_main文件
	cliToolMainPath := cliToolDir + "/" + moduleName + "_cli_tool.go"
	cliTplPath := comm.GetTplPath4Cli()
	CreateTpl(cliTplPath, tplData, cliToolMainPath)
}
