CREATE TABLE IF NOT EXISTS examiner_assignments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    examiner_id INT NOT NULL,
    exam_room_id INT NOT NULL,
    role ENUM('Primary', 'Secondary', 'Backup') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (examiner_id) REFERENCES examiners(id) ON DELETE CASCADE,
    FOREIGN KEY (exam_room_id) REFERENCES exam_rooms(id) ON DELETE CASCADE,
    UNIQUE KEY idx_examiner_room_role (examiner_id, exam_room_id, role)
);
