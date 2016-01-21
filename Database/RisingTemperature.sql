#
#  RisingTemprature.sql
#  LeetCode
#
#  Created by gongshang on 16/01/21.
#  Copyright (c) 2016年 ninerec. All rights reserved.
#

-- Select w1.Id
-- From Weather w1 
--     Inner Join Weather w2 
--     On DATE_SUB(w1.`Date`, INTERVAL 1 DAY)=w2.`Date`
-- Where w1.Temperature > w2.Temperature;

Select w1.Id
From ( 
    Select 
        w.Id,
        if (@prevTemp<w.Temperature And @prevDate = Date_Sub(w.`Date`, Interval 1 Day), True, False) as Rising,
        @prevTemp:=w.Temperature,
        @prevDate:=w.`Date`
    From 
        (Select * From Weather Order By `Date` Asc) as w,
        (Select @prevTemp:=null, @prevDate:=null) as init
    ) as w1
Where w1.Rising=True;