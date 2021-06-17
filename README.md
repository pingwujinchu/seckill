# seckill
该项目使用bazel进行打包管理
# 项目目录
```sh
├── build.sh
├── cache
│   ├── cache
│   ├── cache.cpp
│   └── redis.h
├── client
│   ├── BUILD
│   └── main.cc
├── db
│   ├── db.cpp
│   ├── db.h
│   ├── entity.cpp
│   ├── entity.h
│   ├── init.sql
│   ├── test
│   └── test.cpp
├── Dockerfile
├── LICENSE
├── png
├── README.md
├── server
│   ├── BUILD
│   ├── config
│   │   ├── config.toml
│   │   └── init.sql
│   ├── docker-compose.yml
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── pkg
│   │   ├── cache
│   │   │   └── redis.go
│   │   ├── config
│   │   │   └── config.go
│   │   ├── controller
│   │   │   ├── api
│   │   │   │   └── products.go
│   │   │   ├── base.go
│   │   │   └── dto
│   │   │       ├── code.go
│   │   │       └── response.go
│   │   ├── model
│   │   │   ├── dao.go
│   │   │   └── type.go
│   │   ├── rabitmq
│   │   │   └── rabitmq.go
│   │   └── router
│   │       └── router.go
│   ├── readme.txt
│   └── sec_kill_db
├── tools
│   ├── BUILD
│   └── cmdline.h
└── WORKSPACE
```

# client
client由c++实现

# server
server 由go语言实现


# 构建
构建之前需要安装docker
执行build.sh
```
sudo ./build.sh
```

# 运行
```
docker-compose up
```
客户端查看秒杀产品
```
./client  -s product
```
客户端发起秒杀
```
./client -s seckill -a 1 -n 10
```
其中-n 指的是并行数

查看当前状态
```
./client -s status -a xxxxxx 
```

# 数据库设计

```
+-----------------------+
| Tables_in_sec_kill_db |
+-----------------------+
| orders                |
| products              |
| sec_kills             |
+-----------------------+
```

# 服务端设计
服务端使用go语言的gin框架，主要分为如下模块：

## 秒杀产品管理
秒杀产品存在数据库sec_kills中，主要保存秒杀的产品，及秒杀开始时间及结束时间。
会在项目启动时，将其加载到redis缓存中。这个数据也是读多写少的数据。
## 库存查询
```
因为库存数据是查询多于修改，因此将库存数据放到缓存中
```
## 秒杀任务提交
```
通过缓存数据验证秒杀请求合法后，将秒杀请求传递到rabbitmq队列中。
```
## 库存锁定
```
队列消费者从队列拿到请求后，将会基于乐观锁去获取库存，如果库存>0, 则使用数据库事务修改库存，并生成order记录。并同步更新缓存。
```

## 支付管理
```

```
## 缓存同步
```
```

## 安全性保证
```
服务端提供的接口经过加密，客户端和服务端之间通信使用认证信息进行认证。
```

# 问题排查
如果遇到无法连接数据库，是没有创建数据库，需要创建一下数据库，命令如下：
```
docker exec -it  XXX sh 

mysql -u root -p jdllq@cclfc

create database sec_kill_db
```

如果docker-compose无法启动， 请安装最新版docker-compose， 安装地址如下：
```
sudo curl -L "https://github.com/docker/compose/releases/download/1.24.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
```

