# Movie Booking API

A RESTful API for a movie booking system built with Go, Gin, and PostgreSQL.

## Features

- User authentication with JWT
- GIN and GORM framework implementation
- Register and login
- Movie CRUD management
- Schedule CRUD Management
- Booking CRUD Management (still under construction)
- edit and delete customer data (still under construction)

## Prerequisites

- Go 1.16 or higher
- PostgreSQL 12 or higher
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/NaufalNurFahriza/booking_system_api.git
```

2. Install dependencies:
```bash
go mod tidy
```

3. Set up environment variables in `config/.env`:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=movie_booking_db
JWT_SECRET=your_jwt_secret_key
API_PORT=8080
```

4. Run the application:
```bash
go run main.go
```

## API Endpoints

### Authentication
- `POST /api/register` - Register new customer
- `POST /api/login` - Login customer

### Customer Profile management (still under construction)
- `PUT /api/customers/:id` - Update profile costumer
- `DELETE /api/customers/:id` - Delete profile costumer

### Booking routes (still under construction)
- `POST /api/bookings` - Create new booking
- `PUT /api/bookings` - update booking
- `DELETE /api/bookings/:id` - Delete booking

### admin routes (requires authentication)
### get all data Customer 
- `GET /api/customers/` - List customer's bookings

### Movie management
- `POST /api/movies` - Create movie
- `PUT /api/movies/:id` - Update movie
- `DELETE /api/movies/:id` - Delete movie

### Schedule management
- `POST /api/schedules` - Create schedule
- `PUT /api/schedules/:id` - Update schedule
- `DELETE /api/schedules/:id` - Delete schedule

### Booking management
- `GET /api/bookings` - Create new booking

## Deployment

To deploy to Railway:

1. Install the Railway CLI
2. Login to Railway:
```bash
railway login
```

3. Initialize your project:
```bash
railway init
```

4. Add your environment variables in the Railway dashboard

5. Deploy your application:
```bash
railway up
```

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request
