CREATE TABLE IF NOT EXISTS transaction_items (
    id uuid PRIMARY KEY  DEFAULT gen_random_uuid(),
    transaction_id uuid NOT NULL,
    FOREIGN KEY (transaction_id) REFERENCES transactions(id),
    product_id uuid NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id),
    quantity INT NOT NULL
);