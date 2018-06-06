package main

import (
	"flag"
	"fmt"
	"github.com/eyesofblue/grpchelper/comm"
	"github.com/eyesofblue/grpchelper/logic"
	"os"
)

func usage() {
	fmt.Printf("Usage: %s -c [cmd<%s|%s>] -n [name]\n", os.Args[0], comm.CMD_CREATEPROJ, comm.CMD_ADDINTERFACE)
}

func doCreate(rawName string) {
	logic.Create(rawName)
}

func doAdd(rawName string) {
	logic.Add(rawName)
}

func main() {
	var rawName string
	flag.StringVar(&rawName, "n", "", "name")
	var cmd string
	flag.StringVar(&cmd, "c", "", "cmd")

	flag.Parse()

	if !comm.IsValidRawName(rawName) {
		fmt.Println("Invalid name")
		os.Exit(-1)
	}
	switch cmd {
	case comm.CMD_CREATEPROJ:
		doCreate(rawName)
	case comm.CMD_ADDINTERFACE:
		doAdd(rawName)
	default:
		usage()
	}
}
