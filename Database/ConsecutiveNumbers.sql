#
#  ConsecutiveNumbers.sql
#  LeetCode
#
#  Created by gongshang on 16/01/21.
#  Copyright (c) 2016年 ninerec. All rights reserved.
#

Select Distinct t.Num
From (
    Select 
        Num,
        @count := if (@prev = Num, @count + 1, 1) as `Total`,
        @prev := Num as `Prev`
    From 
        Logs,
        (Select @prev:=Null, @count:=Null) as init 
    ) as t
Where t.Total>=3;
