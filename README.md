# Kubernetes dashboard

## Start Minikube

```
minikube start --driver=docker
```

## Namespaces Operations

### List

Return the namespaces in the k8s cluster. ```curl -X GET localhost:8085/k8s/v1/namespaces```
Response:
```
[
    {
        "name": "default",
        "status": "Active",
        "age": "2021-11-14T23:06:56Z"
    },
    {
        "name": "kube-node-lease",
        "status": "Active",
        "age": "2021-11-14T23:06:49Z"
    }
]
```