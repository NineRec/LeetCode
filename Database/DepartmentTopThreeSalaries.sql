#
#  DuplicateEmails.sql
#  LeetCode
#
#  Created by gongshang on 16/01/20.
#  Copyright (c) 2016年 ninerec. All rights reserved.
#
#  Copyed from the intenet

SELECT D.Name AS Department, E.Name AS Employee, E.Salary AS Salary 
FROM Employee E, Department D
WHERE (SELECT COUNT(DISTINCT(Salary)) FROM Employee 
       WHERE DepartmentId = E.DepartmentId AND Salary > E.Salary) < 4
AND E.DepartmentId = D.Id 
ORDER BY E.DepartmentId, E.Salary DESC;