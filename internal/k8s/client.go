package k8s

import (
	"errors"
	"log"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var errConfig = errors.New("failed to create config")
var errClient = errors.New("failed to create kube client")

type Client struct {
	config *rest.Config
	client kubernetes.Interface
}

// NewClient creates a client to interact with the k8s cluster.
// It will try to use the kube config located in the home folder.
func NewClient() (*Client, error) {
	return NewClientWithKubeConfig("")
}

// NewClientWithKubeConfig creates a client to interact with the k8s cluster.
// a valid kube config path should be provided otherwise it will use the one
// located in the home folder.
func NewClientWithKubeConfig(kubeconfigpath string) (*Client, error) {
	config, err := config(kubeconfigpath)
	if err != nil {
		return nil, err
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Printf("failed to create a client with given kube config: %s", err)
		return nil, errClient
	}

	newClient := Client{
		config: config,
		client: clientSet,
	}
	return &newClient, nil
}

// config reads the kubernetes config from the default location, e.g. $HOME/.kube/config.
func config(kubeconfigpath string) (*rest.Config, error) {
	if home := homedir.HomeDir(); home != "" && kubeconfigpath == "" {
		kubeconfigpath = filepath.Join(home, ".kube", "config")
	}
	if kubeconfigpath == "" {
		log.Println("you should specify a valid kube config path or set up a valid kube config file in home folder")
		return nil, errors.New("please provide or set up a valid kube config file")
	}
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigpath)
	if err != nil {
		log.Printf("failed to build config from config path: %s", err)
		return nil, errConfig
	}
	return config, nil
}
