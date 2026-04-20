-- migrations/000001_init_schema.down.sql
-- Enrollment System Clean Up Migration

DROP TABLE IF EXISTS application_results;
DROP TABLE IF EXISTS application_exam_scores;
DROP TABLE IF EXISTS admission_period_subjects;
DROP TABLE IF EXISTS subjects;
DROP TABLE IF EXISTS examiner_assignments;
DROP TABLE IF EXISTS exam_room_assignments;
DROP TABLE IF EXISTS examiners;
DROP TABLE IF EXISTS exam_rooms;
DROP TABLE IF EXISTS application_documents;
DROP TABLE IF EXISTS academic_records;
DROP TABLE IF EXISTS applications;
DROP TABLE IF EXISTS admission_periods;
DROP TABLE IF EXISTS menus;
DROP TABLE IF EXISTS role_group_permissions;
DROP TABLE IF EXISTS admin_user_role_groups;
DROP TABLE IF EXISTS role_groups;
DROP TABLE IF EXISTS admin_users;
DROP TABLE IF EXISTS user_accounts;
DROP TABLE IF EXISTS ward_units;
DROP TABLE IF EXISTS provinces;
