import { useState } from 'react'
import axios from 'axios'

const api = axios.create({ baseURL: '/api' })

function Login({ onLogin }) {
    const [isRegister, setIsRegister] = useState(false)
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [error, setError] = useState(null)
    const [loading, setLoading] = useState(false)

    const handleSubmit = async (e) => {
        e.preventDefault()
        setError(null)
        setLoading(true)

        try {
            if (isRegister) {
                await api.post('/register', { username, password })
                // Auto-login after registration
            }
            const res = await api.post('/login', { username, password })
            onLogin(res.data.token, res.data.username)
        } catch (err) {
            setError(err.response?.data?.error || 'Connection failed')
        } finally {
            setLoading(false)
        }
    }

    return (
        <div className="app">
            <header className="header">
                <h1>📝 Todo App</h1>
                <p>Full-Stack Learning Project</p>
                <div className="tech-stack">
                    <span className="tech-badge">⚛️ React</span>
                    <span className="tech-badge">🔷 Go</span>
                    <span className="tech-badge">🌐 Gin</span>
                    <span className="tech-badge">🔐 JWT</span>
                </div>
            </header>

            <div className="auth-card">
                <h2 className="auth-title">
                    {isRegister ? '创建账号' : '登录'}
                </h2>

                <form className="auth-form" onSubmit={handleSubmit}>
                    <input
                        type="text"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                        placeholder="用户名 (Username)"
                        autoFocus
                        required
                        minLength={2}
                    />
                    <input
                        type="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        placeholder="密码 (Password, min 6 chars)"
                        required
                        minLength={6}
                    />

                    {error && (
                        <div className="error" onClick={() => setError(null)}>
                            ⚠️ {error}
                        </div>
                    )}

                    <button type="submit" className="btn-add" disabled={loading}>
                        {loading ? '请稍候...' : isRegister ? '注册' : '登录'}
                    </button>
                </form>

                <p className="auth-switch">
                    {isRegister ? '已有账号？' : '没有账号？'}
                    <button
                        className="btn-link"
                        onClick={() => {
                            setIsRegister(!isRegister)
                            setError(null)
                        }}
                    >
                        {isRegister ? '去登录' : '注册一个'}
                    </button>
                </p>
            </div>
        </div>
    )
}

export default Login
