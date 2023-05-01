import React from 'react'
import { useState } from 'react'

function InputBar({childToParent}) {
  const [message, setMessage] = useState('');
  const handleChange = (event) => {
    setMessage(event.target.value);
  }

  const handleKeyDown = (event) => {
    if(event.key === 'Enter') {
      setMessage(event.target.value);
      childToParent(message);
    }
  }
  return (
    <div className='inputbar'>
        <input type='text' name="message" placeholder='Masukan pertanyaan disini' onChange={handleChange} onKeyDown={handleKeyDown}/>
    </div>
  )
}

export default InputBar