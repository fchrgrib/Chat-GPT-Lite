// import './App.css';
import React from 'react'
import Sidebar from './component/Sidebar';
import './style.scss';
import Chat from './component/Chat';
import io from 'socket.io-client'

const socket = io.connect("http://localhost:3001");

function App() {
  const current = new Date();
  const time = current.toLocaleTimeString("en-US");
  // const testClick = () => {
  //   alert(time);
  // };

  return (
    <div className="App">
      <div className='container'>
      {/* <button onClick={testClick}>button</button> */}
        <Sidebar/>
        <Chat/>
      </div>
    </div>
  );
}

export default App;
