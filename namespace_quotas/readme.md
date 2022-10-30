# Namespace and quota


## namespace

create namespace 

```sh
kubectl apply -f ns.yml -n training
```

## resource quota

```sh
kubectl apply -f priority.yml 
kubectl apply -f rq.yml -n training
```

ตรวจสอบ  resource quota

```sh
kubectl describe resourcequota -n training
```

deploy app

```sh
kubectl apply -f bookstore-h.yml -n training
kubectl apply -f bookstore-m.yml -n training
kubectl apply -f bookstore-l.yml -n training
```

## limit resource 

```sh
kubectl apply -f limit.yml -n training
```


สังเกตว่า จะแสดง error ว่าเกิน limit resource

```sh
kubectl apply -f pod-over-resource.yml -n training
```

# clean

```sh
kubectl delete -f ns.yml
kubectl delete -f limit.yml
```