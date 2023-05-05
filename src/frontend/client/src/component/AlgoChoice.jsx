import React, { useEffect, useState } from 'react'

function AlgoChoice({algoData}) {
  const [algo, setAlgo] = useState('')

  useEffect(() => {
    algoData("KMP");
  })

  return (
    <div className='algoChoice'>
        <hr className="solid"></hr>
        <div className='algo-buttons'>
            <div className='KMPButton'>
                <input type="radio" value="KMP" name="algo" onChange={e => algoData(e.target.value)} defaultChecked/> KMP
            </div>
            <div className='BMButton'>
                <input type="radio" value="BM" name="algo" onChange={e => algoData(e.target.value)}/> BM
            </div>
        </div>
    </div>
  )
}

export default AlgoChoice