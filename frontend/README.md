# TaskMorph Frontend 🎯

A modern, responsive React frontend for the TaskMorph AI-enhanced task management system. Built with React, Vite, and Tailwind CSS for a clean, dark, and mobile-friendly user experience.

## ✨ Features

- **🌙 Dark Theme**: Beautiful dark UI optimized for long work sessions
- **📱 Mobile-First**: Fully responsive design that works on all devices
- **🤖 AI Integration**: Seamless integration with Google Gemini AI for task breakdown
- **⚡ Fast Performance**: Built with Vite for lightning-fast development and build times
- **🎨 Modern UI**: Clean, intuitive interface with smooth animations
- **🔐 Secure Authentication**: JWT-based authentication with automatic token management
- **📊 Progress Tracking**: Visual progress indicators and completion statistics

## 🛠️ Tech Stack

- **Frontend Framework**: React 18
- **Build Tool**: Vite
- **Styling**: Tailwind CSS
- **Icons**: Lucide React
- **HTTP Client**: Axios
- **State Management**: React Context API

## 🏃‍♂️ Quick Start

### Prerequisites

- Node.js 16+ and npm
- TaskMorph Backend running on `http://localhost:8080`

### Installation

1. **Navigate to the frontend directory**
   ```bash
   cd frontend
   ```

2. **Install dependencies**
   ```bash
   npm install
   ```

3. **Start the development server**
   ```bash
   npm run dev
   ```

4. **Open your browser**
   ```
   http://localhost:5173
   ```

## 📱 Usage Guide

### 1. Authentication
- Register with name, email, and password
- Login with email and password
- Automatic token management and session persistence

### 2. Creating Tasks
- Click "New Task" to open the task creation form
- Enter a task title and optional deadline  
- Use "Preview AI Breakdown" to see AI-generated steps
- Click "Create Task" to save

### 3. Managing Tasks
- View all tasks in the dashboard grid
- Click the expand arrow to see task steps
- Check off completed steps by clicking the circle icon
- Delete tasks using the trash icon
- Use search and filters to find specific tasks

---

**Ready to transform your tasks with AI? Start the frontend and let TaskMorph do the magic! ✨**+ Vite

This template provides a minimal setup to get React working in Vite with HMR and some ESLint rules.

Currently, two official plugins are available:

- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react) uses [Babel](https://babeljs.io/) for Fast Refresh
- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react-swc) uses [SWC](https://swc.rs/) for Fast Refresh

## Expanding the ESLint configuration

If you are developing a production application, we recommend using TypeScript with type-aware lint rules enabled. Check out the [TS template](https://github.com/vitejs/vite/tree/main/packages/create-vite/template-react-ts) for information on how to integrate TypeScript and [`typescript-eslint`](https://typescript-eslint.io) in your project.
