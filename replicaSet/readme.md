# replicaSet

```sh
kubectl apply -f gb.yml
```

ตรวจสอบจำนวน pods 

```sh
kubectl get pod -l app=guestbook
```


ทำการเพิ่มจำนวน pods หรือที่เรียกว่า replicas

```sh
kubectl apply -f gbscale.yml
```

# clean

```sh
kubectl delete -f gbscale.yml
```