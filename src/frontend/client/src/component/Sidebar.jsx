import React from 'react'
import AlgoChoice from './AlgoChoice'
import History from './History'
import { useState, useEffect } from 'react';

function Sidebar({chatHistory, algorithm}) {
  const[algo, setAlgo] = useState('');
  const[data, setData] = useState([]);

  const API = 'chat'
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
    fetchHistory()
  }, [data]);

  const algoData = (data) => {
    setAlgo(data);
    algorithm(data);
    console.log("Algoritma yang dipakai: " + data);
  }

  const createSession = () => {

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