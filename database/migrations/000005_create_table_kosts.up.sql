CREATE TABLE IF NOT EXISTS kosts (
  id CHAR(32) PRIMARY KEY,
  owner_id CHAR(32) NOT NULL,
  name VARCHAR(250) NOT NULL,
  type VARCHAR(8) NOT NULL,
  description TEXT,
  payment_interval VARCHAR(10) NOT NULL,
  province VARCHAR(250) NOT NULL,
  city VARCHAR(250) NOT NULL,
  district VARCHAR(250) NOT NULL,
  subdistrict VARCHAR(250) NOT NULL,
  latitude NUMERIC(14, 11) NOT NULL,
  longitude NUMERIC(14, 11) NOT NULL,
  image VARCHAR(255),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (owner_id) REFERENCES owners(id)
);
