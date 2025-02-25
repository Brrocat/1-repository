CREATE TABLE tasks (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       description TEXT,
                       message TEXT,
                       is_done BOOLEAN DEFAULT false,
                       created_at TIMESTAMP,
                       updated_at TIMESTAMP,
                       deleted_at TIMESTAMP
);