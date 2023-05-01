import React from 'react'

function AlgoChoice() {
  return (
    <div className='algoChoice'>
        <hr class="solid"></hr>
        <div className='algo-buttons'>
            <div className='KMPButton'>
                <input type="radio" value="KMP" name="algo" /> KMP
            </div>
            <div className='BMButton'>
                <input type="radio" value="BM" name="algo" /> BM
            </div>
        </div>
    </div>
  )
}

export default AlgoChoice