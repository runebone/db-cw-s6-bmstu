db=> select username, trust_level, status from user_data u join struct_annotation s on u.id = s.done_by where s.id = 'b72d276e-a6ca-402b-bcdb-480da5acb9dd';
   username   | trust_level | status
--------------+-------------+--------
 rgallegos988 |           0 |

db=> update struct_annotation set status = 'approved' where id = 'b72d276e-a6ca-402b-bcdb-480da5acb9dd';
UPDATE 1
db=> select username, trust_level, status from user_data u join struct_annotation s on u.id = s.done_by where s.id = 'b72d276e-a6ca-402b-bcdb-480da5acb9dd';
   username   | trust_level |  status
--------------+-------------+----------
 rgallegos988 |           5 | approved
