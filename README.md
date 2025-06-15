# TaskMorph Backend 🚀

**AI-Enhanced Task Management System**

TaskMorph is a powerful backend API that transforms your big tasks into manageable steps using AI. Simply provide a task description, and our Gemini AI integration will break it down into actionable steps.

## 🧠 Features

- **AI-Powered Task Breakdown**: Uses Google's Gemini AI to intelligently break down complex tasks
- **User Authentication**: Secure JWT-based authentication system
- **Task Management**: Full CRUD operations for tasks and steps
- **Progress Tracking**: Real-time progress calculation based on completed steps
- **RESTful API**: Clean, intuitive API endpoints
- **MongoDB Integration**: Scalable NoSQL database storage

## 🛠️ Tech Stack

- **Backend**: Go (Golang) with Gin framework
- **Database**: MongoDB
- **Authentication**: JWT (JSON Web Tokens)
- **AI Integration**: Google Gemini API
- **Architecture**: Clean Architecture pattern

## 📁 Project Structure

```
taskmorph-backend/
├── main.go                 # Application entry point
├── config/
│   └── mongo.go            # MongoDB connection configuration
├── controllers/
│   ├── auth.go             # Authentication handlers
│   └── task.go             # Task management handlers
├── middleware/
│   ├── auth.go             # JWT authentication middleware
│   └── cors.go             # CORS middleware
├── models/
│   ├── user.go             # User data model
│   └── task.go             # Task and Step data models
├── routes/
│   └── routes.go           # API route definitions
├── services/
│   ├── gemini.go           # Gemini AI integration
│   └── jwt.go              # JWT token service
├── utils/
│   ├── database.go         # Database utilities
│   ├── jwt.go              # JWT utilities
│   └── response.go         # Response utilities
├── .env.example            # Environment variables template
├── .gitignore              # Git ignore rules
├── go.mod                  # Go module dependencies
└── README.md               # This file
```

## 🚀 Quick Start

### Prerequisites

- Go 1.19 or higher
- MongoDB (local or Atlas)
- Google Gemini API key

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/Vanaraj10/taskmorph-backend.git
   cd taskmorph-backend
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Environment Setup**
   ```bash
   cp .env.example .env
   ```
   
   Edit `.env` with your configuration:
   ```env
   MONGO_URI=mongodb+srv://username:password@cluster.mongodb.net/taskmorph
   JWT_SECRET=your-super-secret-jwt-key-here
   GEMINI_API_KEY=your-gemini-api-key-here
   PORT=8080
   ```

4. **Run the application**
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## 📚 API Documentation

### Authentication Endpoints

#### Register User
```http
POST /auth/register
```

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "securepassword"
}
```

**Response:**
```json
{
  "message": "User registered successfully"
}
```

#### Login User
```http
POST /auth/login
```

**Request Body:**
```json
{
  "email": "john@example.com",
  "password": "securepassword"
}
```

**Response:**
```json
{
  "message": "Login successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### AI Endpoints

#### Get Task Breakdown
```http
POST /ai/breakdown
```

**Request Body:**
```json
{
  "task": "Build a portfolio website"
}
```

**Response:**
```json
[
  {
    "title": "Plan Website Structure",
    "description": "Create wireframes and define site architecture",
    "is_completed": false
  },
  {
    "title": "Design UI/UX",
    "description": "Create visual designs and user interface mockups",
    "is_completed": false
  }
]
```

### Task Management Endpoints

> **Note**: All task endpoints require authentication. Include the JWT token in the Authorization header:
> ```
> Authorization: Bearer <your-jwt-token>
> ```

#### Create Task
```http
POST /tasks/create
```

**Request Body:**
```json
{
  "title": "Build portfolio website",
  "deadline": "2025-07-15"
}
```

**Response:**
```json
{
  "message": "Task created successfully",
  "task": {
    "id": "60f7b3b3b3b3b3b3b3b3b3b3",
    "title": "Build portfolio website",
    "deadline": "2025-07-15T00:00:00Z",
    "steps": [...],
    "user_id": "60f7b3b3b3b3b3b3b3b3b3b3"
  }
}
```

#### Get All Tasks
```http
GET /tasks/
```

**Response:**
```json
[
  {
    "id": "60f7b3b3b3b3b3b3b3b3b3b3",
    "title": "Build portfolio website",
    "deadline": "2025-07-15T00:00:00Z",
    "progress": 40,
    "steps": [...]
  }
]
```

#### Get Single Task
```http
GET /tasks/:id
```

#### Complete a Step
```http
PATCH /tasks/:taskID/step/:stepID/complete
```

#### Delete Task
```http
DELETE /tasks/:id
```

## 🔐 Authentication

This API uses JWT (JSON Web Tokens) for authentication. After successful login, include the token in the Authorization header:

```
Authorization: Bearer <your-jwt-token>
```

Tokens expire after 72 hours.

## 🗄️ Database Schema

### User Model
```go
type User struct {
    ID       ObjectID `bson:"_id,omitempty"`
    Name     string   `bson:"name"`
    Email    string   `bson:"email"`
    Password string   `bson:"password"`
}
```

### Task Model
```go
type Task struct {
    ID       ObjectID `bson:"_id,omitempty"`
    UserID   string   `bson:"user_id"`
    Title    string   `bson:"title"`
    Deadline time.Time `bson:"deadline"`
    Steps    []Step   `bson:"steps"`
}

type Step struct {
    ID          ObjectID `bson:"_id,omitempty"`
    Title       string   `bson:"title"`
    Description string   `bson:"description"`
    IsCompleted bool     `bson:"is_completed"`
}
```

## 🧪 Testing

You can test the API using tools like:
- **Postman**: Import the collection from `docs/postman-collection.json`
- **curl**: Use the examples provided in this README
- **Thunder Client**: VS Code extension for API testing

### Example cURL Commands

**Register a user:**
```bash
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "securepassword"
  }'
```

**Login:**
```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "securepassword"
  }'
```

**Create a task:**
```bash
curl -X POST http://localhost:8080/tasks/create \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "title": "Learn Go programming",
    "deadline": "2025-08-01"
  }'
```

## 🚦 Status Codes

- `200` - Success
- `201` - Created
- `400` - Bad Request
- `401` - Unauthorized
- `404` - Not Found
- `500` - Internal Server Error

## 🔧 Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `MONGO_URI` | MongoDB connection string | ✅ |
| `JWT_SECRET` | Secret key for JWT signing | ✅ |
| `GEMINI_API_KEY` | Google Gemini API key | ✅ |
| `PORT` | Server port (default: 8080) | ❌ |
| `ENV` | Environment (development/production) | ❌ |

## 📈 Future Enhancements

- [ ] Task categories and tags
- [ ] Due date reminders
- [ ] Task sharing and collaboration
- [ ] Real-time notifications
- [ ] Task templates
- [ ] Analytics and reporting
- [ ] Mobile app integration
- [ ] Subtask management
- [ ] File attachments
- [ ] Task prioritization

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👥 Authors

- **VJ** - *Initial work* - [Vanaraj10](https://github.com/Vanaraj10)

## 🙏 Acknowledgments

- Google Gemini AI for intelligent task breakdown
- MongoDB for scalable data storage
- Gin framework for fast HTTP routing
- JWT for secure authentication

---

**Happy Task Management! 🎯**
