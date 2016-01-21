#
#  DuplicateEmails.sql
#  LeetCode
#
#  Created by gongshang on 16/01/20.
#  Copyright (c) 2016年 ninerec. All rights reserved.
#

Select Email 
From Person 
Group By email 
Having count(1) > 1;
