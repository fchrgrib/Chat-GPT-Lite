import React from 'react'
import { useNavigate } from "react-router-dom";



function HistoryTab({HistoryData, historyDesc}) {
  const navigate = useNavigate()
  const navigateLink = "/" + HistoryData.id
  const API = "chat/" + HistoryData.id + "/delete_chat"

  const deleteRequest = () => {
    //console.log("fetch history before")
    fetch(API, { method: 'DELETE'})
  }
  
  const historyClicked = () => {
    console.log("History " + HistoryData.id + " has been clicked")
    navigate(navigateLink)
  }

  // Current session
  if(HistoryData.id === historyDesc.id){
    return (
      <div className='text-container current' key={HistoryData.id}>
        <div className='history-tab'>
            {HistoryData.lastChat}
        </div>
        <button onClick={deleteRequest}>Clear</button>
      </div>
    )
  }

  // Not current session
  return (
    <div className='text-container' key={HistoryData.id}>
      <div className='history-tab' onClick={historyClicked}>
          {HistoryData.lastChat}
      </div>
    </div>
  )
}

export default HistoryTab