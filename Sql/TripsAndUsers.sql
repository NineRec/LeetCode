#
#  TripsAndUsers.sql
#  LeetCode
#
#  Created by gongshang on 16/01/19.
#  Copyright (c) 2016年 ninerec. All rights reserved.
#

Select t.Request_at as `Day`, round(SUM(t.Status!="completed")/COUNT(t.Id), 2) as `Cancellation Rate`
From Users, Trips as t
Where (t.Client_Id=Users.Users_Id and (t.Request_at BETWEEN '2013-10-01' AND '2013-10-03')
        and Users.Banned="No" and Users.Role="Client" )
Group By Day;