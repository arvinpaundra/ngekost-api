CREATE TABLE IF NOT EXISTS sessions (
  id CHAR(32) PRIMARY KEY,
  user_id CHAR(32) NOT NULL,
  device_id VARCHAR(255) NOT NULL,
  device_name VARCHAR(250) NOT NULL,
  ip_address VARCHAR(15) NOT NULL,
  platform VARCHAR(8) NOT NULL,
  access_token TEXT NOT NULL,
  refresh_token TEXT,
  fcm_token TEXT,
  google_oauth_token TEXT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id)
);