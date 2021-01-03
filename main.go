package main

import (
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/tool/protobuf/pkg/gen"
	"os"
)

func main() {
	versionFlag := flag.Bool("version", false, "print version and exit")
	flag.Parse()
	if *versionFlag {
		fmt.Println("1.0.0")
		os.Exit(0)
	}

	g := GinGenerator()
	gen.Main(g)
}
