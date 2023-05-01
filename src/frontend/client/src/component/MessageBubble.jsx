import React from 'react'
import { useState } from 'react';

function MessageBubble({dataToMessage}) {
  return (
    <div className='messagebubble owner'>
        <p>{dataToMessage? dataToMessage:"Default"}</p>
    </div>
  )
}


export default MessageBubble