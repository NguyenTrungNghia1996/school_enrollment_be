-- migrations/000002_update_admin_users_schema.up.sql

ALTER TABLE admin_users 
ADD COLUMN username VARCHAR(100) NOT NULL UNIQUE AFTER id,
ADD COLUMN full_name VARCHAR(100) NOT NULL AFTER password_hash,
ADD COLUMN phone_number VARCHAR(20) AFTER email,
ADD COLUMN is_super_admin BOOLEAN DEFAULT FALSE AFTER phone_number,
ADD COLUMN is_active BOOLEAN DEFAULT TRUE AFTER is_super_admin;

ALTER TABLE admin_users DROP COLUMN name;
ALTER TABLE admin_users DROP COLUMN status;
