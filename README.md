# Go Serverless Application

A modern serverless application built with Go and Fiber framework, deployed on Vercel.

## 🚀 Features

- Built with Go and Fiber framework
- Serverless deployment on Vercel
- Clerk Authorization
- Hot reloading during development (using Air)
- Clean project structure

## 📋 Prerequisites

- Go 1.24.1
- Node.js and pnpm
- Vercel CLI
- Git
- Air (for development hot reloading)

## 🛠️ Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/go-serverless.git
cd go-serverless
```

2. Install Go dependencies:

```bash
go mod download
```

3. Install Node.js dependencies:

```bash
pnpm install
```

4. Install Air for hot reloading (optional):

```bash
go install github.com/cosmtrek/air@latest
```

5. Set up environment variables:

```bash
cp .env.example .env
```

Edit the `.env` file with your configuration.

## 🔧 Development

You have two options for local development:

### Option 1: Using Air (Recommended for Go development)

```bash
air
```

This will start the server with hot reloading enabled. The application will be available at `http://localhost:8080`

### Option 2: Using Vercel CLI

```bash
pnpm dev
```

This will start the server using Vercel's development environment.

Note: The `main.go` file is only used for local development. In production, the application runs as a serverless function on Vercel.

## 🚀 Deployment

The application is configured for deployment on Vercel. To deploy:

1. Login to Vercel:

```bash
pnpm setup
```

2. Deploy to Vercel:

```bash
pnpm start
```

## 📁 Project Structure

```
.
├── api/           # API routes and handlers
├── app/           # Application core
├── docs/          # Documentation
├── middleware/    # Middleware functions
├── public/        # Static files
└── main.go        # Development entry point (local only)
```

## 🔐 Environment Variables

Create a `.env` file in the root directory with the following variables:

```
# Add your environment variables here
```

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE).
