CREATE TABLE IF NOT EXISTS sales_items(
          id BIGSERIAL PRIMARY KEY,
          sales_id BIGINT NOT NULL,
          product_id BIGINT NOT NULL,
          quantity NUMERIC(12,2),
          price NUMERIC(12,2),
          subtotal NUMERIC(12,2),
          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
          deleted_at TIMESTAMP DEFAULT NULL,

          CONSTRAINT fk_sales_items_sales FOREIGN KEY (sales_id) REFERENCES sales(id),
          CONSTRAINT fk_sales_items_product FOREIGN KEY (product_id) REFERENCES products(id)
)