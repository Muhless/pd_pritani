CREATE TABLE
          IF NOT EXISTS customers (
                    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                    name VARCHAR(30) NOT NULL,
                    phone VARCHAR(12) NOT NULL UNIQUE,
                    address TEXT,
                    company VARCHAR(30),
                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
          )