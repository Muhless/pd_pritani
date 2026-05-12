CREATE TABLE IF NOT EXISTS purchase_items (
     id BIGSERIAL PRIMARY KEY,
     purchase_id BIGINT NOT NULL,
     product_id BIGINT NOT NULL,
     quantity NUMERIC(12, 2) NOT NULL,
     price NUMERIC(12, 2) NOT NULL,
     subtotal NUMERIC(12, 2) NOT NULL,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     deleted_at TIMESTAMP DEFAULT NULL,
     CONSTRAINT fk_purchase_items_purchase FOREIGN KEY (purchase_id) REFERENCES purchases(id),
     CONSTRAINT fk_purchase_items_product FOREIGN KEY (product_id) REFERENCES products(id)
)