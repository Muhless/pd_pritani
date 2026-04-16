CREATE TABLE IF NOT EXISTS supliers (
          id BIGSERIAL PRIMARY KEY,
          name VARCHAR(50) NOT NULL,
          phone VARCHAR(15) NOT NULL,
          address TEXT,
          notes TEXT,
          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
          deleted_at TIMESTAMP DEFAULT NULL
)