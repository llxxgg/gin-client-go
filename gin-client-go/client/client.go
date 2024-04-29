package client

import (
	"context"
	"errors"
	"flag"
	"path/filepath"

	"k8s.io/client-go/rest"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
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
	// namespaceList, err := clientSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	//
	//	if err != nil {
	//		klog.Fatal(err)
	//		return
	//	}
	//
	// namespaces := namespaceList.Items
	//
	//	for _, namespace := range namespaces {
	//		fmt.Println(namespace.Name)
	//	}
}
