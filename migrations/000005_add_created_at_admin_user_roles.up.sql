-- migrations/000005_add_created_at_admin_user_roles.up.sql

ALTER TABLE admin_user_role_groups
ADD COLUMN created_at DATETIME(3) NULL;
