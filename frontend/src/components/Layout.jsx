import React from 'react';
import { LogOut, User, Brain } from 'lucide-react';
import { useAuth } from '../contexts/AuthContext';

const Layout = ({ children }) => {
  const { user, logout } = useAuth();

  return (
    <div className="min-h-screen bg-gray-900">
      {/* Header */}
      <header className="bg-gray-800 border-b border-gray-700">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            {/* Logo */}
            <div className="flex items-center space-x-3">
              <div className="bg-blue-600 p-2 rounded-lg">
                <Brain className="h-6 w-6 text-white" />
              </div>
              <h1 className="text-xl font-bold text-white">TaskMorph</h1>
              <span className="text-xs bg-gray-700 text-gray-300 px-2 py-1 rounded-full">
                AI-Powered
              </span>
            </div>

            {/* User menu */}
            <div className="flex items-center space-x-4">
              <div className="flex items-center space-x-2 text-gray-300">
                <User className="h-4 w-4" />
                <span className="text-sm hidden sm:inline">{user?.email}</span>
              </div>
              <button
                onClick={logout}
                className="flex items-center space-x-2 text-gray-300 hover:text-white transition-colors duration-200 p-2 rounded-lg hover:bg-gray-700"
              >
                <LogOut className="h-4 w-4" />
                <span className="text-sm hidden sm:inline">Logout</span>
              </button>
            </div>
          </div>
        </div>
      </header>

      {/* Main content */}
      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {children}
      </main>
    </div>
  );
};

export default Layout;
