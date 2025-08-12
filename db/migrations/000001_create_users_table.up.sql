CREATE TABLE
          IF NOT EXISTS users (
                    id INT PRIMARY KEY AUTO_INCREMENT,
                    name VARCHAR(30) NOT NULL,
                    phone VARCHAR(15) NOT NULL UNIQUE,
                    role VARCHAR(15) NOT NULL,
                    username VARCHAR(20) NOT NULL UNIQUE,
                    password VARCHAR(255) NOT NULL,
                    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
          )