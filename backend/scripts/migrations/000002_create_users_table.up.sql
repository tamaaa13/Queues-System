CREATE TABLE IF NOT EXISTS users (
  id INT AUTO_INCREMENT,
  username VARCHAR(50) NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
  role ENUM('admin', 'superadmin') NOT NULL,
  category ENUM('A', 'B') DEFAULT NULL, -- hanya untuk admin biasa
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  PRIMARY KEY(id),
  UNIQUE KEY unique_username (username)
);