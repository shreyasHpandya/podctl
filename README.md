# podctl

Small cli utility to list pods in a cluster

### Usage

    make build

    ./podctl --help
    Usage of ./podctl:
        -kubeconfig string
            (optional) absolute path to the kubeconfig file (default "~/.kube/config")
        -name string
            (optional) Pod name filter
        -namespace string
            (optional) namespace filter
