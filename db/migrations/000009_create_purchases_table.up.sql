DO $$ BEGIN
CREATE TYPE purchase_status AS ENUM ('pending', 'paid', 'cancelled');
EXCEPTION
WHEN duplicate_object THEN NULL;
END $$;

CREATE TABLE IF NOT EXISTS purchases (
     id BIGSERIAL PRIMARY KEY,
     po_number VARCHAR(50) NOT NULL UNIQUE,
     employee_id BIGINT NOT NULL,
     supplier_id BIGINT NOT NULL,
     total_price NUMERIC(12, 2) NOT NULL,
     status purchase_status NOT NULL DEFAULT 'pending',
     notes TEXT,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     deleted_at TIMESTAMP DEFAULT NULL,
     CONSTRAINT fk_purchases_employee FOREIGN KEY (employee_id) REFERENCES employees(id),
     CONSTRAINT fk_purchases_supplier FOREIGN KEY (supplier_id) REFERENCES suppliers(id)
)