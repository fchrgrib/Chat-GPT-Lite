import React from 'react'
import InputBar from './InputBar'
import Conversation from './Conversation'
import { useState } from 'react'

function Chat({chatID, algo}) {
  
  const[data, setData] = useState('');

  const parentToChild = () => {
    setData()
  }

  const childToParent = (childData) => {
    setData(childData);
    //alert(childData);
  }
  // if(chatID === 0){
  //   return <p>WELCOME BOIS</p>
  // }

  return (
    <div className='chat'>
        <Conversation chatID={chatID}/>
        <InputBar chatID={chatID} algo={algo}/>
    </div>
  )
}

export default Chat