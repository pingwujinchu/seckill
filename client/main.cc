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
#include "../tools/cmdline.h"
#include "../client/HttpGetPostMethod.h"

using namespace std;

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
    // 也可以定义bool值
    // 通过调用不带类型的add方法
    a.add("gzip", '\0', "gzip when transfer");
    // 运行解析器
    // 只有当所有的参数都有效时他才会返回
    // 如果有无效参数，解析器会输出错误消息，然后退出程序
    // 如果有'--help'或-?'这样的帮助标识被指定，解析器输出帮助信息，然后退出程序
    a.parse_check(argc, argv);
    // 获取输入的参数值
    cout << a.get<string>("type") << "://"
         << a.get<string>("host") << ":"
         << a.get<int>("port") << endl;
    // bool值可以通过调用exsit()方法来判断
    
    string auth = "Basic c2VydmVyOmxvdmU=";
    string host = a.get<string>("host");
    cout << host <<endl;
    string port = "" + a.get<int>("port");
    string service = a.get<string>("service");
    shared_ptr<HttpGetPostMethod> http_request = make_shared<HttpGetPostMethod>();
    int ret  = http_request->HttpGetWithAuth(host, port, service, "showapi_appid=52875&showapi_sign=59c54e39583740bf9708c645c389c9ec&fundCode=519185&needDetails=1", auth);
    if(ret == -1) {
        cout << "Http Socket error!" << endl;
    }
    cout << ret << endl;
    if (a.exist("gzip"))
        cout << "gzip" << endl;
}