import React from 'react'
import io from 'socket.io-client'
const socket = io.connect("https://localhost:3001");


function HistoryTab() {
  return (
    <div className='history-tab'>
        HistoryTab
    </div>
  )
}

export default HistoryTab