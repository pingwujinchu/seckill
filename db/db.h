/*************************************************************************
    > File Name: mydb.h
    > Author:fengxin 
    > Mail:903087053@qq.com 
    > Created Time: 2017年07月21日 星期五 15时17分17秒
 ************************************************************************/

#ifndef _MYDB_H
#define _MYDB_H
#include<iostream>
#include<string>
#include<mysql/mysql.h>
#include "entity.h"
using namespace std;

class DB
{
    public:
    DB();
    ~DB();
    bool initDB(string host,string user,string pwd,string db_name); //连接mysql
    bool exeSQL(string sql);   //执行sql语句
    bool  addProduct(Product product);
    bool  addSecKill(SecKill SecKill);
    private:
    MYSQL *mysql;          //连接mysql句柄指针
    MYSQL_RES *result;    //指向查询结果的指针
    MYSQL_ROW row;       //按行返回的查询信息
};


#endif