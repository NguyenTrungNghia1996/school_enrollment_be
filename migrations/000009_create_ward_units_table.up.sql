-- migrations/000009_create_ward_units_table.up.sql

CREATE TABLE IF NOT EXISTS ward_units (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    province_id INT UNSIGNED NOT NULL,
    code VARCHAR(50) NOT NULL,
    name VARCHAR(255) NOT NULL,
    unit_type ENUM('Ward', 'Commune', 'SpecialZone') NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL,
    UNIQUE INDEX idx_province_name (province_id, name),
    INDEX idx_province (province_id),
    CONSTRAINT fk_ward_units_province FOREIGN KEY (province_id) REFERENCES provinces(id) ON DELETE CASCADE
);
