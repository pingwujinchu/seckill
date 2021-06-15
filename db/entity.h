
#ifndef _ENTITY_H
#define _ENTITY_H
#include<iostream>
#include<string>
#include <ctime>

using namespace std;

class  Product
{
    public:
    Product();
    Product(string productName, int productNumber);
    ~Product();
    bool initDB(string host,string user,string pwd,string db_name); //连接mysql
    bool exeSQL(string sql);   //执行sql语句
    int getProductId();
    string getProductName();
    int getProductNumber();
    private:
    int productId;
    string productName; 
    int productNumber;
};


class  SecKill 
{
    public:
    SecKill();
    SecKill(int  productId, time_t startTime, time_t  endTime);
    ~SecKill();
    SecKill(time_t startTime, time_t endTime);
    private:
    int  secKillId;
    int  productId;
    time_t startTime;
    time_t endTime;
};

#endif