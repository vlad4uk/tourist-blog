import React, { useState } from 'react';
import axios from 'axios';
import AuthLayout from '../components/AuthLayout.tsx';
import SocialButtons from '../components/SocialButtons.tsx';
import FriendsFront from '../assets/FrontFriends.jpg';


// Базовый URL Go-бэкенда
const API_BASE_URL = 'http://localhost:8080/api/auth'; 

const LoginPage: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    setLoading(true);

    try {
      //КЛЮЧЕВОЙ ЗАПРОС С withCredentials: true
      const response = await axios.post(
        `${API_BASE_URL}/login`,
        {
          email: email,
          password: password,
        },
        {
          // ОБЯЗАТЕЛЬНО для отправки куки на Go-бэкенд и получения куки в ответ
          withCredentials: true, 
        }
      );

      // В случае успеха:
      console.log('Login successful:', response.data);
      // Здесь можно сохранить данные пользователя в глобальное состояние (Context/Redux)
      // и перенаправить на /dashboard.
      alert(`Вход успешен! Добро пожаловать, ${response.data.user.username}`);

    } catch (err) {
      if (axios.isAxiosError(err) && err.response) {
        // Отображение ошибки с бэкенда (например, "Invalid credentials" или "Your account has been blocked")
        setError(err.response.data.error || 'Login failed');
      } else {
        setError('An unexpected error occurred. Check server connection.');
      }
    } finally {
      setLoading(false);
    }
  };

  const LoginIllustration = () => (
    <div style={{ textAlign: 'center' }}>
      <img src={FriendsFront} alt="Login friends" style={{ maxWidth: '80%', height: 'auto' }} />
    </div>
  );

  return (
    <AuthLayout illustration={<LoginIllustration />}>
      <h1 className="auth-title">Welcome to sneat!</h1>
      <p className="auth-subtitle">Please sign-in to your account and start the adventure</p>

      <form onSubmit={handleSubmit}>
        {/* Поле Email */}
        <div className="form-field-group">
          <label htmlFor="email" className="form-label">Email or Username</label>
          <input 
            type="text" 
            id="email" 
            className="form-input" 
            placeholder="Enter your email or username" 
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>

        {/* Поле Password */}
        <div className="form-field-group">
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '6px' }}>
            <label htmlFor="password" className="form-label" style={{ marginBottom: 0 }}>Password</label>
            <a href="/forgot-password" style={{ fontSize: '0.875rem', color: '#696cff', textDecoration: 'none' }}>Forgot password?</a>
          </div>
          <input 
            type="password" 
            id="password" 
            className="form-input" 
            placeholder="············"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>

        {/* Checkbox "Remember me" */}
        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '20px' }}>
          <label style={{ fontSize: '0.875rem', color: '#3f4254', display: 'flex', alignItems: 'center' }}>
            <input type="checkbox" style={{ marginRight: '8px' }} />
            Remember me
          </label>
        </div>
        
        {/* Отображение ошибки */}
        {error && <p style={{ color: 'red', textAlign: 'center', fontSize: '0.875rem', marginBottom: '10px' }}>{error}</p>}


        <button type="submit" className="primary-button" disabled={loading}>
          {loading ? 'Logging In...' : 'Login'}
        </button>
      </form>

      <p style={{ textAlign: 'center', fontSize: '0.9rem', marginTop: '20px' }}>
        New on our platform? <a href="/register" style={{ color: '#696cff', textDecoration: 'none', fontWeight: 500 }}>Create an account</a>
      </p>

      <div className="separator">or</div>
      <SocialButtons />
    </AuthLayout>
  );
};

export default LoginPage;