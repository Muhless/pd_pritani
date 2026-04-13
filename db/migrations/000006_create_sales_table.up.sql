DO $$ BEGIN
          CREATE TYPE sales_status AS ENUM ('pending','paid','cancelled');
EXCEPTION
          WHEN duplicate_object THEN NULL;
END $$;

CREATE TABLE IF NOT EXISTS sales (
          id BIGSERIAL PRIMARY KEY,
          invoice_number VARCHAR(50) NOT NULL UNIQUE,
          employee_id BIGINT NOT NULL,
          customer_id BIGINT NOT NULL,
          total_price NUMERIC(15,2),
          status sales_status NOT NULL DEFAULT 'pending',
          notes TEXT,
          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
          deleted_at TIMESTAMP DEFAULT NULL,

          CONSTRAINT fk_sales_customer FOREIGN KEY (customer_id) REFERENCES customers(id),
          CONSTRAINT fk_sales_employee FOREIGN KEY (employee_id) REFERENCES employees(id)
)