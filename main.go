package main

import (
	"flag"
	"fmt"

	buildlist "github.com/hmiyado/bitrise-check-queued-builds/buildlist"
)

func main() {
	flag.Parse()
	appSlug := flag.Arg(0)

	fmt.Println(buildlist.ListNum(appSlug))
}
