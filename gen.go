package main

import (
	"github.com/eyesofblue/grpchelper/comm"
	"github.com/eyesofblue/grpchelper/model"
	"os"
	"text/template"
	"time"
)

func main() {
	// moduleName := "cashredpack"
	// fileName := moduleName + ".go"
	goPath := comm.GetGoPath()
	if len(goPath) == 0 {
		panic("Not Found GOPATH")
	}
	goSrcPath := goPath + "/src"

	tplPath := goSrcPath + "/github.com/eyesofblue/grpchelper/tpl/cli.tpl"

	funcMap := template.FuncMap{"Time2Date": comm.Time2Date}
	tmpl := template.Must(template.New("cli.tpl").Funcs(funcMap).ParseFiles(tplPath))

	prefixFromGoPath, err := comm.GetPrefixFromGoSrcPath()
	if err != nil {
		panic(err)
	}

	date := time.Now() //.Format("2006-01-02")
	cli := model.CliTplModel{ModuleName: "cashredpack", Date: date, SvrIp: "127.0.0.1", SvrPort: 8888, PrefixFromGoPath: prefixFromGoPath}

	err = tmpl.Execute(os.Stdout, cli) //将struct与模板合成，合成结果放到os.Stdout里
	if err != nil {
		panic(err)
	}
}
