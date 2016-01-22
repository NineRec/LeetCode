#
#  DepartmentTopThreeSalaries.sql
#  LeetCode
#
#  Created by gongshang on 16/01/20.
#  Copyright (c) 2016年 ninerec. All rights reserved.
#

-- Copyed from the intenet
SELECT D.Name AS Department, E.Name AS Employee, E.Salary AS Salary 
FROM Employee E, Department D
WHERE (SELECT COUNT(DISTINCT(Salary)) FROM Employee 
       WHERE DepartmentId = E.DepartmentId AND Salary > E.Salary) < 4
AND E.DepartmentId = D.Id 
ORDER BY E.DepartmentId, E.Salary DESC;

-- Result With Argument
Select 
    D.Name as Department, T.Name as Employee, T.salary as Salary
From
    Department as D, (
        Select 
            E.Name, E.Salary, E.DepartmentId,
            @rank:=if (@prevDept=E.DepartmentId, 
                if(@prevSalary=E.Salary, @rank, @rank+1),
                1) as Rank,
            @prevDept:=E.DepartmentId, @prevSalary:=E.Salary
        From
            (Select * From Employee Order By DepartmentId Asc, Salary Desc) E,
            (Select @prevDept:=Null,@prevSalary:=Null,@rank:=Null) init
    ) as T
Where T.Rank <= 3 And D.Id=T.DepartmentId;
