import React from 'react'
import MessageBubbleRight from './MessageBubbleRight'
import MessageBubbleLeft from './MessageBubbleLeft'
import { useEffect, useState } from 'react';

function Conversation({parentToChild}) {
  // const [file, getFile] = useState
  const[data, setData] = useState([]);
  const API = 'http://localhost:3000/Chats'
  const fetchChats = () => {
    fetch(API).then((res) => res.json()).then((res) => {
      console.log(res);
      setData(res);
    })
  }

  useEffect(() => {
    fetchChats()
  }, [data]);



  // const dataToMessage = () => {
  //   setData({parentToChild});
  // }

  return (
    <div className='conversation'>
      {
        data && data.map((item) => {
          if(item.from === "user"){
            return <MessageBubbleRight dataToMessage={item.chat}/>
          }else{
            return <MessageBubbleLeft dataToMessage={item.chat}/>
          }
        })
      }
        <div className='flex-container'></div>
    </div>
  )
}

export default Conversation