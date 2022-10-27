# replicaSet

```sh
kubectl apply -f gb.yml
```

ตรวจสอบจำนวน pods 

```sh
kubectl get pod -l app=guestbook
```

## scale up

ทำการเพิ่มจำนวน pods หรือที่เรียกว่า replicas

```sh
kubectl apply -f gbscale.yml
```

## self-healing ด้วยการลบ pod

```sh
kubectl delete pod -l app=guestbook
```

ตรวจสอบ Pod พบว่ากำลังสร้างขึ้นมาใหม่

```sh
kubectl get pod -l app=guestbook
```

## ทำการ update version ของ application 

```sh
kubectl apply -f gbv2.yml
```

```sh
kubectl get pod -l app=guestbook
```

```sh
kubectl get rs/guestbook -o jsonpath="{.spec.template.spec.containers[?(@.name=='guestbook')].image}"
```

จะพบว่าไม่มีการเปลียน image เป็น v2

# clean

```sh
kubectl delete -f gbscale.yml
```