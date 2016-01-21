#
#  DepartmentHighestSalary.sql
#  LeetCode
#
#  Created by gongshang on 16/01/20.
#  Copyright (c) 2016年 ninerec. All rights reserved.
#

Select 
    d.Name as Department, e.Name as Employee, e.Salary as Salary
From 
    Employee as e, Department as d,
    (Select DepartmentId, Max(Salary) as MaxSalary From Employee Group By DepartmentId) as t
Where
    e.DepartmentId = d.Id And e.DepartmentId = t.DepartmentId And e.Salary = t.MaxSalary;
