import { useState, useRef, useEffect } from 'react';
import './App.css';

function App() {
  const [input, setInput] = useState('');
  const [messages, setMessages] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const messagesEndRef = useRef(null);

  const sendMessage = async () => {
    if (!input.trim() || loading) return;
    setError(null);
    setMessages(prev => [...prev, { sender: 'user', text: input }]);
    setLoading(true);

    try {
      const res = await fetch('http://localhost:8080/chat', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ message: input }),
      });
      if (!res.ok) throw new Error(`服务器返回状态: ${res.status}`);

      const data = await res.json();
      setMessages(prev => [...prev, { sender: 'ai', text: data.reply }]);
    } catch (err) {
      setError('发送消息失败，请重试');
    } finally {
      setInput('');
      setLoading(false);
    }
  };

  const handleKeyDown = (e) => {
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault();
      sendMessage();
    }
  };

  useEffect(() => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' });
  }, [messages]);

  return (
    <div className="app-container">
      <h1 className="app-title">AI 聊天机器人</h1>
      <div className="chat-window">
        {messages.length === 0 && (
          <p className="empty-tip">请输入消息开始对话...</p>
        )}
        {messages.map((msg, idx) => (
          <div
            key={idx}
            className={`message-row ${msg.sender === 'user' ? 'message-user' : 'message-ai'}`}
          >
            <div className="message-bubble">{msg.text}</div>
          </div>
        ))}
        <div ref={messagesEndRef} />
      </div>

      {error && <div className="error-tip">{error}</div>}

      <div className="input-area">
        <textarea
          className="message-input"
          placeholder="输入你的消息，按回车发送，Shift+回车换行..."
          value={input}
          onChange={e => setInput(e.target.value)}
          onKeyDown={handleKeyDown}
          disabled={loading}
          rows={3}
        />
        <button
          className="send-button"
          onClick={sendMessage}
          disabled={loading}
          title="发送消息"
        >
          {loading ? '发送中...' : '发送'}
        </button>
      </div>
    </div>
  );
}

export default App;