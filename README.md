# Movie Booking API

A RESTful API for a movie booking system built with Go, Gin, and PostgreSQL.

## Features

- User authentication with JWT
- Movie management
- Cinema and screen management
- Schedule management
- Booking system with seat availability checking
- Admin panel for content management

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

### Public Endpoints
- `GET /api/movies` - List all movies
- `GET /api/movies/:id` - Get movie details
- `GET /api/cinemas` - List all cinemas
- `GET /api/cinemas/:id` - Get cinema details
- `GET /api/schedules` - List all schedules
- `GET /api/schedules/:id` - Get schedule details

### Protected Endpoints (requires authentication)
- `POST /api/bookings` - Create new booking
- `GET /api/bookings` - List customer's bookings
- `GET /api/bookings/:id` - Get booking details
- `PUT /api/bookings/:id/cancel` - Cancel booking

### Admin Endpoints (requires authentication)
- `POST /api/admin/movies` - Create movie
- `PUT /api/admin/movies/:id` - Update movie
- `DELETE /api/admin/movies/:id` - Delete movie
- `POST /api/admin/cinemas` - Create cinema
- `PUT /api/admin/cinemas/:id` - Update cinema
- `DELETE /api/admin/cinemas/:id` - Delete cinema
- `POST /api/admin/screens` - Create screen
- `GET /api/admin/screens` - List all screens
- `GET /api/admin/screens/:id` - Get screen details
- `PUT /api/admin/screens/:id` - Update screen
- `DELETE /api/admin/screens/:id` - Delete screen
- `POST /api/admin/schedules` - Create schedule
- `PUT /api/admin/schedules/:id` - Update schedule
- `DELETE /api/admin/schedules/:id` - Delete schedule

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
