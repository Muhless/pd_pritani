CREATE TYPE employee_status AS ENUM ('active', 'inactive');

CREATE TABLE
          IF NOT EXISTS employees (
                    id BIGSERIAL PRIMARY KEY,
                    user_id INT NOT NULL UNIQUE,
                    name VARCHAR(30) NOT NULL,
                    phone VARCHAR(12) NOT NULL,
                    address VARCHAR(30) NOT NULL,
                    photo TEXT,
                    status employee_status NOT NULL DEFAULT 'active',
                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    CONSTRAINT fk_employee_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
          )