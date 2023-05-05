import React, { useEffect } from 'react'
import InputBar from './InputBar'
import Conversation from './Conversation'
import { useState } from 'react'

function Chat({chatID, algo}) {
  
  const[data, setData] = useState('');
  const[trigger, setTrigger] = useState(0);

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

  const triggerSignal = (trig) => {
    setTrigger((trig) => trig + 1)
    console.log("THERES A CHANGE")
  }

  // useEffect(() =>{
  //   console.log("THERES A CHANGE")
  // }, [trigger])

  return (
    <div className='chat'>
        <Conversation chatID={chatID} triggerSignal={trigger}/>
        <InputBar chatID={chatID} algo={algo} triggerSignal={triggerSignal}/>
    </div>
  )
}

export default Chat