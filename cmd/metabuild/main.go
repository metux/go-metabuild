package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	E "github.com/metux/go-metabuild/engine"
)

var (
	argDefaults string
	argConf     string
	args        []string
)

func usage() {
	fmt.Printf("Usage: %s -conf <fn> -global <fn> [command ...]\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Printf("Available commands\n")
	fmt.Printf("    build [stage]       run build\n")
	os.Exit(1)
}

func cmdBuild() {
	stage := E.StagePackage

	if len(args) > 1 {
		stage = E.Stage(args[1])
	}

	engine, err := E.Load(argConf, argDefaults)
	if err != nil {
		log.Printf("ERR: %s\n", err)
	} else {
		engine.RunStage(stage)
	}
}

func main() {
	flag.StringVar(&argDefaults, "global", "", "global defaults yaml")
	flag.StringVar(&argConf, "conf", "", "project config yaml")
	flag.Parse()
	args = flag.Args()

	if argDefaults == "" || argConf == "" || len(args) == 0 {
		usage()
	}

	switch args[0] {
	case "build":
		cmdBuild()
	default:
		usage()
	}
}
