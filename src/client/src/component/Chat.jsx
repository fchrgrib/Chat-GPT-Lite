import React from 'react'
import InputBar from './InputBar'
import Conversation from './Conversation'

function Chat() {
  return (
    <div className='chat'>
        <Conversation/>
        <InputBar/>
    </div>
  )
}

export default Chat