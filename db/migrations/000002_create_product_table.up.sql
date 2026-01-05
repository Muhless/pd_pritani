CREATE TABLE
          IF NOT EXISTS products (
                    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                    name VARCHAR(30) NOT NULL,
                    photo TEXT,
                    type VARCHAR(30) NOT NULL,
                    stock INT NOT NULL,
                    price NUMERIC(12, 2) NOT NULL,
                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
          )