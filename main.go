package main

import (
	"fmt"
	"os"

	"github.com/koron/gtc/gtcore"
)

func main() {
	err := gtcore.DefaultCatalog.Merge([]gtcore.Tool{
		{
			Path: "github.com/sohgett/mygtc",
			Desc: "My own go tools catalog",
		},
		{
			Path: "github.com/mattn/sudo",
			Desc: "sudo for windows",
		},
	}...).Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
	}
}
