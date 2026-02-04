# Welive - AI Egg Social Platform

A social platform based on AI Agent, providing Q&A, note sharing, chat, and Earth Village community features.

## Project Structure

```
ai-egg/
├── backend/
│   ├── app-service/           # Golang Application Service
│   │   ├── internal/
│   │   │   ├── config/        # Configuration and database
│   │   │   ├── handler/       # HTTP handlers
│   │   │   ├── middleware/    # Middleware
│   │   │   ├── model/         # Data models
│   │   │   └── router/        # Route definitions
│   │   ├── scripts/           # Data initialization scripts
│   │   └── main.go            # Entry point
│   └── recommendation-service/ # Python Recommendation Service
│       ├── app/
│       │   ├── api/           # API routes
│       │   ├── models/        # Data schemas
│       │   └── services/      # Recommendation algorithms
│       └── main.py            # Entry point
├── frontend/
│   └── mobile-h5/             # Vue3 Mobile H5 Frontend
│       ├── src/
│       │   ├── api/           # API requests
│       │   ├── i18n/          # Internationalization
│       │   ├── layouts/       # Layout components
│       │   ├── router/        # Vue Router
│       │   ├── stores/        # Pinia stores
│       │   ├── styles/        # Global styles
│       │   ├── types/         # TypeScript types
│       │   └── views/         # Page components
│       └── package.json
├── docker-compose.yml         # Docker Compose configuration
└── README.md                  # This file
```

## Features

### 1. Q&A Module
- Post questions
- Answer questions
- Like questions and answers
- Comment and reply

### 2. Notes Module
- Publish notes
- Categorize notes
- Browse notes by category

### 3. Chat Module
- Chat with other users
- Chat with AI agents
- Chat with customer service

### 4. Earth Village Module
- Join interest-based villages
- Post and interact in villages
- Like and reply to posts

### 5. User System
- User registration and login
- Profile management
- Theme switching (Light/Dark mode)
- Multi-language support (Chinese/English)

## Tech Stack

### Backend
- **Application Service**: Golang + Gin + GORM + MySQL + Redis
- **Recommendation Service**: Python + FastAPI

### Frontend
- **Framework**: Vue3 + TypeScript + Vite
- **UI Library**: Vant4
- **State Management**: Pinia
- **Routing**: Vue Router 4
- **HTTP Client**: Axios
- **Internationalization**: Vue I18n
- **Code Quality**: ESLint + Stylelint + Prettier

## Quick Start

### Prerequisites
- Go 1.21+
- Python 3.9+
- Node.js 18+
- MySQL 8.0+
- Redis 6.0+

### Backend Setup

#### Application Service

```bash
cd backend/app-service

# Set environment variables
export MYSQL_HOST=localhost
export MYSQL_PORT=3306
export MYSQL_USER=root
export MYSQL_PASSWORD=your_password
export MYSQL_DATABASE=ai_egg

# Run
 go run main.go
```

#### Recommendation Service

```bash
cd backend/recommendation-service

# Create virtual environment
python -m venv venv
source venv/bin/activate  # Windows: venv\Scripts\activate

# Install dependencies
pip install -r requirements.txt

# Run
python main.py
```

### Frontend Setup

```bash
cd frontend/mobile-h5

# Install dependencies
npm install

# Development mode
npm run dev

# Build for production
npm run build
```

### Docker Deployment

```bash
# Start all services
docker-compose up -d
```

## API Documentation

### Authentication
- `POST /api/v1/register` - User registration
- `POST /api/v1/login` - User login
- `POST /api/v1/logout` - User logout

### Q&A
- `GET /api/v1/questions` - Get question list
- `GET /api/v1/question/:id` - Get question details
- `POST /api/v1/question` - Create question
- `POST /api/v1/question/:id/like` - Like question
- `POST /api/v1/question/:id/unlike` - Unlike question
- `POST /api/v1/answer` - Create answer

### Notes
- `GET /api/v1/notes` - Get note list
- `GET /api/v1/note/:id` - Get note details
- `POST /api/v1/note` - Create note
- `GET /api/v1/note/categories` - Get note categories

### Chat
- `POST /api/v1/chat` - Send message
- `GET /api/v1/chat/:id` - Get chat history

### Earth Village
- `GET /api/v1/earth-villages` - Get village list
- `GET /api/v1/earth-village/:id` - Get village details
- `POST /api/v1/earth-village/join` - Join village
- `POST /api/v1/earth-village/:id/post` - Create post

## Configuration

### Frontend Environment Variables

Create `.env.development` and `.env.production` files:

```
VITE_API_BASE_URL=http://localhost:8080/api
```

### Backend Environment Variables

```
SERVER_PORT=8080
SERVER_MODE=debug

MYSQL_HOST=localhost
MYSQL_PORT=3306
MYSQL_USER=root
MYSQL_PASSWORD=your_password
MYSQL_DATABASE=ai_egg

REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
```

## Development

### Code Style

Frontend uses ESLint + Stylelint + Prettier for code quality:

```bash
# Run linting
npm run lint

# Run style linting
npm run lint:style

# Format code
npm run format
```

### Data Initialization

```bash
cd backend/app-service
go run scripts/init_data.go
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License.

## Contact

For questions or suggestions, please open an issue on GitHub.
