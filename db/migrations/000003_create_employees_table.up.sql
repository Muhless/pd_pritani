DO $$ BEGIN
CREATE TYPE employee_status AS ENUM ('active', 'inactive');
EXCEPTION
WHEN duplicate_object THEN NULL;
END $$;

CREATE TABLE
          IF NOT EXISTS employees (
                    id BIGSERIAL PRIMARY KEY,
                    user_id BIGINT NOT NULL UNIQUE,
                    name VARCHAR(30) NOT NULL,
                    phone VARCHAR(15) NOT NULL,
                    address TEXT,
                    photo TEXT,
                    status employee_status NOT NULL DEFAULT 'active',
                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    deleted_at TIMESTAMP DEFAULT NULL,
                    CONSTRAINT fk_employee_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
          )