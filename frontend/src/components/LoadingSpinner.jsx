import React from 'react';
import { Brain } from 'lucide-react';

const LoadingSpinner = ({ size = 'md', text = '' }) => {
  const sizeClasses = {
    sm: 'h-4 w-4',
    md: 'h-6 w-6',
    lg: 'h-8 w-8',
    xl: 'h-12 w-12',
  };

  return (
    <div className="flex flex-col items-center justify-center space-y-3">
      <div className="relative">
        <div className={`animate-spin rounded-full border-b-2 border-blue-500 ${sizeClasses[size]}`}></div>
        <Brain className={`absolute inset-0 text-blue-500/20 ${sizeClasses[size]}`} />
      </div>
      {text && (
        <p className="text-gray-400 text-sm animate-pulse">{text}</p>
      )}
    </div>
  );
};

export default LoadingSpinner;
