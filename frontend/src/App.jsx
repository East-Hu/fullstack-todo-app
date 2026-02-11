import { useState, useEffect } from 'react'
import axios from 'axios'

// API base — Vite proxy forwards /api to Go backend
const api = axios.create({ baseURL: '/api' })

function App() {
    const [todos, setTodos] = useState([])
    const [newTodo, setNewTodo] = useState('')
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(null)

    // Fetch todos on mount
    useEffect(() => {
        fetchTodos()
    }, [])

    const fetchTodos = async () => {
        try {
            setLoading(true)
            setError(null)
            const res = await api.get('/todos')
            setTodos(res.data || [])
        } catch (err) {
            setError('无法连接到后端服务器，请确保 Go 后端正在运行 (go run main.go)')
            console.error('Fetch error:', err)
        } finally {
            setLoading(false)
        }
    }

    const addTodo = async (e) => {
        e.preventDefault()
        if (!newTodo.trim()) return

        try {
            const res = await api.post('/todos', { title: newTodo.trim() })
            setTodos([...todos, res.data])
            setNewTodo('')
        } catch (err) {
            setError('添加失败: ' + (err.response?.data?.error || err.message))
        }
    }

    const toggleTodo = async (todo) => {
        try {
            const res = await api.put(`/todos/${todo.ID}`, {
                completed: !todo.completed,
            })
            setTodos(todos.map((t) => (t.ID === todo.ID ? res.data : t)))
        } catch (err) {
            setError('更新失败: ' + (err.response?.data?.error || err.message))
        }
    }

    const deleteTodo = async (id) => {
        try {
            await api.delete(`/todos/${id}`)
            setTodos(todos.filter((t) => t.ID !== id))
        } catch (err) {
            setError('删除失败: ' + (err.response?.data?.error || err.message))
        }
    }

    const completedCount = todos.filter((t) => t.completed).length

    return (
        <div className="app">
            <header className="header">
                <h1>📝 Todo App</h1>
                <p>Full-Stack Learning Project</p>
                <div className="tech-stack">
                    <span className="tech-badge">⚛️ React</span>
                    <span className="tech-badge">🔷 Go</span>
                    <span className="tech-badge">🌐 Gin</span>
                    <span className="tech-badge">🗄️ GORM</span>
                </div>
            </header>

            <form className="input-form" onSubmit={addTodo}>
                <input
                    type="text"
                    value={newTodo}
                    onChange={(e) => setNewTodo(e.target.value)}
                    placeholder="添加新任务... (Add a new todo)"
                    autoFocus
                />
                <button type="submit" className="btn-add">
                    添加
                </button>
            </form>

            {error && (
                <div className="error" onClick={() => setError(null)}>
                    ⚠️ {error}
                </div>
            )}

            {loading ? (
                <div className="loading">
                    <div className="loading-spinner"></div>
                    <p>连接后端中...</p>
                </div>
            ) : (
                <>
                    {todos.length > 0 && (
                        <div className="stats">
                            <span>
                                共 <span className="count">{todos.length}</span> 个任务
                            </span>
                            <span>
                                已完成 <span className="count">{completedCount}</span> /{' '}
                                {todos.length}
                            </span>
                        </div>
                    )}

                    <div className="todo-list">
                        {todos.length === 0 ? (
                            <div className="empty-state">
                                <div className="emoji">🎉</div>
                                <p>暂无任务，添加你的第一个 Todo 吧！</p>
                            </div>
                        ) : (
                            todos.map((todo) => (
                                <div
                                    key={todo.ID}
                                    className={`todo-item ${todo.completed ? 'completed' : ''}`}
                                >
                                    <div
                                        className={`checkbox ${todo.completed ? 'checked' : ''}`}
                                        onClick={() => toggleTodo(todo)}
                                    />
                                    <span className="todo-text">{todo.title}</span>
                                    <button
                                        className="btn-delete"
                                        onClick={() => deleteTodo(todo.ID)}
                                    >
                                        删除
                                    </button>
                                </div>
                            ))
                        )}
                    </div>
                </>
            )}
        </div>
    )
}

export default App
