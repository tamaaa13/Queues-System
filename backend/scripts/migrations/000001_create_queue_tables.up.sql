CREATE TABLE IF NOT EXISTS queue (
  id INT NOT NULL AUTO_INCREMENT,
  category ENUM('A', 'B') NOT NULL,
  number VARCHAR(10) NOT NULL,
  status ENUM('waiting', 'called', 'done') DEFAULT 'waiting',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  queue_date DATE GENERATED ALWAYS AS (DATE(created_at)) STORED, 

  PRIMARY KEY(id),
  UNIQUE KEY unique_number_per_day (category, number, queue_date)
);