package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	E "github.com/metux/go-metabuild/engine"
)

func main() {
	var defaults string
	var conf string

	flag.StringVar(&defaults, "global", "", "global defaults yaml")
	flag.StringVar(&conf, "conf", "", "project config yaml")
	flag.Parse()

	if len(defaults) == 0 || len(conf) == 0 {
		fmt.Printf("Usage: %s -conf <fn> -global <fn>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	log.Println("global", defaults)

	stage := E.StagePackage

	engine, err := E.Load(conf, defaults)
	if err != nil {
		log.Printf("ERR: %s\n", err)
	} else {
		engine.RunStage(stage)
	}
}
