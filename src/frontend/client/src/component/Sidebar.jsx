import React from 'react'
import AlgoChoice from './AlgoChoice'
import History from './History'
import { useState, useEffect } from 'react';

function Sidebar({chatHistory, algorithm}) {
  const[algo, setAlgo] = useState('');
  const[data, setData] = useState([]);

  const API = 'chat'
  const APIput = API + '/add_chat'
  const fetchHistory = () => {
    fetch(API).then((res) => {
      //console.log("Ini data fetch sebelum jadi json: " + res);
      return res.json();
    }).then((resAfter) => {
      // console.log("Ini data fetch setelah jadi json: " + res);
      setData(resAfter.history);
    })
  }

  useEffect(() => {
    //console.log("use effect in progress")
    setInterval(()=>{
      fetchHistory()
    }, 2000)
  }, []);

  const algoData = (data) => {
    setAlgo(data);
    algorithm(data);
    console.log("Algoritma yang dipakai: " + data);
  }

  const createSession = () => {
    fetch(APIput, {
      method: 'PUT',
      headers: {"Content-Type" : "application/json"}
    })
  }

  return (
    <div className='sidebar'>
        <History chatHistory={chatHistory}/>
        <div className='flex-container'></div>
        <button onClick={createSession}>+</button>
        <AlgoChoice algoData={algoData}/>
    </div>
  )
}

export default Sidebar