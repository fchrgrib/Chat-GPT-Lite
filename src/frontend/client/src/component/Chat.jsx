import React from 'react'
import InputBar from './InputBar'
import Conversation from './Conversation'
import { useState } from 'react'

function Chat() {
  
  const[data, setData] = useState('');

  const parentToChild = () => {
    setData()
  }

  const childToParent = (childData) => {
    setData(childData);
    //alert(childData);
  }
  return (
    <div className='chat'>
        <Conversation parentToChild={data}/>
        <InputBar childToParent={childToParent}/>
    </div>
  )
}

export default Chat