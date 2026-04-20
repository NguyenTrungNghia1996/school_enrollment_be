CREATE TABLE IF NOT EXISTS application_exam_scores (
    id INT AUTO_INCREMENT PRIMARY KEY,
    application_id INT NOT NULL,
    subject_id INT NOT NULL,
    raw_score DECIMAL(5,2),
    bonus_score DECIMAL(5,2) DEFAULT 0.00,
    final_score DECIMAL(5,2),
    is_absent BOOLEAN DEFAULT FALSE,
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (application_id) REFERENCES applications(id) ON DELETE CASCADE,
    FOREIGN KEY (subject_id) REFERENCES subjects(id) ON DELETE CASCADE,
    UNIQUE KEY idx_app_subject_score (application_id, subject_id)
);
