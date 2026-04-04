CREATE TABLE IF NOT EXISTS customers (
          id BIGSERIAL PRIMARY KEY,
          name VARCHAR(50) NOT NULL,
          company_name VARCHAR(50),
          phone VARCHAR(15) NOT NULL UNIQUE,
          email VARCHAR(30) UNIQUE,
          address TEXT,
          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
          deleted_at TIMESTAMP DEFAULT NULL
)