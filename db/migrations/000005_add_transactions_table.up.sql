CREATE TABLE IF NOT EXISTS transactions (
    id uuid PRIMARY KEY  DEFAULT gen_random_uuid(),
    customer_id uuid NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers(id),
    paid INT NOT NULL,
    change INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);