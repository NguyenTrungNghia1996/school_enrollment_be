-- migrations/000003_permission_engine_schema.down.sql

ALTER TABLE role_group_permissions
DROP COLUMN permission_value,
CHANGE COLUMN permission_key permission_code VARCHAR(100) NOT NULL;

ALTER TABLE menus
DROP COLUMN menu_key,
DROP COLUMN permission_bit;
