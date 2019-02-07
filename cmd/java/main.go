package main

import (
	"flag"
	"os"
	"fmt"
)

type (
	CmdOptions struct {
		help bool
		version bool
		classpath string
		class string
		args []string
	}
)

func usage() {
	fmt.Printf("%s [-options] class [args...]\n", os.Args[0])
}

func parseCmd() *CmdOptions {
	cmd := new(CmdOptions)
	flag.Usage = usage
	flag.BoolVar(&cmd.help, "help", false, "print help message")
	flag.BoolVar(&cmd.version, "version", false, "print version")
	flag.StringVar(&cmd.classpath, "cp", "", "set classpath")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func main()  {
	opt := parseCmd()
	if opt.class == "" && opt.classpath == "" {
		// no class
		usage()
	} else if opt.version {
		fmt.Println("myvm 0.0.1")
	} else {
		start(opt)
	}
}

func start(opt *CmdOptions) {
	fmt.Printf("class: %s, classpath: %s, args: %v\n", 
		opt.class, opt.classpath, opt.args)
}