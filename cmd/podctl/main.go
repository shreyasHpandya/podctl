package main

import (
	"fmt"

	"github.com/shreyasHpandya/podctl/pkg/client"
	"github.com/shreyasHpandya/podctl/pkg/config"
	"github.com/shreyasHpandya/podctl/pkg/pod"
	v1 "k8s.io/api/core/v1"
)

func main() {
	config.InitConf()

	pods, err := pod.List(client.GetClient(), *config.Conf.Namespace, *config.Conf.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	printer(pods)
}

func printer(pods []v1.Pod) {
	fmt.Println("NameSpace               ", "PodName                      ", "Count")
	for _, p := range pods {
		fmt.Printf("%-25.24s", p.Namespace)
		fmt.Printf("%-30.29s", p.Name)
		upContainers := 0
		for _, st := range p.Status.ContainerStatuses {
			if st.Ready {
				upContainers += 1
			}
		}
		fmt.Printf("%v/%v\n", upContainers, len(p.Status.ContainerStatuses))
	}
}
