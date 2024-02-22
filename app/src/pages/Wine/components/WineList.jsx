import { toast } from "react-toastify";
import api from "../../../api/api"


function Winwineist({ wines, store, setStore }) {

    const createData = ({uri,name}) => {
        let data = new FormData();
        data.append('uri',uri)
        data.append('name',name)
        return data
    }
    const handleDownload = ({name,assets}) => {
        toast('Downloading',{
            autoClose: 2*60*1000,
        })
        api('getWine',{
            method:'POST',
            body:createData({uri:assets[0].browser_download_url, name})
        })
        .then(r=>{
            setStore({...store,DOWNLOADED:r})
            toast.dismiss();
            toast('Downloaded')
            
        })
    }

    return (
        <>
            <span className="text-[#C4CBDA] whitespace-nowrap">WINE list:</span>
            <div className="text-[#D0D7E5] bg-[#252930] max-h-[250px] overflow-scroll">
                {wines.map((wine, i) => (
                    <div key={wine.name + i} className="p-2.5 hover:bg-[#31343D] flex justify-between items-center">
                        <span>{wine.name}</span>
                        <button onClick={() => handleDownload(wine)}>Download</button>
                    </div>
                ))}
            </div>
            
        </>
    )
}

export default Winwineist