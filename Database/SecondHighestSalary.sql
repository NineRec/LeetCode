#
#  SecondHighestSalary.sql
#  LeetCode
#
#  Created by gongshang on 16/01/21.
#  Copyright (c) 2016年 ninerec. All rights reserved.
#

Select max(Salary) as SecondHighestSalary
From Employee 
Where Salary < (Select max(Salary) From Employee);
