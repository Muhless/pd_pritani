CREATE TABLE
          IF NOT EXISTS profiles (
                    name VARCHAR(30) NOT NULL,
                    phone VARCHAR(15) NOT NULL UNIQUE,
                    photo VARCHAR(255),
                    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                    user_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE
          )