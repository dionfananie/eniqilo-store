CREATE TABLE IF NOT EXISTS products (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50) NOT NULL,
    sku VARCHAR(30) NOT NULL,
    category VARCHAR(20) NOT NULL, 
    image_url VARCHAR(255) NOT NULL,
    notes VARCHAR(200) NOT NULL,
    price INT NOT NULL,
    stock INT NOT NULL,
    location VARCHAR(200) NOT NULL,
    is_available BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);