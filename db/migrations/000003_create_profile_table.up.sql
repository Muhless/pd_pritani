CREATE TABLE
          IF NOT EXISTS profiles (
                    id GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                    name VARCHAR(30) NOT NULL,
                    phone VARCHAR(15) NOT NULL UNIQUE,
                    photo TEXT,
                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    user_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE
          )