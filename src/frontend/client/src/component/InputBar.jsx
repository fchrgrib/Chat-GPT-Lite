import React, { useEffect } from 'react'
import { useState } from 'react'

function InputBar({chatID, algo, triggerSignal}) {
  const [dataAlgo, setDataAlgo] = useState('');
  const [message, setMessage] = useState('');
  const [trigger, setTrigger] = useState(0);
  const handleChange = (event) => {
    setMessage(event.target.value);
  }

  // useEffect(() => {
  //   fetchHistory()
  // }, [data]);

  const handleKeyDown = (event) => {
    const current = new Date();
    const currentTime = current.toLocaleTimeString("en-US");
    const API = "chat/" + chatID + "/post_chat"
    if(event.key === 'Enter') {
      fetch(API, {
        method: 'POST',
        body: JSON.stringify({

          chat: event.target.value,
          type: algo,
          
        }),
        headers: {"Content-Type" : "application/json"}
      })
      console.log("Dari input bar: ");
      console.log(chatID);
      console.log("Entered chat: " + event.target.value);
      event.target.value = ""
      setTrigger((trigger) => trigger + 1)
      triggerSignal(trigger)
      console.log("REQUETS HAS BEEN SENT")
    }
  }

  return (
    <div className='inputbar'>
        <input type='text' name="message" placeholder='Masukan pertanyaan disini' onChange={handleChange} onKeyDown={handleKeyDown}/>
    </div>
  )
}

export default InputBar