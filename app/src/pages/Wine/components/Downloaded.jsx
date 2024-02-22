import api from "../../../api/api"


const Downloaded = ({downloaded, current, store, setStore }) => {

  const createData = ({name}) => {
    let data = new FormData();
    data.append('name',name)
    return data
}
  const handleCurrent = (name) => api('changeCurrent',{
      method:'POST',
      body:createData({name})
  })
  .then(res=>{
    console.log(res)
    setStore({...store,CURRENT:name})
  })

  return (
    <>
    <span className="text-[#C4CBDA] whitespace-nowrap">WINES:</span>
    <div className="text-[#D0D7E5] bg-[#252930]">
       {downloaded.map((wine, i) => (
            <div key={wine + i} className="p-2.5 hover:bg-[#31343D] flex justify-between items-center">
                <span>{wine.name} {wine.name===current && '- current'}</span>
                <button onClick={() => handleCurrent(wine.name)}>Choose</button>
            </div>
        ))} 
    </div>
    
</>
  )
}

export default Downloaded