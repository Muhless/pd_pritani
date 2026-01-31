CREATE TABLE
          IF NOT EXISTS sales_items (
                    id BIGSERIAL PRIMARY KEY,
                    sales_id BIGINT NOT NULL,
                    product_id BIGINT NOT NULL,
                    quantity INT NOT NULL,
                    subtotal NUMERIC(15, 2),
                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    CONSTRAINT fk_sales_item_sales FOREIGN KEY (sales_id) REFERENCES sales (id) ON DELETE CASCADE,
                    CONSTRAINT fk_sales_item_product FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE RESTRICT,
                    CONSTRAINT unique_sales_product UNIQUE (sales_id, product_id)
          )