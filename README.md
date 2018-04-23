# Hello server

This is a demo of a multi-service application.

## Deploy to minikube on your machine

1.  Install [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) and [minikube](https://github.com/kubernetes/minikube/releases).
1.  Start minikube.
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
1.  Deploy to kubernetes.
    ```
    ./build/apply.sh
    ```
1.  Open frontend in browser.
    ```
    minikube service frontend
    ```

## TODO

* Add label for easy teardown of all services.
* Config service persistent volume and only 1 replica
* Don't emit service definitions when 0 replicas?
* Image versioning? Deploy dev version? Local docker registry?
* remove internalPort?
