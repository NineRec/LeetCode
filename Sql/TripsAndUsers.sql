#
#  TripsAndUsers.sql
#  LeetCode
#
#  Created by gongshang on 16/01/19.
#  Copyright (c) 2016年 ninerec. All rights reserved.
#

Select t1.`Request_at` as `Day`, round(COUNT(t2.`Id`)/COUNT(t1.`Id`), 2) as `Cancellation Rate`
From Users, Trips as t1 left join Trips as t2 on (t1.Id=t2.Id and t2.Status!='completed')
Where (t1.Client_Id=Users.Users_Id and t1.Request_at in ('2013-10-01', '2013-10-02', '2013-10-03')
        and Users.Banned="No" and Users.Role="Client" )
Group By Day;
