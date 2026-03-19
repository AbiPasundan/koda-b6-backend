CREATE TABLE IF NOT EXISTS forgot_password (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INT NOT NULL,
    token TEXT,
    created_at TIMESTAMP,
    expires_at TIMESTAMP,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
