// Central package to hold all the configs, can be accessed in all packges
package config

import (
	"flag"
	"path/filepath"

	"k8s.io/client-go/util/homedir"
)

type ConfType struct {
	Kubeconfig *string
	Namespace  *string
	Name       *string
}

var Conf ConfType

func init() {
	if home := homedir.HomeDir(); home != "" {
		Conf.Kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		Conf.Kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	Conf.Namespace = flag.String("namespace", "", "(optional) namespace filter")
	Conf.Name = flag.String("name", "", "(optional) Pod name filter")
}
func InitConf() {
	flag.Parse()
}
