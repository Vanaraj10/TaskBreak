import React from 'react';
import { AlertTriangle, RefreshCw } from 'lucide-react';

const ErrorBoundary = ({ error, onRetry }) => {
  return (
    <div className="min-h-screen bg-gray-900 flex items-center justify-center px-4">
      <div className="max-w-md w-full text-center">
        <div className="bg-gray-800 rounded-lg p-8 border border-gray-700">
          <AlertTriangle className="h-16 w-16 text-red-500 mx-auto mb-4" />
          <h2 className="text-xl font-bold text-white mb-2">Oops! Something went wrong</h2>
          <p className="text-gray-400 mb-6">
            {error || 'An unexpected error occurred. Please try again.'}
          </p>
          {onRetry && (
            <button
              onClick={onRetry}
              className="btn-primary flex items-center space-x-2 mx-auto"
            >
              <RefreshCw className="h-4 w-4" />
              <span>Try again</span>
            </button>
          )}
        </div>
      </div>
    </div>
  );
};

export default ErrorBoundary;
