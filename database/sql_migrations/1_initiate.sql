
-- Tabel User dengan role
CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL DEFAULT 'customer', -- 'admin' atau 'customer'
    phone VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Movie
CREATE TABLE Movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    genre VARCHAR(255) NOT NULL,
    duration INTEGER NOT NULL,
    description TEXT,
    release_date DATE NOT NULL,
    rating VARCHAR(255) NOT NULL, -- 'PG-13' atau 'R'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Schedule
CREATE TABLE Schedules (
    id SERIAL PRIMARY KEY,
    movie_id INTEGER NOT NULL,
    show_time TIME NOT NULL, --- contoh '15:30:00'
    ticket_price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (movie_id) REFERENCES Movies(id) ON DELETE CASCADE
);

-- Tabel Booking dengan status
CREATE TABLE Bookings (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    schedule_id INTEGER NOT NULL,
    booking_date DATE NOT NULL, --- contoh '2024-12-28'
    total_price DECIMAL(10, 2) NOT NULL CHECK (total_price > 0),
    payment_status VARCHAR(20) DEFAULT 'unpaid', --- unpaid atau paid
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE,
    FOREIGN KEY (schedule_id) REFERENCES Schedules(id) ON DELETE CASCADE,
);