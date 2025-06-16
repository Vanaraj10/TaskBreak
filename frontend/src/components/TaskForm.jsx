import React, { useState } from 'react';
import { Plus, Calendar, Sparkles } from 'lucide-react';
import { taskService, aiService } from '../services/api';

const TaskForm = ({ onTaskCreated, onClose }) => {
  const [formData, setFormData] = useState({
    title: '',
    deadline: '',
  });
  const [loading, setLoading] = useState(false);
  const [previewSteps, setPreviewSteps] = useState([]);
  const [showPreview, setShowPreview] = useState(false);
  const [error, setError] = useState('');

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
    setError('');
  };

  const handlePreview = async () => {
    if (!formData.title.trim()) {
      setError('Please enter a task title');
      return;
    }

    setLoading(true);
    setError('');

    try {
      const response = await aiService.getTaskBreakdown(formData.title);
      setPreviewSteps(response.data);
      setShowPreview(true);
    } catch (err) {
      setError('Failed to generate task breakdown. Please try again.');
    } finally {
      setLoading(false);
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError('');

    try {
      await taskService.createTask(formData);
      onTaskCreated();
      onClose();
    } catch (err) {
      setError(err.response?.data?.error || 'Failed to create task');
    } finally {
      setLoading(false);
    }
  };

  const today = new Date().toISOString().split('T')[0];

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50">
      <div className="bg-gray-800 rounded-lg shadow-xl w-full max-w-2xl max-h-[90vh] flex flex-col">
        {/* Header */}
        <div className="p-6 border-b border-gray-700">
          <h2 className="text-xl font-bold text-white flex items-center space-x-2">
            <Plus className="h-5 w-5" />
            <span>Create New Task</span>
          </h2>
        </div>

        {/* Content */}
        <div className="flex-1 overflow-y-auto p-6">
          <form onSubmit={handleSubmit} className="space-y-6">
            {/* Task Title */}
            <div>
              <label htmlFor="title" className="block text-sm font-medium text-gray-300 mb-2">
                Task Title
              </label>
              <input
                id="title"
                name="title"
                type="text"
                required
                className="input-field w-full"
                placeholder="e.g., Build a portfolio website"
                value={formData.title}
                onChange={handleChange}
              />
            </div>

            {/* Deadline */}
            <div>
              <label htmlFor="deadline" className="block text-sm font-medium text-gray-300 mb-2">
                Deadline (Optional)
              </label>
              <div className="relative">
                <Calendar className="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-gray-400" />
                <input
                  id="deadline"
                  name="deadline"
                  type="date"
                  min={today}
                  className="input-field pl-10 w-full"
                  value={formData.deadline}
                  onChange={handleChange}
                />
              </div>
            </div>

            {/* AI Preview Button */}
            {formData.title && !showPreview && (
              <div>
                <button
                  type="button"
                  onClick={handlePreview}
                  disabled={loading}
                  className="btn-secondary w-full flex items-center justify-center space-x-2"
                >
                  <Sparkles className="h-4 w-4" />
                  <span>{loading ? 'Generating...' : 'Preview AI Breakdown'}</span>
                </button>
              </div>
            )}

            {/* AI Preview */}
            {showPreview && previewSteps.length > 0 && (
              <div className="bg-gray-900 rounded-lg p-4 border border-gray-700">
                <h3 className="text-lg font-medium text-white mb-3 flex items-center space-x-2">
                  <Sparkles className="h-5 w-5 text-blue-400" />
                  <span>AI-Generated Steps</span>
                </h3>
                <div className="space-y-2 max-h-48 overflow-y-auto">
                  {previewSteps.map((step, index) => (
                    <div key={index} className="bg-gray-800 rounded-lg p-3 border border-gray-600">
                      <h4 className="font-medium text-gray-200 text-sm">{step.title}</h4>
                      <p className="text-gray-400 text-xs mt-1">{step.description}</p>
                    </div>
                  ))}
                </div>
              </div>
            )}

            {error && (
              <div className="bg-red-900 border border-red-700 text-red-100 px-4 py-3 rounded-lg">
                {error}
              </div>
            )}
          </form>
        </div>

        {/* Footer */}
        <div className="p-6 border-t border-gray-700 flex justify-end space-x-3">
          <button
            type="button"
            onClick={onClose}
            className="btn-secondary"
          >
            Cancel
          </button>
          <button
            onClick={handleSubmit}
            disabled={loading || !formData.title.trim()}
            className="btn-primary disabled:opacity-50 disabled:cursor-not-allowed"
          >
            {loading ? 'Creating...' : 'Create Task'}
          </button>
        </div>
      </div>
    </div>
  );
};

export default TaskForm;
