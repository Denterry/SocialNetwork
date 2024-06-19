INSERT INTO users (username, password, role, email)
VALUES ('testuser@example.com', 'password123', 'USER', 'testuser@example.com')
ON CONFLICT (username) DO NOTHING;