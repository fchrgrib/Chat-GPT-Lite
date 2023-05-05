import React from 'react'
import Sidebar from './Sidebar'
import Chat from './Chat'
import { useState } from 'react';

function AppPanel({chatHistory}) {
    const[algo, setAlgo] = useState('');

    const algorithm = (data) => {
        setAlgo(data);
    }
    console.log(chatHistory.id)
  return (
    <div className='appPanel'>
        <Sidebar chatHistory={chatHistory} algorithm={algorithm}/>
        <Chat chatID={chatHistory.id} algo={algo}/>
    </div>
  )
}

export default AppPanel