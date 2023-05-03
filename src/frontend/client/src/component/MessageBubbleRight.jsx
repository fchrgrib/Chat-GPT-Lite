import React from 'react'
// import { useEffect, useState } from 'react';

function MessageBubbleRight({dataToMessage}) {
    return (
      <div className='messagebubble owner'>
          <p>{dataToMessage}</p>
      </div>
    )
}


export default MessageBubbleRight