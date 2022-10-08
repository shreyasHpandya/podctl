package podprinter

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
)

func Printer(pods []v1.Pod) {
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
