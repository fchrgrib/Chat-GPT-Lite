import React from 'react'
import HistoryTab from './HistoryTab'
import { useState, useEffect } from 'react';

function History() {
  const[data, setData] = useState([]);
  const API = 'http://localhost:3000/ChatHistory'
  const fetchHistory = () => {
    fetch(API).then((res) => res.json()).then((res) => {
      console.log(res);
      setData(res);
    })
  }

  useEffect(() => {
    fetchHistory()
  }, []);


  return (
    <div className='history'>
        <div className='history-text'>History</div>
        <div className='history-container'>
          {
            data && data.map((item) => {
              return <HistoryTab HistoryData={item.lastChat}/>
            })
          } 
        </div>
    </div>
  )
}

export default History