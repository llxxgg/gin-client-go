package service

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	"operator/client"
)

func GetNamespace(ctx context.Context) []string{
	clientSet := client.GetClientSet(ctx)
	if clientSet == nil {
		return nil
	}
	namespaceList, err := clientSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		klog.Fatal(err)
		return nil
	}
	namespaces := namespaceList.Items
	var result []string
	for _, namespace := range namespaces {
		result = append(result, namespace.Name)
	}
	return result
}
