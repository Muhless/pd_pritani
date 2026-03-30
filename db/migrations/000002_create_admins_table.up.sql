DO $$ BEGIN
    CREATE TYPE admin_status AS ENUM ('active', 'inactive');
EXCEPTION
    WHEN duplicate_object THEN NULL;
END $$;

CREATE TABLE
          IF NOT EXISTS admins (
                    id BIGSERIAL PRIMARY KEY,
                    user_id BIGINT NOT NULL UNIQUE,
                    permissions VARCHAR(255),
                    name VARCHAR(50) NOT NULL,
                    email VARCHAR(50) NOT NULL,
                    phone VARCHAR(15) NOT NULL,
                    address TEXT,
                    photo TEXT,
                    status admin_status DEFAULT 'active',
                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    deleted_at TIMESTAMP NULL,
                    CONSTRAINT fk_admin_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
          )