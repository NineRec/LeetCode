#
#  CombineTwoTables.sql
#  LeetCode
#
#  Created by gongshang on 16/01/21.
#  Copyright (c) 2016年 ninerec. All rights reserved.
#

Select p.FirstName, p.LastName, a.City, a.State
From Person p left join Address a on p.PersonId=a.PersonId;