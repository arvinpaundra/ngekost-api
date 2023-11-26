CREATE TABLE IF NOT EXISTS owners (
  id CHAR(32) PRIMARY KEY,
  user_id CHAR(32) NOT NULL,
  fullname VARCHAR(100) NOT NULL,
  gender VARCHAR(8) NOT NULL,
  phone VARCHAR(15) NOT NULL,
  city VARCHAR(50) NOT NULL,
  address VARCHAR(250) NOT NULL,
  birthdate DATE,
  status VARCHAR(10),
  photo VARCHAR(255),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id)
);