package main

import (
	"flag"
	"fmt"
	"github.com/eyesofblue/grpchelper/comm"
	"github.com/eyesofblue/grpchelper/logic"
	"os"
)

func usage() {
	fmt.Printf("Usage: \n\t%s -c %s -n [name] -i [ip] -p [port]\t\t/*创建一个新工程*/\n\t%s -c %s -n [name]\t\t\t\t/*增加一个新接口*/\n", os.Args[0], comm.CMD_CREATEPROJ, os.Args[0], comm.CMD_ADDINTERFACE)
}

func doCreate(rawName string, ip string, port uint) {
	logic.Create(rawName, ip, port)
}

func doAdd(rawName string) {
	logic.Add(rawName)
}

func main() {
	var rawName string
	flag.StringVar(&rawName, "n", "", "name")
	var cmd string
	flag.StringVar(&cmd, "c", "", "cmd")

	var ip string
	flag.StringVar(&ip, "i", "", "ip")

	var port uint
	flag.UintVar(&port, "p", 10317, "port")

	flag.Parse()

	if !comm.IsValidPath(rawName) {
		fmt.Println("Invalid module path")
		os.Exit(-1)
	}
	switch cmd {
	case comm.CMD_CREATEPROJ:
		doCreate(rawName, ip, port)
	case comm.CMD_ADDINTERFACE:
		doAdd(rawName)
	default:
		usage()
	}
}
