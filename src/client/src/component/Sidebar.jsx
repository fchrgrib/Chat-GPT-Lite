import React from 'react'
import AlgoChoice from './AlgoChoice'
import History from './History'

function Sidebar() {
  return (
    <div className='sidebar'>
        <History/>
        <div className='flex-container'></div>
        <AlgoChoice/>
    </div>
  )
}

export default Sidebar