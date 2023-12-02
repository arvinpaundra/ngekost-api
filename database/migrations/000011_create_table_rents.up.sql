CREATE TABLE IF NOT EXISTS rents (
  id CHAR(32) PRIMARY KEY,
  room_id CHAR(32) NOT NULL,
  lessee_id CHAR(32) NOT NULL,
  start_date DATE NOT NULL,
  end_date DATE,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (room_id) REFERENCES rooms(id),
  FOREIGN KEY (lessee_id) REFERENCES lessees(id)
);