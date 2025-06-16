import React, { useState } from 'react';
import { Calendar, CheckCircle, Circle, Trash2, ChevronDown, ChevronUp } from 'lucide-react';
import { taskService } from '../services/api';

const TaskCard = ({ task, onTaskUpdated, onTaskDeleted }) => {
  const [expanded, setExpanded] = useState(false);
  const [loading, setLoading] = useState(false);

  const formatDate = (dateString) => {
    const date = new Date(dateString);
    return date.toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
    });
  };

  const isOverdue = (deadline) => {
    return new Date(deadline) < new Date() && task.progress < 100;
  };

  const handleStepComplete = async (stepId) => {
    setLoading(true);
    try {
      await taskService.completeStep(task.id, stepId);
      onTaskUpdated();
    } catch (error) {
      console.error('Failed to complete step:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleDeleteTask = async () => {
    if (window.confirm('Are you sure you want to delete this task?')) {
      setLoading(true);
      try {
        await taskService.deleteTask(task.id);
        onTaskDeleted(task.id);
      } catch (error) {
        console.error('Failed to delete task:', error);
      } finally {
        setLoading(false);
      }
    }
  };

  const completedSteps = task.steps?.filter(step => step.is_completed).length || 0;
  const totalSteps = task.steps?.length || 0;
  const progressPercentage = totalSteps > 0 ? Math.round((completedSteps / totalSteps) * 100) : 0;

  return (
    <div className="card p-6 transition-all duration-200 hover:shadow-xl">
      {/* Task Header */}
      <div className="flex items-start justify-between mb-4">
        <div className="flex-1">
          <h3 className="text-lg font-semibold text-white mb-2">{task.title}</h3>
          
          {/* Deadline */}
          {task.deadline && (
            <div className={`flex items-center space-x-2 text-sm ${
              isOverdue(task.deadline) 
                ? 'text-red-400' 
                : 'text-gray-400'
            }`}>
              <Calendar className="h-4 w-4" />
              <span>{formatDate(task.deadline)}</span>
              {isOverdue(task.deadline) && (
                <span className="bg-red-900 text-red-100 px-2 py-1 rounded-full text-xs">
                  Overdue
                </span>
              )}
            </div>
          )}
        </div>

        {/* Actions */}
        <div className="flex items-center space-x-2">
          <button
            onClick={() => setExpanded(!expanded)}
            className="p-2 text-gray-400 hover:text-white transition-colors duration-200"
          >
            {expanded ? (
              <ChevronUp className="h-5 w-5" />
            ) : (
              <ChevronDown className="h-5 w-5" />
            )}
          </button>
          <button
            onClick={handleDeleteTask}
            disabled={loading}
            className="p-2 text-gray-400 hover:text-red-400 transition-colors duration-200 disabled:opacity-50"
          >
            <Trash2 className="h-5 w-5" />
          </button>
        </div>
      </div>

      {/* Progress Bar */}
      <div className="mb-4">
        <div className="flex justify-between items-center mb-2">
          <span className="text-sm text-gray-400">Progress</span>
          <span className="text-sm font-medium text-gray-300">
            {completedSteps}/{totalSteps} steps ({progressPercentage}%)
          </span>
        </div>
        <div className="w-full bg-gray-700 rounded-full h-2">
          <div
            className={`h-2 rounded-full transition-all duration-300 ${
              progressPercentage === 100 
                ? 'bg-green-500' 
                : progressPercentage > 50 
                  ? 'bg-blue-500' 
                  : 'bg-yellow-500'
            }`}
            style={{ width: `${progressPercentage}%` }}
          />
        </div>
      </div>

      {/* Steps (Expanded) */}
      {expanded && task.steps && task.steps.length > 0 && (
        <div className="space-y-3">
          <h4 className="text-sm font-medium text-gray-300 border-b border-gray-700 pb-2">
            Task Steps
          </h4>
          <div className="space-y-2 max-h-64 overflow-y-auto">
            {task.steps.map((step, index) => (
              <div
                key={step.id || index}
                className={`flex items-start space-x-3 p-3 rounded-lg transition-all duration-200 ${
                  step.is_completed 
                    ? 'bg-green-900/20 border border-green-800' 
                    : 'bg-gray-700 hover:bg-gray-600'
                }`}
              >
                <button
                  onClick={() => !step.is_completed && handleStepComplete(step.id)}
                  disabled={loading || step.is_completed}
                  className={`mt-0.5 transition-colors duration-200 ${
                    step.is_completed
                      ? 'text-green-400'
                      : 'text-gray-400 hover:text-blue-400'
                  }`}
                >
                  {step.is_completed ? (
                    <CheckCircle className="h-5 w-5" />
                  ) : (
                    <Circle className="h-5 w-5" />
                  )}
                </button>
                <div className="flex-1 min-w-0">
                  <h5 className={`text-sm font-medium ${
                    step.is_completed 
                      ? 'text-green-300 line-through' 
                      : 'text-gray-200'
                  }`}>
                    {step.title}
                  </h5>
                  <p className={`text-xs mt-1 ${
                    step.is_completed 
                      ? 'text-green-400/70' 
                      : 'text-gray-400'
                  }`}>
                    {step.description}
                  </p>
                </div>
              </div>
            ))}
          </div>
        </div>
      )}

      {/* Completion Badge */}
      {progressPercentage === 100 && (
        <div className="mt-4 bg-green-900 border border-green-700 rounded-lg p-3 text-center">
          <CheckCircle className="h-6 w-6 text-green-400 mx-auto mb-1" />
          <p className="text-green-300 text-sm font-medium">Task Completed!</p>
        </div>
      )}
    </div>
  );
};

export default TaskCard;
