import React from 'react';
import { AuthProvider, useAuth } from './contexts/AuthContext';
import AuthForm from './components/AuthForm';
import Dashboard from './components/Dashboard';
import Layout from './components/Layout';

function AppContent() {
  const { user, loading } = useAuth();

  if (loading) {
    return (
      <div className="min-h-screen bg-gray-900 flex items-center justify-center">
        <div className="flex items-center space-x-3 text-gray-400">
          <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-500"></div>
          <span className="text-lg">Loading TaskMorph...</span>
        </div>
      </div>
    );
  }

  if (!user) {
    return <AuthForm onSuccess={() => window.location.reload()} />;
  }

  return (
    <Layout>
      <Dashboard />
    </Layout>
  );
}

function App() {
  return (
    <AuthProvider>
      <AppContent />
    </AuthProvider>
  );
}

export default App;
