import React, { useState, useEffect, useRef } from 'react';
import './App.css';

function App() {
  const [logs, setLogs] = useState([]);
  const [connectionStatus, setConnectionStatus] = useState('connecting');
  const wsRef = useRef(null);
  const logContainerRef = useRef(null);

  useEffect(() => {
    const websocket = new WebSocket('ws://localhost:8080/ws');
    wsRef.current = websocket;

    websocket.onopen = () => {
      console.log('Connected to WebSocket server');
      setConnectionStatus('connected');
      setLogs(prev => [...prev, 'Connected to WebSocket server']);
    };

    websocket.onmessage = (event) => {
      console.log('Message from server:', event.data);
      setLogs(prevLogs => [...prevLogs, event.data]);
    };

    websocket.onclose = () => {
      console.log('Disconnected from WebSocket server');
      setConnectionStatus('disconnected');
      setLogs(prev => [...prev, 'Disconnected from WebSocket server']);
    };

    websocket.onerror = (error) => {
      console.error('WebSocket error:', error);
      setConnectionStatus('disconnected');
      setLogs(prev => [...prev, `WebSocket error: ${error.message || 'Unknown error'}`]);
    };

    return () => {
      if (wsRef.current && wsRef.current.readyState === WebSocket.OPEN) {
        wsRef.current.close();
      }
    };
  }, []);

  useEffect(() => {
    if (logContainerRef.current) {
      logContainerRef.current.scrollTop = logContainerRef.current.scrollHeight;
    }
  }, [logs]);

  const getLogType = (log) => {
    if (log.includes('[ERROR]')) return 'error';
    return 'info';
  };

  return (
    <div className="App">
      <div className="header">
        <h1>Auto Messenger Logs</h1>
        <p>Real-time monitoring of message delivery and system events</p>
      </div>
      
      <div className="log-container" ref={logContainerRef}>
        {logs.length === 0 ? (
          <div className="no-logs">
            <svg width="80" height="80" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M21 7L13 15L9 11L3 17" stroke="#888888" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"/>
              <path d="M21 12V7H16" stroke="#888888" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"/>
            </svg>
            <div>No logs yet. Waiting for messages...</div>
          </div>
        ) : (
          logs.map((log, index) => (
            <div key={index} className={`log-entry ${getLogType(log)}`}>
              {log}
            </div>
          ))
        )}
      </div>
      
      <div className="status-indicator">
        <div className={`status-dot ${connectionStatus}`}></div>
        <div>
          {connectionStatus === 'connected' && 'Connected to server'}
          {connectionStatus === 'disconnected' && 'Disconnected from server'}
          {connectionStatus === 'connecting' && 'Connecting to server...'}
        </div>
      </div>
      
      <div className="footer">
        Auto Messenger © {new Date().getFullYear()}
      </div>
    </div>
  );
}

export default App; 