import React, { useState, useEffect } from 'react';
import { Plus, Search, Filter, Brain, CheckCircle, Clock, AlertCircle } from 'lucide-react';
import { taskService } from '../services/api';
import TaskCard from './TaskCard';
import TaskForm from './TaskForm';

const Dashboard = () => {
  const [tasks, setTasks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [showTaskForm, setShowTaskForm] = useState(false);
  const [searchTerm, setSearchTerm] = useState('');
  const [filterStatus, setFilterStatus] = useState('all'); // all, completed, pending, overdue
  const [error, setError] = useState('');

  useEffect(() => {
    fetchTasks();
  }, []);

  const fetchTasks = async () => {
    try {
      setLoading(true);
      const response = await taskService.getTasks();
      setTasks(response.data);
      setError('');
    } catch (err) {
      setError('Failed to fetch tasks');
      console.error('Failed to fetch tasks:', err);
    } finally {
      setLoading(false);
    }
  };

  const handleTaskCreated = () => {
    fetchTasks();
  };

  const handleTaskUpdated = () => {
    fetchTasks();
  };

  const handleTaskDeleted = (taskId) => {
    setTasks(tasks.filter(task => task.id !== taskId));
  };

  const isOverdue = (deadline) => {
    return new Date(deadline) < new Date();
  };

  const filteredTasks = tasks.filter(task => {
    // Search filter
    const matchesSearch = task.title.toLowerCase().includes(searchTerm.toLowerCase());
    
    // Status filter
    let matchesStatus = true;
    switch (filterStatus) {
      case 'completed':
        matchesStatus = task.progress === 100;
        break;
      case 'pending':
        matchesStatus = task.progress < 100 && (!task.deadline || !isOverdue(task.deadline));
        break;
      case 'overdue':
        matchesStatus = task.progress < 100 && task.deadline && isOverdue(task.deadline);
        break;
      default:
        matchesStatus = true;
    }
    
    return matchesSearch && matchesStatus;
  });

  const stats = {
    total: tasks.length,
    completed: tasks.filter(task => task.progress === 100).length,
    pending: tasks.filter(task => task.progress < 100 && (!task.deadline || !isOverdue(task.deadline))).length,
    overdue: tasks.filter(task => task.progress < 100 && task.deadline && isOverdue(task.deadline)).length,
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center min-h-64">
        <div className="flex items-center space-x-3 text-gray-400">
          <div className="animate-spin rounded-full h-6 w-6 border-b-2 border-blue-500"></div>
          <span>Loading your tasks...</span>
        </div>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between space-y-4 sm:space-y-0">
        <div>
          <h1 className="text-2xl font-bold text-white">Task Dashboard</h1>
          <p className="text-gray-400 mt-1">Manage your AI-powered tasks efficiently</p>
        </div>
        <button
          onClick={() => setShowTaskForm(true)}
          className="btn-primary flex items-center space-x-2"
        >
          <Plus className="h-5 w-5" />
          <span>New Task</span>
        </button>
      </div>

      {/* Stats Cards */}
      <div className="grid grid-cols-2 lg:grid-cols-4 gap-4">
        <div className="card p-4">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-gray-400 text-sm">Total Tasks</p>
              <p className="text-2xl font-bold text-white">{stats.total}</p>
            </div>
            <Brain className="h-8 w-8 text-blue-500" />
          </div>
        </div>
        
        <div className="card p-4">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-gray-400 text-sm">Completed</p>
              <p className="text-2xl font-bold text-green-400">{stats.completed}</p>
            </div>
            <CheckCircle className="h-8 w-8 text-green-500" />
          </div>
        </div>
        
        <div className="card p-4">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-gray-400 text-sm">In Progress</p>
              <p className="text-2xl font-bold text-yellow-400">{stats.pending}</p>
            </div>
            <Clock className="h-8 w-8 text-yellow-500" />
          </div>
        </div>
        
        <div className="card p-4">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-gray-400 text-sm">Overdue</p>
              <p className="text-2xl font-bold text-red-400">{stats.overdue}</p>
            </div>
            <AlertCircle className="h-8 w-8 text-red-500" />
          </div>
        </div>
      </div>

      {/* Filters */}
      <div className="flex flex-col sm:flex-row space-y-4 sm:space-y-0 sm:space-x-4">
        {/* Search */}
        <div className="relative flex-1">
          <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-gray-400" />
          <input
            type="text"
            placeholder="Search tasks..."
            className="input-field pl-10 w-full"
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
          />
        </div>
        
        {/* Status Filter */}
        <div className="relative">
          <Filter className="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-gray-400" />
          <select
            className="input-field pl-10 pr-8 appearance-none cursor-pointer"
            value={filterStatus}
            onChange={(e) => setFilterStatus(e.target.value)}
          >
            <option value="all">All Tasks</option>
            <option value="pending">In Progress</option>
            <option value="completed">Completed</option>
            <option value="overdue">Overdue</option>
          </select>
        </div>
      </div>

      {/* Error Message */}
      {error && (
        <div className="bg-red-900 border border-red-700 text-red-100 px-4 py-3 rounded-lg">
          {error}
        </div>
      )}

      {/* Tasks Grid */}
      {filteredTasks.length > 0 ? (
        <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
          {filteredTasks.map((task) => (
            <TaskCard
              key={task.id}
              task={task}
              onTaskUpdated={handleTaskUpdated}
              onTaskDeleted={handleTaskDeleted}
            />
          ))}
        </div>
      ) : (
        <div className="text-center py-12">
          <div className="max-w-md mx-auto">
            <Brain className="h-16 w-16 text-gray-600 mx-auto mb-4" />
            <h3 className="text-lg font-medium text-gray-300 mb-2">
              {searchTerm || filterStatus !== 'all' 
                ? 'No tasks match your filters' 
                : 'No tasks yet'}
            </h3>
            <p className="text-gray-500 mb-6">
              {searchTerm || filterStatus !== 'all'
                ? 'Try adjusting your search or filter criteria'
                : 'Create your first AI-powered task to get started'}
            </p>
            {!searchTerm && filterStatus === 'all' && (
              <button
                onClick={() => setShowTaskForm(true)}
                className="btn-primary"
              >
                Create Your First Task
              </button>
            )}
          </div>
        </div>
      )}

      {/* Task Form Modal */}
      {showTaskForm && (
        <TaskForm
          onTaskCreated={handleTaskCreated}
          onClose={() => setShowTaskForm(false)}
        />
      )}
    </div>
  );
};

export default Dashboard;
