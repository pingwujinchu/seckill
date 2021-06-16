# seckill
该项目使用bazel进行打包管理
# 项目目录
```sh
├── [  67]  build.sh
├── [4.0K]  cache
│   ├── [ 14K]  cache
│   ├── [ 350]  cache.cpp
│   └── [ 939]  redis.h
├── [4.0K]  client
│   ├── [ 151]  BUILD
│   └── [1.7K]  main.cc
├── [4.0K]  db
│   ├── [2.7K]  db.cpp
│   ├── [ 865]  db.h
│   ├── [ 667]  entity.cpp
│   ├── [ 799]  entity.h
│   ├── [ 803]  init.sql
│   ├── [ 28K]  test
│   └── [1.4K]  test.cpp
├── [   0]  Dockerfile
├── [1.0K]  LICENSE
├── [  63]  README.md
├── [4.0K]  server
│   ├── [ 263]  BUILD
│   ├── [   0]  Dockerfile
│   ├── [ 151]  go.mod
│   ├── [6.3K]  go.sum
│   ├── [ 204]  main.go
│   └── [4.0K]  pkg
│       ├── [4.0K]  cache
│       ├── [4.0K]  dao
│       │   └── [1.6K]  dao.go
│       ├── [4.0K]  entity
│       │   └── [ 379]  type.go
│       └── [4.0K]  kafka
├── [4.0K]  tools
│   ├── [ 165]  BUILD
│   └── [ 18K]  cmdline.h
└── [ 592]  WORKSPACE
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

# 问题排查
如果遇到无法连接数据库，是没有创建数据库，需要创建一下数据库，命令如下：
```
docker exec -it  XXX sh 

mysql -u root -p jdllq@cclfc

create database sec_kill_db
```

