#
#  DeletedDuplicatedEmails.sql
#  LeetCode
#
#  Created by gongshang on 16/01/21.
#  Copyright (c) 2016年 ninerec. All rights reserved.
#

Delete p1
From Person as p1, 
    (Select 
        p2.Id,
        if (@prevEmail=p2.EMail, True, False) as Dup_Mail,
        @prevEmail:=p2.Email as prevEmail
    From 
        (Select * From Person Order By Email, Id Asc) as p2,
        (Select @prevEmail:=null) as init
    ) as p3
Where p3.Dup_Mail=True And p1.Id=p3.Id;
