import React from 'react'
import MessageBubbleRight from './MessageBubbleRight'
import MessageBubbleLeft from './MessageBubbleLeft'
import { useEffect, useState } from 'react';

function Conversation({chatID, triggerSignal}) {
  // const [file, getFile] = useState
  const[data, setData] = useState('');
  const API = 'chat/' + chatID
  const[trigger, setTrigger] = useState(triggerSignal)

  const dataFetch = async () => {
    const fetched = await (await fetch(API)).json();

    setData(fetched.chats);
    //console.log("CONVERSATION ID: " + chatID)
    console.log(fetched);
  }

  useEffect(() => {
    dataFetch();

    const interval = setInterval(() => {
      console.log('This will be called every 1 seconds');
      dataFetch();
    }, 1000);

    return () => clearInterval(interval);
    console.log("CHANGED")
    // fetchChats()
  }, [data]);

  // useEffect(() => {
    
  //   dataFetch();
  // },[triggerSignal])

  // while(!data){
  //   console.log("Fetching failed conversation data")
  //   fetchChats()
  // }



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