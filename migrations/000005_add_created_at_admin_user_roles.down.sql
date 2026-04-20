-- migrations/000005_add_created_at_admin_user_roles.down.sql

ALTER TABLE admin_user_role_groups
DROP COLUMN created_at;
