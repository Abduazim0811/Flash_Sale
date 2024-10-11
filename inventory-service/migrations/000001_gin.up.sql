CREATE TABLE IF NOT EXISTS inventory(
    id TEXT NOT NULL,               
    product_id TEXT NOT NULL,              
    quantity INT NOT NULL,                 
    status VARCHAR(50) NOT NULL DEFAULT 'available', 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    FOREIGN KEY (product_id) REFERENCES products(id)
);
