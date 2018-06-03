package main

import (
	"github.com/grpchelper/model"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err)
	}
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

func main() {
	tmpl, err := template.ParseFiles("./tpl/cli.tpl")
	if err != nil {
		panic(err)
	}

	pathPrefix := getCurrentDirectory()
	date := time.Now().Format("2006-01-02")
	cli := model.CliTplModel{ModuleName: "cashredpack", Date: date, SvrIp: "127.0.0.1", SvrPort: "8888", PathPrefix: pathPrefix}

	err = tmpl.Execute(os.Stdout, cli) //将struct与模板合成，合成结果放到os.Stdout里
	if err != nil {
		panic(err)
	}
}
