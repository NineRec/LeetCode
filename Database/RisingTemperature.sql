#
#  RisingTemprature.sql
#  LeetCode
#
#  Created by gongshang on 16/01/21.
#  Copyright (c) 2016年 ninerec. All rights reserved.
#

Select w1.Id
From Weather w1 
    Inner Join Weather w2 
    On DATE_SUB(w1.`Date`, INTERVAL 1 DAY)=w2.`Date`
Where w1.Temperature > w2.Temperature;