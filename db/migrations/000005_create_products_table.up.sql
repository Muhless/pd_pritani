DO $$ BEGIN
          CREATE TYPE product_category AS ENUM ('rice','bran');
EXCEPTION
          WHEN duplicate_object THEN NULL;
END $$;

CREATE TABLE IF NOT EXISTS products (
          id BIGSERIAL PRIMARY KEY,
          name VARCHAR(50) NOT NULL,
          category product_category NOT NULL DEFAULT 'rice', 
          stock NUMERIC(12,2) DEFAULT 0,
          price NUMERIC(12,2) NOT NULL,
          unit VARCHAR(20) NOT NULL,
          photo TEXT,
          description TEXT,
          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
          deleted_at TIMESTAMP DEFAULT NULL
)