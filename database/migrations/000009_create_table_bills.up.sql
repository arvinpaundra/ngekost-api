CREATE TABLE IF NOT EXISTS bills (
  id CHAR(32) PRIMARY KEY,
  lessee_id CHAR(32) NOT NULL,
  invoice VARCHAR(20) NOT NULL,
  kost_name VARCHAR(250) NOT NULL
  room_name VARCHAR(250) NOT NULL
  rent_type VARCHAR(8) NOT NULL,
  status VARCHAR(8) NOT NULL,
  price NUMERIC(11, 2) NOT NULL,
  service_fee NUMERIC(11, 2) NOT NULL,
  total NUMERIC(11, 2) NOT NULL,
  due_date DATE NOT NULL,
  payment_date DATE,
  discount NUMERIC(3, 2),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (lessee_id) REFERENCES lessees(id)
);