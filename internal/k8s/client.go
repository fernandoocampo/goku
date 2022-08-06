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
var errNilConfig = errors.New("the given config object is empty")

type client struct {
	config *rest.Config
	client kubernetes.Interface
}

// NewClientWithKubeConfig creates a client to interact with the k8s cluster.
// based on the given kube config object.
func NewClient(config *rest.Config) (*client, error) {
	if config == nil {
		return nil, errNilConfig
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Printf("failed to create a client with given kube config: %s", err)
		return nil, errClient
	}

	newClient := client{
		config: config,
		client: clientSet,
	}
	return &newClient, nil
}

// LoadConfig reads the kubernetes config from the given path e.g. $HOME/.kube/config.
func LoadConfig(kubeconfigpath string) (*rest.Config, error) {
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

// LoadDefaultKubeConfig load the default kube config file
func LoadDefaultKubeConfig() (*rest.Config, error) {
	return LoadConfig(getDefaultKubeConfigPath())
}

func getDefaultKubeConfigPath() string {
	return filepath.Join(homedir.HomeDir(), ".kube", "config")
}
