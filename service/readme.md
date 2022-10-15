# Service 

```sh
kubectl apply -f deployment.yml
```

## expose service

```sh
kubectl apply -f service.yml
```

ตรวจสอบ service

```sh
kubectl get svc -l app=myapp
```

ทดสอบการ เรียก api ผ่าน port forward

```sh
kubectl port-forward svc/myapp 8080:80
```

ทดลองการเรียก api

```sh
curl localhost:8080
```

## Node port

ทำให้ภายนอก cluster เข้าถึง api เรียกผ่าน Service ด้วยการใช้ NodePort 

```sh
kubectl apply -f nodeport.yml
```

ตรวจสอบ nodeport โดยจะ random เริ่มตั้งแต่ 30000 - 32000 โดย default

```sh
kubectl get svc -l app=myapp
```

ทดสอบการ เรียก api ผ่าน node port

```sh 
curl NODE_IP:NODE_PORT
```

NODE_IP มาจาก IP ของ Node/worker

NODE_PORT มาจาก service port 


## External service

### map external service ด้วย domain name

```sh
kubectl apply -f externalName.yml
```

debug external service  ด้วย pod


```sh
kubectl apply -f busybox.yml
```

install curl บน pod 

```sh
kubectl exec -it busybox -- apk add curl
```

เรียก external service 

```sh
kubectl exec -it busybox -- curl https://mflow -Lk
```

### map external service ด้วย IPs

```sh
kubectl apply -f externalIP.yml
```

เรียก external service 

```sh
kubectl exec -it busybox -- curl https://mflowip -Lk
```

# clean

```sh
kubectl delete -f busybox.yml
kubectl delete -f deployment.yml
kubectl delete -f service.yml
kubectl delete -f nodeport.yml
kubectl delete -f externalIP.yml
kubectl delete -f externalName.yml
```