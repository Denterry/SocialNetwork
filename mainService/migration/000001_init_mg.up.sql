CREATE DATABASE auth_db;
CREATE USER authiniter WITH SUPERUSER PASSWORD 'qwerty123456';
GRANT ALL PRIVILEGES ON DATABASE auth_db TO authiniter;