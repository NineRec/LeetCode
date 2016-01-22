#
#  DepartmentHighestSalary.sql
#  LeetCode
#
#  Created by gongshang on 16/01/22.
#  Copyright (c) 2016年 ninerec. All rights reserved.
#

Select
    s.Score,
    @rank := @rank + (@prev <> (@prev := Score)) Rank
From 
    (Select * From Scores Order By Score Desc) as s,
    (Select @prev:=-1, @Rank:=0) as init;
