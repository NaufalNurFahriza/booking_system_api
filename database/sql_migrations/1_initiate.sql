
-- Tabel Users dengan role
CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'customer', -- 'admin' atau 'customer'
    phone VARCHAR(15),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Vehicles
CREATE TABLE Vehicles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    type VARCHAR(50) NOT NULL,
    license_plate VARCHAR(15) UNIQUE NOT NULL,
    price_per_day DECIMAL(10, 2) NOT NULL CHECK (price_per_day > 0),
    available BOOLEAN DEFAULT TRUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Bookings dengan status
CREATE TABLE Bookings (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    vehicle_id INT NOT NULL,
    booking_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL CHECK (total_price > 0),
    status VARCHAR(20) DEFAULT 'pending', -- pending, confirmed, completed, cancelled
    FOREIGN KEY (user_id) REFERENCES Users (id) ON DELETE CASCADE,
    FOREIGN KEY (vehicle_id) REFERENCES Vehicles (id) ON DELETE CASCADE,
    CHECK (end_date >= start_date)
);

-- Tabel Payments
CREATE TABLE Payments (
    id SERIAL PRIMARY KEY,
    booking_id INT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL CHECK (amount > 0),
    payment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    payment_method VARCHAR(50) NOT NULL, -- cash, credit_card, transfer
    payment_status VARCHAR(20) DEFAULT 'pending', -- pending, completed, failed, refunded
    transaction_id VARCHAR(100),
    notes TEXT,
    FOREIGN KEY (booking_id) REFERENCES Bookings (id) ON DELETE CASCADE
);