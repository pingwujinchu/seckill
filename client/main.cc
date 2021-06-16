#include <iostream>
#include <string>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <sys/types.h>
#include <netinet/in.h>
#include <stdlib.h>
#include <string>
#include <cstring>
#include <memory>
#include <cstdlib>
#include <pthread.h>
#include "../tools/cmdline.h"

using namespace std;

void *execThread(void *arg){
    string cmd =  *static_cast<string*>(arg);
    system((char *)cmd.data());
    return ((void *)0);
}

int main(int argc, char *argv[])
{
    // 创建一个命令行解析器
    cmdline::parser a;
    // 添加指定类型的输入参数
    // 第一个参数：长名称
    // 第二个参数：短名称（'\0'表示没有短名称）
    // 第三个参数：参数描述
    // 第四个参数：bool值，表示该参数是否必须存在（可选，默认值是false）
    // 第五个参数：参数的默认值（可选，当第四个参数为false时该参数有效）
    a.add<string>("service", 's', "服务类型", true, "可选：seckill，product，order");

    a.add<string>("host", 'h', "服务地址", false, "127.0.0.1");
    // 第六个参数用来对参数添加额外的限制
    // 这里端口号被限制为必须是1到65535区间的值，通过cmdline::range(1, 65535)进行限制
    a.add<int>("port", 'p', "port number", false, 9988, cmdline::range(1, 65535));
    // cmdline::oneof() 可以用来限制参数的可选值
    a.add<string>("type", 't', "protocol type", false, "http", cmdline::oneof<string>("http", "https", "ssh", "ftp"));

    a.add<string>("arg", 'a', "路径参数", false, "");

    a.add<int>("number", 'n', "并发数", false, 1);
    // 也可以定义bool值
    // 通过调用不带类型的add方法
    a.add("gzip", '\0', "gzip when transfer");
    // 运行解析器
    // 只有当所有的参数都有效时他才会返回
    a.parse_check(argc, argv);
    // bool值可以通过调用exsit()方法来判断
    string host = a.get<string>("host");
    string port = std::to_string(a.get<int>("port"));
 
    string service = a.get<string>("service");

    int number = a.get<int>("number");

    string cmd = "curl --user server:love " + host + ":" + port + "/" + service;
    
    if (service.compare("seckill") == 0) {
        cmd = cmd + "/" + a.get<string>("arg") + " -X POST";
    }else if(service.compare("status") == 0){
        cmd = cmd + "/" + a.get<string>("arg");
    }

    // 解析json
    cmd = cmd + " -s | jq";

    pthread_t tids[number];
    for(int i = 0; i < number; ++i)
    {
        //参数依次是：创建的线程id，线程参数，调用的函数，传入的函数参数
        int ret = pthread_create(&tids[i], NULL, &execThread, static_cast<void*>(new std::string(cmd)));
        if (ret != 0)
        {
           cout << "pthread_create error: error_code=" << ret << endl;
        }
    }
    //等各个线程退出后，进程才结束，否则进程强制结束了，线程可能还没反
    pthread_exit(NULL);
    return 0;
}