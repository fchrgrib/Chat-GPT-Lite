import React, { useState, useEffect, Link } from 'react'
import Sidebar from './component/Sidebar';
import './style.scss';
import Chat from './component/Chat';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import AppPanel from './component/AppPanel';
import {useDispatch} from 'react-redux'

export const useData = (url) => {
  const [state, setState] = useState();

  useEffect(() => {
    const dataFetch = async () => {
      const data = await (await fetch(url)).json();

      setState(data);
    };

    dataFetch();
  }, [url]);

  return { data: state };
};

function App() {
  const current = new Date();
  const time = current.toLocaleTimeString("en-US");

  const[data, setData] = useState([]);
  const[def, setDef] = useState([]);
  const[req, setReq] = useState('');
  //const API = 'http://localhost:3000/ChatHistory'
  const API = 'chat'
  const fetchHistory = async () => {
    //console.log("fetch history before")
    fetch(API).then((res) => {
      //console.log("Ini data fetch sebelum jadi json: " + res);
      return res.json();
    }).then((resAfter) => {
      // console.log("Ini data fetch setelah jadi json: " + res);
      setData(resAfter.history);
      setDef(resAfter)
      // console.log("Here")
      // console.log(data)
    })
  }
  
  // const {template} = useData(API);
  // const {data} = template.history;

  //const dispatch = useDispatch();

  // const fetchLoad = (req) => {
  //   setReq(req);
  //   dispatch(debouncedSearch(req));
  // }

  useEffect(() => {
    const dataFetch = async () => {
      const fetched = await (await fetch(API)).json();

      setData(fetched.history);
      setDef(fetched);
    }
    dataFetch();
  }, []);

  if(data){
    return (
      <div className="App">
        {/* <div className='container'> */}
          {/* <Sidebar/> */}
          <Router>
            <Routes>
              
              {
                data && data.map((item) => {
                  // return <Route path={"/" + item.id} element={<Chat chatID={item.id}/>}/>
                  return <Route key={item.id} path={`/${item.id}`} element={<AppPanel chatHistory={item}/>}/>
                  
                })
              }
              <Route path='/' element={<AppPanel chatHistory={def}/>}/>
            </Routes>
          </Router>
        {/* </div> */}
      </div>
    );
  }
  // console.log("Failed to get Data")
  // fetchHistory()
  // return <p>Loading...</p>
  
}

export default App;