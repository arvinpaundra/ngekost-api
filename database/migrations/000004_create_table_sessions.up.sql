CREATE TABLE IF NOT EXISTS sessions (
  id VARCHAR(32) PRIMARY KEY,
  user_id VARCHAR(32) NOT NULL,
  device_id VARCHAR(255) NOT NULL,
  device_name VARCHAR(250) NOT NULL,
  ip_address CHAR(15) NOT NULL,
  platform CHAR(8) NOT NULL,
  access_token TEXT NOT NULL,
  refresh_token TEXT,
  fcm_token TEXT,
  google_oauth_token TEXT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id)
);