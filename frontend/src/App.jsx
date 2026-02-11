import { useState } from 'react'
import { Routes, Route, Navigate, useNavigate } from 'react-router-dom'
import Login from './pages/Login.jsx'
import Todos from './pages/Todos.jsx'

function App() {
    const [token, setToken] = useState(localStorage.getItem('token'))
    const [username, setUsername] = useState(localStorage.getItem('username') || '')
    const navigate = useNavigate()

    const handleLogin = (newToken, newUsername) => {
        localStorage.setItem('token', newToken)
        localStorage.setItem('username', newUsername)
        setToken(newToken)
        setUsername(newUsername)
        navigate('/')
    }

    const handleLogout = () => {
        localStorage.removeItem('token')
        localStorage.removeItem('username')
        setToken(null)
        setUsername('')
        navigate('/login')
    }

    return (
        <Routes>
            <Route
                path="/login"
                element={
                    token ? <Navigate to="/" /> : <Login onLogin={handleLogin} />
                }
            />
            <Route
                path="/"
                element={
                    token ? (
                        <Todos token={token} username={username} onLogout={handleLogout} />
                    ) : (
                        <Navigate to="/login" />
                    )
                }
            />
        </Routes>
    )
}

export default App
