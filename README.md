## port_forward
A base forward use TCP UDP KCP 



### Link 

* `https://github.com/1lann` thx man


### Quick

* 编译运行即可 主要功能 用于做中转加速。
* 大致方式 部署在阿里云或者其他BGP多线运营商，转发流量到其他地区的机房（US，UK等）
* 实现了`Socat`大致的UDP TCP转发功能 

### Test And Check

```bash
./port_forward  0.0.0.0:2341 47.93.244.110:52 udp
./port_forward  0.0.0.0:2341 47.93.244.110:22 tcp
# 检查端口是否在线
nc -vuz 127.0.0.1  2341


found 0 associations
found 1 connections:
     1:	flags=82<CONNECTED,PREFERRED>
	outif (null)
	src 127.0.0.1 port 60327
	dst 127.0.0.1 port 2341
	rank info not available

Connection to 127.0.0.1 port 2341 [udp/xiostatus] succeeded!
```

### Start and Stop

* release 里面是编译好的可以直接，执行`install.sh`

```bash
# start
port_forward 0.0.0.0:53 47.93.244.11:53 udp

#stop 

ps -ef | grep port_forward | awk '{print$2}' | xargs kill -9 

```


## One Key

```bash
wget -N https://github.com/uk0/port_forward/raw/master/release.run  && chmod +x ./release.run  && sh -c ./release.run
```