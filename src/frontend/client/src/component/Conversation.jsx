import React from 'react'
import MessageBubbleRight from './MessageBubbleRight'
import MessageBubbleLeft from './MessageBubbleLeft'
import { useEffect, useState } from 'react';

function Conversation({chatID}) {
  // const [file, getFile] = useState
  const[data, setData] = useState([]);
  const API = 'chat/' + chatID
  const fetchChats = () => {
    // fetch(API).then((res) => res.json()).then((resChats) => {
    //   //console.log(res);
    //   //console.log("This is the id: " + chatID)
    //   setData(resChats.chats);
    //   console.log(data)
    // })

    fetch(API).then((res) => {
      //console.log("Ini data fetch sebelum jadi json: " + res);
      return res.json();
    }).then((resAfter) => {
      // console.log("Ini data fetch setelah jadi json: " + res);
      setData(resAfter.chats);
      console.log(data)
    })
  }

  useEffect(() => {
    console.log("fetch chat");
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