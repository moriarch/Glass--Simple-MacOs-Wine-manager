import { toast } from 'react-toastify';
import api from '../../../api/api'
import { useRef } from 'react';

function ChangeDir({ store, setStore }) {

    const toastId = useRef(null);

    const notify = () => toastId.current = toast("Loanding");

    const handleChangeFolder = () => {
        notify()
        api('changeFolder').then(data => {
            toast.dismiss(toastId.current)
            setStore({ ...store, ...data })
        })
    }

    return (
        <>
            <span className="text-[#C4CBDA] whitespace-nowrap">Bottles and Wine path:</span>
            <div className="text-[#D0D7E5] bg-[#252930] p-2.5 px-4 overflow-hidden text-ellipsis flex-1">{store?.WINE_PATH}</div>
            <button onClick={handleChangeFolder}>Change Folder</button>
        </>
    )
}

export default ChangeDir