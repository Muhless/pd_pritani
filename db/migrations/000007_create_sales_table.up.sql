CREATE TABLE
          IF NOT EXISTS sales (
                    id BIGSERIAL PRIMARY KEY,
                    sales_date DATE NOT NULL,
                    total_amount NUMERIC(15, 2),
                    paid_amount NUMERIC(15, 2),
                    remaining_amount NUMERIC(15, 2),
                    status sales_status,
                    note TEXT,
                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
          )