import React from 'react'
import { useState } from 'react'

function InputBar({childToParent}) {
  const [message, setMessage] = useState('');
  const handleChange = (event) => {
    setMessage(event.target.value);
  }

  const handleKeyDown = (event) => {
    if(event.key === 'Enter') {
      fetch('http://localhost:3000/Chats', {
        method: 'POST',
        body: JSON.stringify({
          id: Math.random().toString(36).slice(2),
          from: "user",
          chat: event.target.value
        }),
        headers: {"Content-Type" : "application/json"}
      })
      console.log(event.target.value);
      event.target.value = ""

      console.log("REQUETS HAS BEEN SENT")
      childToParent(true);
    }
  }

  return (
    <div className='inputbar'>
        <input type='text' name="message" placeholder='Masukan pertanyaan disini' onChange={handleChange} onKeyDown={handleKeyDown}/>
    </div>
  )
}

export default InputBar