import React from 'react'
import api from '../api/api'
import { toast } from 'react-toastify';

const Bottle = ({ store, setStore, navigate, props }) => {

  const createFormData = () => {
    let data = new FormData();
    data.append('name',props.name)
    return data
  }

  const handleClick = () => {
    toast("Loanding file chooser")
    api('runApp',{
      method:'POST',
      body: createFormData()
    })
  }

  return (
    <div>
        <div className='flex items-center justify-between'>
            <div>Bottle: {props.name}</div>
            <button onClick={()=>navigate({path:'/'})}>Back</button>
        </div>
        <button onClick={handleClick}>Execute app</button>
    </div>
  )
}

export default Bottle