CREATE TABLE IF NOT EXISTS payments (
  id CHAR(32) PRIMARY KEY,
  bill_id CHAR(32) NOT NULL,
  nominal NUMERIC(11, 2) NOT NULL,
  method VARCHAR(10) NOT NULL,
  status VARCHAR(15) NOT NULL,
  via VARCHAR(10) NOT NULL,
  raw JSON,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (bill_id) REFERENCES bills(id)
)