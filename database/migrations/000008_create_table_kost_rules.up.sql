CREATE TABLE IF NOT EXISTS kost_rules (
  id CHAR(32) PRIMARY KEY,
  kost_id CHAR(32) NOT NULL,
  title VARCHAR(100) NOT NULL,
  description VARCHAR(255),
  priority VARCHAR(8) NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (kost_id) REFERENCES kosts(id)
);