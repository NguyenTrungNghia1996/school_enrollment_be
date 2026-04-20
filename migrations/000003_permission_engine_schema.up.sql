-- migrations/000003_permission_engine_schema.up.sql

ALTER TABLE role_group_permissions
CHANGE COLUMN permission_code permission_key VARCHAR(100) NOT NULL,
ADD COLUMN permission_value BIGINT NOT NULL DEFAULT 0;

-- Drop primary key to change it, depending on exact MySQL definition. If it fails, we will safely drop and recreate.
-- Actually, the easiest is to recreate the primary key if needed, or if only renaming column, it might just work.
-- Let's drop primary key and add again to be safe.
-- ALTER TABLE role_group_permissions DROP PRIMARY KEY, ADD PRIMARY KEY(role_group_id, permission_key);
-- Usually CHANGE COLUMN renames the column in the primary key automatically in MySQL.

ALTER TABLE menus
ADD COLUMN menu_key VARCHAR(100) UNIQUE,
ADD COLUMN permission_bit INT DEFAULT 0;
