package main

import (
	"fmt"

	"github.com/shreyasHpandya/podctl/internal/podprinter"
	"github.com/shreyasHpandya/podctl/pkg/client"
	"github.com/shreyasHpandya/podctl/pkg/config"
	"github.com/shreyasHpandya/podctl/pkg/pod"
)

func main() {
	//pasrse input flags and set it to global config object
	config.InitConf()
	runCmd()
}

func runCmd() {
	pods, err := pod.List(client.GetClient(), *config.Conf.Namespace, *config.Conf.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	podprinter.Printer(pods)
}
