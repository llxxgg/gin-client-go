package client

import (
	"context"
	"errors"
	"flag"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func GetRestConfig(ctx context.Context) (config *rest.Config, err error) {
	var kubeConfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "absolute path to kube config")
	} else {
		err = errors.New("read config error")
		return
	}
	flag.Parse()

	config, err = clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		klog.Fatal(err)
	}
	return
}

func GetClientSet(ctx context.Context) *kubernetes.Clientset {
	config, err := GetRestConfig(ctx)
	if err != nil {
		return nil
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatal(err)
		return nil
	}
	return clientSet
}
