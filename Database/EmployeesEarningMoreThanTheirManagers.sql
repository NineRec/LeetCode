#
#  EmployeesEarningMoreThanTheirManagers.sql
#  LeetCode
#
#  Created by gongshang on 16/01/21.
#  Copyright (c) 2016年 ninerec. All rights reserved.
#

Select e1.Name as Employee
From Employee e1 
    Inner Join Employee e2 
    On e1.ManagerId=e2.Id
Where e1.Salary > e2.Salary;