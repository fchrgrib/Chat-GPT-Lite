import React from 'react'
import MessageBubble from './MessageBubble'

function Conversation() {
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