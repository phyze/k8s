# Ingress

## Allow all and HTTP

```sh
kubectl apply -f bookstore.yml
kubectl apply -f financial.yml
```

```sh
kubectl apply -f ingress-http.yml
```

observe ingress

```sh
kubectl get ing 
```

try to call bookstore service 

```sh
curl NODE_IP:NODE_PORT/bookstore
```

```sh
curl NODE_IP:NODE_PORT/financial
```



## HTTPS 


generate self signed

```sh
openssl req -subj '/CN=YOUR_DOMAIN/O=YOUR_ORGANIZATION/C=YOU_COUNTRY' -new -newkey rsa:2048 -sha256 -days 365 -nodes -x509 -keyout server.key -out server.crt
```

create tls.yml like below

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: tls
type: kubernetes.io/tls
data:
  tls.crt: 
  tls.key: 
```

encode crt and key to base64 then copy base64 paste on tls.yml

```sh
cat server.crt | base64
```

```sh
cat server.key | base64
```

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: tls
type: kubernetes.io/tls
data:
  tls.crt: |
    LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS...
  tls.key: |
    LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUV2UUl...
```

apply tls.yml

```sh
kubectl apply tls.yml
```

apply https

```sh
kubectl apply -f ingress-https.yml
```


observe ingress

```sh
kubectl get ing
```

try to call bookstore and financial 

```sh
curl https://YOU_DOMAIN:NODE_PORT/bookstore --insecure --resolve YOUR_DOMAIN:NODE_PORT:NODE_IP
```

```sh
curl https://YOU_DOMAIN:NODE_PORT/financial --insecure --resolve YOUR_DOMAIN:NODE_PORT:NODE_IP
```

observe cert

```sh
curl https://YOU_DOMAIN:NODE_PORT/bookstore --insecure --resolve YOUR_DOMAIN:NODE_PORT:NODE_IP -Ivv
```

try to call api via http  but ingress not allow

```sh
curl http://YOU_DOMAIN:NODE_PORT/financial --resolve YOUR_DOMAIN:NODE_PORT:NODE_IP
```

allow http 

```sh
kubectl apply -f ingress-https-allow-http.yml
```

```sh
curl http://YOU_DOMAIN:NODE_PORT/financial --resolve YOUR_DOMAIN:NODE_PORT:NODE_IP
```

# Clean 

```sh
kubectl delete -f ingress-https-allow-http.yml
kubectl delete -f ingress-http.yml
kubectl delete -f tls.yml
kubectl delete -f financial.yml
kubectl delete -f bookstore.yml
```