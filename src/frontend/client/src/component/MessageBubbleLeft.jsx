import React from 'react'
import { useState } from 'react';

function MessageBubbleLeft({dataToMessage}) {
  return (
    <div className='messagebubble'>
        <p>{dataToMessage}</p>
    </div>
  )
}


export default MessageBubbleLeft