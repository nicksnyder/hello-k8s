# Demo application

This is a demo of a multi-service application.

## Deploy to minikube

1.  Install minikube
    ```
    $ minikube start
    Starting local Kubernetes v1.10.0 cluster...
    Starting VM...
    Downloading Minikube ISO
    150.53 MB / 150.53 MB [============================================] 100.00% 0s
    Getting VM IP address...
    Moving files into cluster...
    Downloading kubelet v1.10.0
    Downloading kubeadm v1.10.0
    Finished Downloading kubeadm v1.10.0
    Finished Downloading kubelet v1.10.0
    Setting up certs...
    Connecting to cluster...
    Setting up kubeconfig...
    Starting cluster components...
    Kubectl is now configured to use the cluster.
    Loading cached images from config file.
    ```
1.  Deploy to kubernetes
    ```
    helm template ./chart > /tmp/chart && kubectl apply -f /tmp/chart
    ```
