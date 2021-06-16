#ifndef HTTPGETPOSTMETHOD_H
#define HTTPGETPOSTMETHOD_H
#include <iostream>
 
class HttpGetPostMethod
{
    public:
        HttpGetPostMethod();
        virtual ~HttpGetPostMethod();
        // Http GET 请求
        int HttpGet(std::string host, std::string path, std::string get_content);
        int HttpGetWithAuth(std::string host,std::string port,  std::string path, std::string get_content, std::string auth);
        // Http POST 请求
        int HttpPost(std::string host, std::string path, std::string post_content);
        std::string get_request_return();
        std::string get_main_text();
        int get_return_status_code();
 
    protected:
 
    private:
        // 记录请求返回的状态码
        int return_status_code_;
        // 记录请求返回所有数据
        std::string request_return_;
        // 记录请求返回的报文部分
        std::string main_text_;
        // HTTP请求过程中使用到的Socket通信部分
        std::string HttpSocket(std::string host, std::string request_str);
        // 将HTTP请求返回的数据分解
        void AnalyzeReturn(void);
};
 
#endif // HTTPGETPOSTMETHOD_H