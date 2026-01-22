CREATE TABLE
          IF NOT EXISTS admins (
                    id BIGSERIAL PRIMARY KEY,
                    user_id INT NOT NULL UNIQUE,
                    name VARCHAR(50) NOT NULL,
                    email VARCHAR(50) NOT NULL,
                    phone VARCHAR(20) NOT NULL,
                    photo TEXT,
                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    CONSTRAINT fk_admin_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
          )