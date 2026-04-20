CREATE TABLE IF NOT EXISTS application_results (
    id INT AUTO_INCREMENT PRIMARY KEY,
    application_id INT NOT NULL UNIQUE,
    total_score DECIMAL(5,2) DEFAULT 0.00,
    priority_score DECIMAL(5,2) DEFAULT 0.00,
    additional_score DECIMAL(5,2) DEFAULT 0.00,
    final_total_score DECIMAL(5,2) DEFAULT 0.00,
    ranking INT,
    result_status ENUM('Pending', 'Passed', 'Failed', 'Waitlisted') DEFAULT 'Pending',
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (application_id) REFERENCES applications(id) ON DELETE CASCADE
);
