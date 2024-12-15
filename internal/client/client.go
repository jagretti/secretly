package client

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func BuildKubernetesClient() (client kubernetes.Interface, err error) {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}
