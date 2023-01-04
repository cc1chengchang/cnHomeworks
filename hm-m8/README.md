
1. create Tls key+cert
ubuntu@master:~/chchang/homeworks/hm-m8$ openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=cncamp.com/O=cncamp" -addext "subjectAltName = DNS:cncamp.com"
Generating a RSA private key
........................................+++++
...........................+++++
writing new private key to 'tls.key'
-----
ubuntu@master:~/chchang/homeworks/hm-m8$ ls
tls.crt  tls.key

2. Create secret
ubuntu@master:~/chchang/homeworks/hm-m8$ sudo kubectl create secret tls cncamp-tls --cert=./tls.crt --key=./tls.key
secret/cncamp-tls created

ubuntu@master:~/chchang/homeworks/hm-m8$ sudo kubectl get secret
NAME                  TYPE                                  DATA   AGE
cncamp-tls            kubernetes.io/tls                     2      37m


3. create pod + service
ubuntu@master:~/chchang/homeworks/hm-m8$ sudo kubectl apply -f ./httpserver.yaml
deployment.apps/httpserver created
service/httpserver created

ubuntu@master:~/chchang/homeworks/hm-m8$ sudo kubectl get svc
NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
httpserver   ClusterIP   10.102.13.177   <none>        80/TCP    34s

ubuntu@master:~/chchang/homeworks/hm-m8$ sudo kubectl get ep
NAME         ENDPOINTS                                  AGE
httpserver   192.168.104.24:8090,192.168.166.153:8090   50s

access service
ubuntu@node1:~$ curl -k -v http://10.102.13.177
*   Trying 10.102.13.177:80...
* TCP_NODELAY set
* Connected to 10.102.13.177 (10.102.13.177) port 80 (#0)
> GET / HTTP/1.1
> Host: 10.102.13.177
> User-Agent: curl/7.68.0
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Fri, 25 Nov 2022 19:01:39 GMT
< Content-Length: 181
< Content-Type: text/plain; charset=utf-8
<
hello [stranger]
===================Details of the http request header:============
Accept=[*/*]
User-Agent=[curl/7.68.0]
===================Server go version:============
Version=
* Connection #0 to host 10.102.13.177 left intact
ubuntu@node1:~$ curl -k -v http://10.102.13.177/healthz
*   Trying 10.102.13.177:80...
* TCP_NODELAY set
* Connected to 10.102.13.177 (10.102.13.177) port 80 (#0)
> GET /healthz HTTP/1.1
> Host: 10.102.13.177
> User-Agent: curl/7.68.0
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Fri, 25 Nov 2022 19:01:51 GMT
< Content-Length: 4
< Content-Type: text/plain; charset=utf-8
<
200
* Connection #0 to host 10.102.13.177 left intact
ubuntu@node1:~$

4. create ingress
ubuntu@master:~/chchang/homeworks/hm-m8$ sudo kubectl create -f ./ingress.yaml
ingress.networking.k8s.io/httpserver-ingress created

ubuntu@master:~/chchang/homeworks/hm-m8$ sudo kubectl get ingress
NAME                 CLASS    HOSTS        ADDRESS   PORTS     AGE
httpserver-ingress   <none>   cncamp.com             80, 443   20s

ubuntu@master:~/chchang/homeworks/hm-m8$ sudo kubectl describe ingress
Name:             httpserver-ingress
Namespace:        default
Address:
Default backend:  default-http-backend:80 (<error: endpoints "default-http-backend" not found>)
TLS:
  cncamp-tls terminates cncamp.com
Rules:
  Host        Path  Backends
  ----        ----  --------
  cncamp.com
              /   httpserver:80 (192.168.104.24:8090,192.168.166.153:8090)
Annotations:  kubernetes.io/ingress.allow-http: false
Events:       <none>


5. clean up
ubuntu@master:~/chchang/homeworks/hm-m8$ sudo kubectl delete -f ./ingress.yaml
ingress.networking.k8s.io "httpserver-ingress" deleted

ubuntu@master:~/chchang/homeworks/hm-m8$ sudo kubectl delete -f ./httpserver.yaml
deployment.apps "httpserver" deleted
service "httpserver" deleted

