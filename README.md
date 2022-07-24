# goku
kubectl trainer inspired by the [kubectl cheat sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/) site. The idea is to wrap kubectl commands, trying to abbreviate them, displaying the original command using `kubectl` when the user calls goku.

## How to test

* to run integration tests run the following command

```sh
go test -timeout 2s -run ^TestNewClient$ github.com/fernandoocampo/goku/internal/k8s -integration
```

## Known issues

1. if you got `failed to create a client with given kube config: exec plugin: invalid apiVersion "client.authentication.k8s.io/v1alpha1"` it means that you have a kube config using a deprecated version api. The `client.authentication.k8s.io/v1alpha1` has been deprecated and removed from Kubernetes 1.24. 

    * Update the $HOME/.kube/config and change these sections

    ```yaml
    apiVersion: client.authentication.k8s.io/v1alpha1
    ```
    to
    ```yaml
    apiVersion: client.authentication.k8s.io/v1beta1
    ```

    * Update the AWS CLI to the latest version and regenerate the kubeconfig with:

    ```sh
    aws eks update-kubeconfig --name ${EKS_CLUSTER_NAME} --region ${REGION}
    ```

