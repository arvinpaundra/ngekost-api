CREATE TABLE IF NOT EXISTS rooms (
  id CHAR(32) PRIMARY KEY,
  kost_id CHAR(32) NOT NULL,
  name VARCHAR(250) NOT NULL,
  quantity SMALLINT NOT NULL,
  price NUMERIC(11, 2) NOT NULL,
  category VARCHAR(50),
  description TEXT,
  image VARCHAR(255),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (kost_id) REFERENCES kosts(id)
);