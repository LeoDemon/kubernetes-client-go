package client_test

import (
	"context"
	"flag"
	"path/filepath"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func TestNode(t *testing.T) {
	t.Logf("get kubernetes nodes: \n")
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err1 := kubernetes.NewForConfig(config)
	if err1 != nil {
		panic(err1.Error())
	}

	nodes, err2 := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err2 != nil {
		panic(err2.Error())
	}
	if nil != nodes {
		for i := 0; i < len(nodes.Items); i++ {
			t.Logf("node[%d]: %s %s", i, nodes.Items[i].Name, nodes.Items[i].Status.Addresses[0].Address)
		}
	}
}
