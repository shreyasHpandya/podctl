package client

import (
	"fmt"
	"os"

	"github.com/shreyasHpandya/podctl/pkg/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var client *kubernetes.Clientset

func InitClient() {
	var err error
	config, err := clientcmd.BuildConfigFromFlags("", *config.Conf.Kubeconfig)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// create the clientset
	client, err = kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func GetClient() *kubernetes.Clientset {
	if client == nil {
		InitClient()
	}
	return client
}
