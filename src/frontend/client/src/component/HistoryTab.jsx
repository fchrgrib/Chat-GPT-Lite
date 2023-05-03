import React from 'react'
import io from 'socket.io-client'
const socket = io.connect("https://localhost:3001");


function HistoryTab({HistoryData}) {
  return (
    <div className='history-tab'>
        {HistoryData}
    </div>
  )
}

export default HistoryTab