package pod

import (
	"context"

	v1 "k8s.io/api/core/v1"
	// "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func List(clientset *kubernetes.Clientset, namespace string, podname string) ([]v1.Pod, error) {
	if podname == "" {
		pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})

		return pods.Items, err
	}

	pod, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), podname, metav1.GetOptions{})
	return []v1.Pod{*pod}, err
}
