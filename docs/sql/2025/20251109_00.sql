DROP TABLE IF EXISTS Reports;

CREATE TABLE
    IF NOT EXISTS reports (
        `id` INT AUTO_INCREMENT PRIMARY KEY,
        `title` VARCHAR(255) NOT NULL,
        `description` TEXT NOT NULL,
        `category` ENUM ('Infrastruktur', 'Pelayanan', 'Keamanan') NOT NULL,
        `location` VARCHAR(255) NOT NULL,
        `photo_url` VARCHAR(255),
        `ticket_code` VARCHAR(255) NOT NULL,
        `status` ENUM ('New', 'In Review', 'Rejected', 'Resolved') DEFAULT 'New',
        `status_desc` TEXT,
        `status_proof_url` VARCHAR(255),
        `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `created_by` VARCHAR(255),
        `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );