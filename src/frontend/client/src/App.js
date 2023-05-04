import React, { useState, useEffect } from 'react'
import Sidebar from './component/Sidebar';
import './style.scss';
import Chat from './component/Chat';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import AppPanel from './component/AppPanel';

function App() {
  const current = new Date();
  const time = current.toLocaleTimeString("en-US");

  const[data, setData] = useState([]);
  const[def, setDef] = useState([]);
  //const API = 'http://localhost:3000/ChatHistory'
  const API = 'chat'
  const fetchHistory = () => {
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

  useEffect(() => {
    //console.log("use effect in progress")
    setInterval(()=>{
      fetchHistory()
    }, 2000)
  }, []);

  return (
    <div className="App">
      {/* <div className='container'> */}
        {/* <Sidebar/> */}
        <Router>
          <Routes>
            {
              data && data.map((item) => {
                // return <Route path={"/" + item.id} element={<Chat chatID={item.id}/>}/>
                console.log("ID: " + item.id)
                const path = "/" + item.id
                console.log("Path created: " + path)
                return <Route path={"/" + item.id} element={<AppPanel chatHistory={item}/>}/>
                
              })
            }
            <Route path='/' element={<AppPanel chatHistory={def}/>}/>
          </Routes>
        </Router>
      {/* </div> */}
    </div>
  );
}

export default App;