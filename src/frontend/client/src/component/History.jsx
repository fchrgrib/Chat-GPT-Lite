import React, { useReducer } from 'react'
import HistoryTab from './HistoryTab'
import { useState, useEffect} from 'react';

const initialState = {
  loading: true,
  post: {}
}

const reducer = (state, action) => {
  if(action.type == 'SUCCESS'){
    return {
      loading: false,
      post: action.payload
    }
  }
  return state
}

function History({chatHistory}) {
  const[data, setData] = useState([]);
  const[state, dispatch] = useReducer(reducer, initialState)
  const API = 'chat'

  useEffect(() => {
    const dataFetch = async () => {
      const fetched = await (await fetch(API)).json();

      setData(fetched.history);
    }
    dataFetch();

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