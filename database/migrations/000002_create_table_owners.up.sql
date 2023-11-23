CREATE TABLE IF NOT EXISTS owners (
  id VARCHAR(32) PRIMARY KEY,
  user_id VARCHAR(32) NOT NULL,
  fullname VARCHAR(100) NOT NULL,
  gender CHAR(8) NOT NULL,
  phone CHAR(15) NOT NULL,
  city VARCHAR(50) NOT NULL,
  address VARCHAR(250) NOT NULL,
  birthdate DATE,
  status CHAR(10),
  photo VARCHAR(255),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id)
);