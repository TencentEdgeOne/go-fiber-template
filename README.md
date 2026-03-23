# Go Cloud Functions on EdgeOne Pages - Fiber Framework

A full-stack demo application built with Next.js + Tailwind CSS frontend and Go Fiber backend, showcasing how to deploy Go Cloud Functions using the Fiber framework on EdgeOne Pages with RESTful API routing.

## 🚀 Features

- **Fiber Framework Integration**: Express-inspired Go web framework built on fasthttp, with middleware, JSON parsing, and route groups
- **Modern UI Design**: Dark theme with #1c66e5 accent color, responsive layout with interactive elements
- **Interactive API Testing**: Built-in API endpoint panel — click "Call" to test each REST endpoint in real-time
- **RESTful API Design**: Complete Todo CRUD operations with structured route groups (`/api/todos`)
- **TypeScript Support**: Complete type definitions and type safety on the frontend

## 🛠️ Tech Stack

### Frontend
- **Next.js 15** - React full-stack framework (with Turbopack)
- **React 19** - User interface library
- **TypeScript** - Type-safe JavaScript
- **Tailwind CSS 4** - Utility-first CSS framework

### UI Components
- **shadcn/ui** - High-quality React components
- **Lucide React** - Beautiful icon library
- **class-variance-authority** - Component style variant management
- **clsx & tailwind-merge** - CSS class name merging utilities

### Backend
- **Go 1.21** - Cloud Functions runtime
- **Fiber v2** - Express-inspired Go web framework built on fasthttp

## 📁 Project Structure

```
go-fiber/
├── cloud-functions/                # Go Cloud Functions source
│   ├── api.go                     # Fiber app with all REST API routes
│   ├── go.mod                     # Go module definition
│   └── go.sum                     # Go dependency checksums
├── src/
│   ├── app/                       # Next.js App Router
│   │   ├── globals.css           # Global styles (dark theme)
│   │   ├── layout.tsx            # Root layout
│   │   └── page.tsx              # Main page (API testing UI)
│   ├── components/               # React components
│   │   └── ui/                   # UI base components
│   │       ├── button.tsx        # Button component
│   │       └── card.tsx          # Card component
│   └── lib/                      # Utility functions
│       └── utils.ts              # Common utilities (cn helper)
├── public/                        # Static assets
│   ├── eo-logo-blue.svg          # EdgeOne logo (blue)
│   └── eo-logo-white.svg         # EdgeOne logo (white)
├── package.json                   # Project configuration
└── README.md                     # Project documentation
```

## 🚀 Quick Start

### Requirements

- Node.js 18+
- pnpm (recommended) or npm
- Go 1.21+ (for local development)

### Install Dependencies

```bash
pnpm install
# or
npm install
```

### Development Mode

```bash
edgeone pages dev
```

Visit [http://localhost:8088](http://localhost:8088) to view the application.

### Build Production Version

```bash
edgeone pages build
```

## 🎯 Core Features

### 1. Fiber REST API Routes

All API endpoints are defined in a single `cloud-functions/api.go` file using Fiber's route groups:

| Method | Route | Description |
|--------|-------|-------------|
| GET | `/` | Welcome message with route list |
| GET | `/health` | Health check |
| GET | `/api/todos` | List all todos |
| POST | `/api/todos` | Create a new todo |
| GET | `/api/todos/:id` | Get todo by ID |
| PATCH | `/api/todos/:id/toggle` | Toggle todo completion |
| DELETE | `/api/todos/:id` | Delete a todo |

### 2. Interactive API Testing Panel

- 7 pre-configured API endpoint cards with "Call" buttons
- Real-time JSON response display with syntax highlighting
- POST request support with pre-filled JSON body
- Loading states and error handling

### 3. Fiber Framework Convention

The Go backend uses Fiber's standard patterns — route groups, body parsing, and built-in middleware:

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
    app := fiber.New()
    app.Use(recover.New())

    app.Get("/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "status":    "ok",
            "framework": "fiber",
        })
    })

    api := app.Group("/api")
    api.Get("/todos", listTodos)
    api.Post("/todos", createTodo)
    api.Get("/todos/:id", getTodo)
    api.Patch("/todos/:id/toggle", toggleTodo)
    api.Delete("/todos/:id", deleteTodo)

    app.Listen(":9000")
}
```

### 4. Data Model

```go
type Todo struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Completed bool      `json:"completed"`
    CreatedAt time.Time `json:"createdAt"`
}
```

## 🔧 Configuration

### Tailwind CSS Configuration
The project uses Tailwind CSS 4 with custom color variables:

```css
:root {
  --primary: #1c66e5;        /* Primary color */
  --background: #000000;     /* Background color */
  --foreground: #ffffff;     /* Foreground color */
}
```

### Component Styling
Uses `class-variance-authority` to manage component style variants with multiple preset styles.

## 📚 Documentation

- **EdgeOne Pages Official Docs**: [https://edgeone.ai/document/go-functions](https://edgeone.ai/document/go-functions)
- **Fiber Framework**: [https://docs.gofiber.io](https://docs.gofiber.io)
- **Next.js Documentation**: [https://nextjs.org/docs](https://nextjs.org/docs)
- **Tailwind CSS Documentation**: [https://tailwindcss.com/docs](https://tailwindcss.com/docs)

## 🚀 Deployment Guide

### EdgeOne Pages Deployment

1. Push code to GitHub repository
2. Create new project in EdgeOne Pages console
3. Select GitHub repository as source
4. Configure build command: `edgeone pages build`
5. Deploy project

### Go Fiber Cloud Function

Create `cloud-functions/api.go` in project root with your Fiber application:

```go
package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()

    app.Get("/hello", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "message": "Hello from Go Fiber on EdgeOne Pages!",
        })
    })

    app.Listen(":9000")
}
```

## Deploy

[![Deploy with EdgeOne Pages](https://cdnstatic.tencentcs.com/edgeone/pages/deploy.svg)](https://edgeone.ai/pages/new?from=github&template=go-fiber)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/github/choosealicense.com/blob/gh-pages/_licenses/mit.txt) file for details.
