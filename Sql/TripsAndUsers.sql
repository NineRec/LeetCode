#
#  TripsAndUsers.sql
#  LeetCode
#
#  Created by gongshang on 16/01/19.
#  Copyright (c) 2016年 ninerec. All rights reserved.
#

Select 
    t.Request_at as `Day`, 
    round(SUM(t.Status!="completed")/COUNT(t.Id), 2) as `Cancellation Rate`
From Trips as t Inner Join Users as u 
    on (t.Client_Id=u.Users_Id and u.Banned="No")
Where t.Request_at BETWEEN '2013-10-01' AND '2013-10-03'
Group By t.Request_at;