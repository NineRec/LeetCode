#
#  CustomersWhoNeverOrder.sql
#  LeetCode
#
#  Created by gongshang on 16/01/21.
#  Copyright (c) 2016年 ninerec. All rights reserved.
#

Select c.Name
From Customers as c left join Orders as o On c.Id=o.CustomerId
Where o.Id Is Null;
