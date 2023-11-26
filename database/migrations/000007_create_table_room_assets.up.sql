CREATE TABLE IF NOT EXISTS room_assets (
  id CHAR(32) PRIMARY KEY,
  room_id CHAR(32) NOT NULL,
  url VARCHAR(255) NOT NULL,
  type VARCHAR(8) NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (room_id) REFERENCES rooms(id)
);