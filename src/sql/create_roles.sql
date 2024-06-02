CREATE ROLE user_role;
GRANT INSERT ON TABLE term_annotation      TO user_role;
GRANT INSERT ON TABLE term_annotation_part TO user_role;
GRANT INSERT ON TABLE struct_annotation    TO user_role;
GRANT SELECT ON TABLE annotation_task      TO user_role;
GRANT SELECT ON TABLE sentence             TO user_role;
GRANT SELECT ON TABLE token                TO user_role;

CREATE ROLE moderator_role;
GRANT user_role                         TO moderator_role;
GRANT INSERT ON TABLE annotation_task   TO moderator_role;
GRANT UPDATE ON TABLE annotation_task   TO moderator_role;
GRANT UPDATE ON TABLE term_annotation   TO moderator_role;
GRANT UPDATE ON TABLE struct_annotation TO moderator_role;

CREATE ROLE admin_role WITH SUPERUSER;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO admin_role;

CREATE USER dbusr WITH PASSWORD 'dbusr';
CREATE USER moder WITH PASSWORD 'moder';
CREATE USER admin WITH PASSWORD 'admin';

GRANT user_role      TO dbusr;
GRANT moderator_role TO moder;
GRANT admin_role     TO admin;
