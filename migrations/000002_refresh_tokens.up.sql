CREATE TABLE IF NOT EXISTS refresh_tokens(
    user_id UUID REFERENCES users(id) ON DELETE CASCADE, 
    refesh_token VARCHAR(255) NOT NULL
);