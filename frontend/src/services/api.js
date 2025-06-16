import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080';

// Create axios instance
const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor to add auth token
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Response interceptor to handle auth errors
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token');
      localStorage.removeItem('user');
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

export const authService = {
  register: (userData) => api.post('/auth/register', userData),
  login: (credentials) => api.post('/auth/login', credentials),
  logout: () => {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
  },
};

export const taskService = {
  getTasks: () => api.get('/tasks/'),
  getTask: (id) => api.get(`/tasks/${id}`),
  createTask: (taskData) => api.post('/tasks/create', taskData),
  deleteTask: (id) => api.delete(`/tasks/${id}`),
  completeStep: (taskId, stepId) => api.patch(`/tasks/${taskId}/step/${stepId}/complete`),
};

export const aiService = {
  getTaskBreakdown: (task) => api.post('/ai/breakdown', { task }),
};

export default api;
