db=> select username, trust_level, status from user_data u join struct_annotation s on u.id = s.done_by where s.id = '595c15b7-a289-4f3c-8069-4c678baaa8ad';
     username      | trust_level | status
-------------------+-------------+--------
 christopher097599 |           0 |

db=> update struct_annotation set status = 'rejected' where id = '595c15b7-a289-4f3c-8069-4c678baaa8ad';
UPDATE 1
db=> select username, trust_level, status from user_data u join struct_annotation s on u.id = s.done_by where s.id = '595c15b7-a289-4f3c-8069-4c678baaa8ad';
     username      | trust_level |  status
-------------------+-------------+----------
 christopher097599 |          -4 | rejected
