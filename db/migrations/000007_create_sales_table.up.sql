CREATE TABLE
          IF NOT EXISTS sales (
                    id BIGSERIAL PRIMARY KEY,
                    customer_id BIGINT NOT NULL,
                    sales_date DATE NOT NULL,
                    total_amount NUMERIC(15, 2),
                    paid_amount NUMERIC(15, 2),
                    remaining_amount NUMERIC(15, 2),
                    status sales_status,
                    note TEXT,
                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    CONSTRAINT fk_sales_customer FOREIGN KEY (customer_id) REFERENCES customers (id) ON DELETE RESTRICT
          )