import { useEffect, useState } from "react"
import ChangeDir from "./components/ChangeDir"
import api from "../../api/api"
import WineList from "./components/WineList"
import Downloaded from "./components/Downloaded"

const Wine = ({ store, setStore }) => {
  const [wines, setWines] = useState([])
  
  useEffect(() => {
    api('main').then(data => setStore({ ...store, ...data }))
    api('wineReleases').then(data => setWines(data))
  }, [])



  return (
    <div className="flex flex-col gap-5">
      
      <ChangeDir store={store} setStore={setStore} />
      {store?.WINE_PATH && store?.WINE_PATH != "false" && <>
        <WineList wines={wines} store={store} setStore={setStore}/>
      </>}
      {store?.DOWNLOADED && <Downloaded downloaded={store?.DOWNLOADED} current={store.CURRENT} setStore={setStore} store={store}/> }

    </div>
  )
}

export default Wine