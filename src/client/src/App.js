// import './App.css';
import React from 'react'
import Sidebar from './component/Sidebar';
import './style.scss';
import Chat from './component/Chat';
import io from 'socket.io-client'

const socket = io.connect("htttp//localhost:3001");

function App() {
  return (
    <div className="App">
      <div className='container'>
        <Sidebar/>
        <Chat/>
      </div>
    </div>
  );
}

export default App;
