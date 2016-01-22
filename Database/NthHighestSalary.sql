#
#  NthHighestSalary.sql
#  LeetCode
#
#  Created by gongshang on 16/01/22.
#  Copyright (c) 2016年 ninerec. All rights reserved.
#

CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  Declare M Int;
  Set M = N - 1;

  RETURN (
    Select Distinct Salary From Employee Order By Salary Desc Limit M,1
  );
END
