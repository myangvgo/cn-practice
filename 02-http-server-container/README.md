# Containerize an Http Server

1. Build the httpserver image based on Makefile and Dockerfile

```shell
# in 02-http-server-container directory
make root
# build httpserver to binary
make build
# build httpserver docker image
make release
# push to docker registry
make push
```

2. Run an httpserver container from the httpserver docker image

```shell
# in the ubuntu server run
docker run myangvgo/httpserver:v1.0

# find the container id
docker ps

root@k8snode:~# docker ps
CONTAINER ID   IMAGE                      COMMAND                  CREATED         STATUS         PORTS     NAMES
27febd8187c0   myangvgo/httpserver:v1.0   "/bin/sh -c /httpser…"   3 minutes ago   Up 3 minutes   80/tcp    eager_meninsky

# find the container pid
docker inspect 27febd8187c0 | grep -i pid
root@k8snode:~# docker inspect 27febd8187c0 | grep -i pid
            "Pid": 3552,
            "PidMode": "",
            "PidsLimit": null,

# check ip 
nsenter -t 3552 -n ip addr
root@k8snode:~# nsenter -t 3552 -n ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
5: eth0@if6: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever

# visit http server from ubuntu server
docker inspect 27febd8187c0 | grep -i ipaddress
root@k8snode:~# docker inspect 27febd8187c0  | grep -i ipaddress
            "SecondaryIPAddresses": null,
            "IPAddress": "172.17.0.2",
                    "IPAddress": "172.17.0.2",

curl 172.17.0.2
curl 172.17.0.2/healthz
curl 172.17.0.2/notfound
```