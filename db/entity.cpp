#include <iostream>
#include <string>
#include "entity.h"

using namespace std;

Product::Product()
{

}

Product::Product(string productName, int productNumber)
{
        this->productNumber = productNumber;
        this->productName = productName;
}

int Product::getProductId()
{
        return this->productId;
}

string Product::getProductName()
{
        return this->productName;
}

int Product::getProductNumber()
{
        return this->productNumber;
}


SecKill::SecKill()
{

}

SecKill::SecKill(int  productId, time_t startTime, time_t  endTime)
{
        this->productId = productId;
        this->startTime = startTime;
        this->endTime = endTime;
}