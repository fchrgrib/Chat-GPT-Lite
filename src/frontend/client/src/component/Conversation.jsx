import React from 'react'
import MessageBubble from './MessageBubble'
import { useState } from 'react';

function Conversation({parentToChild}) {
  const[data, setData] = useState('');

  const dataToMessage = () => {
    setData({parentToChild});
  }

  return (
    <div className='conversation'>
        <MessageBubble/>
        <MessageBubble/>
        <MessageBubble/>
        <MessageBubble/>
        <MessageBubble/>
        <MessageBubble/>
        <MessageBubble/>
        <MessageBubble/>
        <MessageBubble/>
        <div className='flex-container'></div>
    </div>
  )
}

export default Conversation