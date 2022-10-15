# configmap

apply configmap

```sh
kubectl apply -f cm.yml
```

lookup

```sh
kubectl apply -f nginx.yml
```

forward port to nginx

```sh
kubectl port-forward pod/nginx 8080:80
```

call nginx

```sh
curl localhost:8080
```