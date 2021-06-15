/*************************************************************************
    > File Name: test.cpp
    > Author:fengxin 
    > Mail:903087053@qq.com 
    > Created Time: 2017年07月21日 星期五 16时43分55秒
 ************************************************************************/

#include<iostream>
#include <string>
#include <fstream>
#include <sstream>
#include <stdlib.h>
#include <vector>
#include <cstring>
#include"db.h"
using namespace std;


//从文件读入到string里
string readFileIntoString(char *filename)
{
    ifstream ifile(filename);
    //将文件读入到ostringstream对象buf中
    ostringstream buf;
    char ch;
    while (buf && ifile.get(ch))
        buf.put(ch);
    //返回与流对象buf关联的字符串
    return buf.str();
}

vector<string> split(const string &str, const  string  &delim){
        vector<string>  res;
        
}

int main(int argc, char *argv[])
{
    DB db; 
    printf("test");
    //连接数据库
    db.initDB("127.0.0.1","root","123456","sec_kill_db");
    printf("test");
    //将用户信息添加到数据库
    bool  result =  db.exeSQL("INSERT accounts values('fengxin','123');");
    cout << result;
    if (!result){
             string data = readFileIntoString("init.sql");
             db.exeSQL(data);
    }
    db.exeSQL("INSERT accounts values('axin','456');");
    //将所有用户信息读出，并输出。
    db.exeSQL("SELECT * from accounts;");
    return 0;
}
