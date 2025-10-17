// src/App.tsx
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import LoginPage from './pages/LoginPage.tsx';
import RegisterPage from './pages/RegisterPage.tsx';

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        {/* Страницы для внешнего вида */}
        <Route path="/login" element={<LoginPage />} />
        <Route path="/register" element={<RegisterPage />} />
        
        {/* Добавьте другие маршруты, например, для забытого пароля и главной страницы */}
        <Route path="/" element={<LoginPage />} />
      </Routes>
    </Router>
  );
};

export default App;