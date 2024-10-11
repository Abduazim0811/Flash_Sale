CREATE TABLE IF NOT EXISTS products (
    id TEXT,           
    name VARCHAR(255) NOT NULL,  
    description TEXT,                
    price DECIMAL(10, 2) NOT NULL,
    discount_price DECIMAL(10,2),   
    stock INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  
    is_active BOOLEAN DEFAULT TRUE  
);;