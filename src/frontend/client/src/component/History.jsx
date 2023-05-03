import React from 'react'
import HistoryTab from './HistoryTab'
import { useState, useEffect } from 'react';


function History({chatHistory}) {
  const[data, setData] = useState([]);
  // const API = 'http://localhost:3000/ChatHistory'
  // const fetchHistory = () => {
  //   fetch(API).then((res) => res.json()).then((res) => {
  //     //console.log(res);
  //     setData(res);
  //   })
  // }
  const API = 'chat'
  const fetchHistory = () => {
    //console.log("fetch history before")
    fetch(API).then((res) => {
      //console.log("Ini data fetch sebelum jadi json: " + res);
      return res.json();
    }).then((resAfter) => {
      // console.log("Ini data fetch setelah jadi json: " + res);
      setData(resAfter.history);
      console.log("Here")
      console.log(data)
    })
  }

  useEffect(() => {
    fetchHistory()
  }, [data]);


  return (
    <div className='history'>
        <div className='history-text'>History</div>
        <div className='history-container'>
          {
            data && data.map((item) => {
              return <HistoryTab HistoryData={item} historyDesc={chatHistory}/>
            })
          } 
        </div>
    </div>
  )
}

export default History