## godman
### 
### Containers from scratch  
###
```shell
make build
```

```shell
./godman run alpine /bin/bash
```

```example
go run cmd/main.go run alpine  /bin/sh                     
Using config file: /home/andrey/git_project/golang/godman/config/config.yaml
Init images paths..
Hello godman!
Starter PID: 108149
Using config file: /home/andrey/git_project/golang/godman/config/config.yaml
Init images paths..
INFO[0000] Downloading metadata for alpine:latest, please wait... 
INFO[0002] Id: 8ca4688f4f356596b5ae539337c9941abc78eda10021d35cbc52659c74d9b443 
INFO[0002] Checking if image exists under another name... 
INFO[0002] /home/andrey/git_project/golang/godman/containers/tmp/8ca4688f4f356596b5ae539337c9941abc78eda10021d35cbc52659c74d9b443/8ca4688f4f356596b5ae539337c9941abc78eda10021d35cbc52659c74d9b443.tar 
INFO[0004] Uncompressing layer to: /home/andrey/git_project/golang/godman/containers/storage/overlay/96526aa774ef0126ad0fe9e9a95764c5fc37f409ab9e97021e7b4775d82bf6fa/diff  
INFO[0004] Save Layer 96526aa774ef0126ad0fe9e9a95764c5fc37f409ab9e97021e7b4775d82bf6fa.  
change root to /home/andrey/git_project/golang/godman/containers/storage/containers/5b4390987744c8ee2a86ddd415d13455/merged
mount overlay fs root ...
Container PID: 1
/ # cat /etc/os-release 
NAME="Alpine Linux"
ID=alpine
VERSION_ID=3.18.4
PRETTY_NAME="Alpine Linux v3.18"
HOME_URL="https://alpinelinux.org/"
BUG_REPORT_URL="https://gitlab.alpinelinux.org/alpine/aports/-/issues"
/ # 
```
### Show images 
```shell
./godman images
```
```
go run cmd/main.go images                                         
Using config file: /home/andrey/git_project/golang/godman/config/config.yaml
Init images paths..
NAME                                          TA      IMAGE ID                                                          
registry.home.local/busybox-kaniko            latest  14315177050d00d01658077ac50dfd2f9c28cd5a062d3d5f8ccca8b66ea94f19  
gcr.io/kubernetes-e2e-test-images/echoserver  2.2     4081d9a831083d9e57c49a95632feaf0103bd4db2c9fa1e01b48b7b1136a946d  
alpine                                        latest  8ca4688f4f356596b5ae539337c9941abc78eda10021d35cbc52659c74d9b443  
registry.home.local/busybox                   latest  a416a98b71e224a31ee99cff8e16063554498227d2b696152a9c3e0aa65e5824  
```