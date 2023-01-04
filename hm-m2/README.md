
how to run

ubuntu@VM-0-43-ubuntu:~/chchang/homeworks/hm2/httpserver$ docker build . -t httpserver:v1.0
ubuntu@VM-0-43-ubuntu:~/chchang/homeworks/hm2/httpserver$ docker image ls
REPOSITORY   TAG       IMAGE ID       CREATED         SIZE
httpserver   v1.0      98853362d448   7 seconds ago   85.2MB
ubuntu       latest    216c552ea5ba   11 days ago     77.8MB
ubuntu@VM-0-43-ubuntu:~/chchang/homeworks/hm2/httpserver$
ubuntu@VM-0-43-ubuntu:~/chchang/homeworks/hm2/httpserver$ docker run -d httpserver:v1.0
95e5d3402b6ed68fd2f61059b8aceaeed03c52473fff1e5b68618463c953d82e
ubuntu@VM-0-43-ubuntu:~/chchang/homeworks/hm2/httpserver$ docker ps
CONTAINER ID   IMAGE             COMMAND                  CREATED          STATUS          PORTS     NAMES
95e5d3402b6e   httpserver:v1.0   "/bin/sh -c /httpser…"   34 seconds ago   Up 33 seconds   80/tcp    suspicious_burnell
ubuntu@VM-0-43-ubuntu:~/chchang/homeworks/hm2/httpserver$
ubuntu@VM-0-43-ubuntu:~/chchang/homeworks/hm2/httpserver$ docker inspect --format "{{.State.Pid}}" 95e5d3402b6e
36888
ubuntu@VM-0-43-ubuntu:~/chchang/homeworks/hm2/httpserver$

ubuntu@VM-0-43-ubuntu:~/chchang/homeworks/hm2/httpserver$ sudo nsenter -t 36888  -n ip a s
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
4: eth0@if5: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
ubuntu@VM-0-43-ubuntu:~/chchang/homeworks/hm2/httpserver$

