-- Hapus objek database yang ada untuk memastikan skrip bisa dijalankan ulang
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS users;
DROP TYPE IF EXISTS user_role;

-- Mengaktifkan ekstensi untuk generate UUID jika belum ada
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Membuat tipe data ENUM untuk peran pengguna
-- Ini memastikan kolom 'role' hanya bisa diisi dengan nilai yang sudah ditentukan.
CREATE TYPE user_role AS ENUM ('admin', 'user');

-- 1. Tabel untuk data pengguna (users)
CREATE TABLE users (
    id UUID     PRIMARY KEY     DEFAULT uuid_generate_v4(),
    full_name   VARCHAR(255)    NOT NULL,
    email       VARCHAR(255)    UNIQUE NOT NULL,            -- Kolom email harus unik
    password    VARCHAR(255)    NOT NULL,                   -- Akan menyimpan password yang sudah di-hash
    role        user_role       NOT NULL DEFAULT 'user',    -- Default peran adalah 'user'
    created_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW()
);

-- 2. Tabel untuk data produk (products)
-- Tabel ini memiliki relasi foreign key ke tabel users
CREATE TABLE products (
    id UUID     PRIMARY KEY     DEFAULT uuid_generate_v4(),
    name        VARCHAR(255)    NOT NULL,
    price       INT             NOT NULL CHECK (price >= 0),    -- Memastikan harga tidak negatif
    user_id     UUID,                                           -- Kolom untuk menyimpan ID pengguna yang membuat produk
    created_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW(),

    -- Mendefinisikan Foreign Key Constraint
    -- Ini menghubungkan products.user_id dengan users.id
    CONSTRAINT fk_user
        FOREIGN KEY(user_id) 
        REFERENCES users(id)
        ON DELETE SET NULL -- Jika user dihapus, user_id di produk ini akan menjadi NULL
);

-- Membuat index pada foreign key untuk mempercepat query join
CREATE INDEX idx_products_user_id ON products(user_id);

-- INSERT DATA SAMPLE (opsional)
-- Password di-hash menggunakan bcrypt untuk keamanan
-- Password 'OnlinePHP' di-hash menjadi '$2y$10$TN7zDKb1jY9Dmmi3JKujWebUXWdcSMQd5Pq5qHjA6jAeUWKECo9tG'
INSERT INTO users (full_name, email, password, role) VALUES
('Admin User', 'h2xkR@example.com', '$2y$10$TN7zDKb1jY9Dmmi3JKujWebUXWdcSMQd5Pq5qHjA6jAeUWKECo9tG', 'admin'),
('User Biasa', 'l7bTg@example.com', '$2y$10$TN7zDKb1jY9Dmmi3JKujWebUXWdcSMQd5Pq5qHjA6jAeUWKECo9tG', 'user');

INSERT INTO products (name, price, user_id) VALUES
('Produk A', 10000, (SELECT id FROM users WHERE email = 'h2xkR@example.com')),
('Produk B', 20000, (SELECT id FROM users WHERE email = 'h2xkR@example.com')),
('Produk C', 15000, (SELECT id FROM users WHERE email = 'l7bTg@example.com'));