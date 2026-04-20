-- migrations/000002_update_admin_users_schema.down.sql

ALTER TABLE admin_users
ADD COLUMN name VARCHAR(100) NOT NULL AFTER email,
ADD COLUMN status ENUM('active', 'inactive') DEFAULT 'active' AFTER password_hash;

ALTER TABLE admin_users
DROP COLUMN username,
DROP COLUMN full_name,
DROP COLUMN phone_number,
DROP COLUMN is_super_admin,
DROP COLUMN is_active;
