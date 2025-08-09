import { useState } from 'react';

function App() {
  const [input, setInput] = useState('');
  const [messages, setMessages] = useState([]);

  const sendMessage = async () => {
    if (!input.trim()) return;

    // 把用户输入先加到聊天记录
    setMessages(prev => [...prev, { sender: 'user', text: input }]);

    try {
      const res = await fetch('http://localhost:8080/chat', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ message: input }),
      });
      const data = await res.json();

      // 把 AI 回复加入聊天记录
      setMessages(prev => [...prev, { sender: 'ai', text: data.reply }]);
    } catch (err) {
      console.error('发送消息失败', err);
    }

    setInput('');
  };

  return (
    <div style={{ maxWidth: '600px', margin: '0 auto', padding: '20px' }}>
      <h1>AI 聊天</h1>
      <div style={{ border: '1px solid #ccc', padding: '10px', height: '400px', overflowY: 'auto' }}>
        {messages.map((msg, idx) => (
          <div key={idx} style={{ textAlign: msg.sender === 'user' ? 'right' : 'left' }}>
            <b>{msg.sender === 'user' ? '我' : 'AI'}: </b>
            <span>{msg.text}</span>
          </div>
        ))}
      </div>
      <div style={{ marginTop: '10px', display: 'flex' }}>
        <input
          style={{ flex: 1, padding: '8px' }}
          value={input}
          onChange={(e) => setInput(e.target.value)}
          placeholder="输入消息..."
        />
        <button style={{ marginLeft: '5px', padding: '8px' }} onClick={sendMessage}>发送</button>
      </div>
    </div>
  );
}

export default App;


