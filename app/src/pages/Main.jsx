import { useEffect, useState } from "react"
import api from "../api/api"
import { toast } from "react-toastify";


const Main = ({ store, setStore, navigate }) => {

  useEffect(() => {
    api('main').then(data => setStore({ ...store, ...data }))
  }, [])

  const createBottle = (event) => {
    event.preventDefault();
    toast('Run wine bottle creator')
    const data = new FormData(event.target)
    api('createBottle', {
      method: 'POST',
      body: data
    }).then(res => setStore({ ...store, BOTTLES: res }))
  }

  if(store.CURRENT=="false") return <button onClick={()=>navigate({path:'/wine'})}>Setup wine</button>

  return (
    <div className="flex flex-col gap-5">

      <span>Wine: {store.CURRENT}</span>

      <form onSubmit={createBottle} className="flex gap-5">
        <input
          type="text"
          name="name"
          placeholder="Type bottle name"
          className="text-[#D0D7E5] bg-[#252930] p-2.5 px-4 flex-1"
          required />
        <button type="submit">Create bottle</button>
      </form>

      <div className="flex flex-col gap-5">
        {store?.BOTTLES && store.BOTTLES.map(bottle => (
          <div 
            key={bottle.name} 
            className="text-[#D0D7E5] bg-[#252930] p-2.5 cursor-pointer hover:opacity-40"
            onClick={()=>navigate({path:'/bottle',props:bottle})}
          >{bottle.name}</div>
        ))}
      </div>
    </div>
  )
}

export default Main