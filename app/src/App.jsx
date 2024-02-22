import { useState } from 'react'
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

import Main from './pages/Main'
import Wine from './pages/Wine/Wine'
import Bottle from './pages/Bottle';

const pages = {
  '/': Main,
  '/wine':Wine,
  '/bottle':Bottle
}

function App() {

  const [page, setPage] = useState({path:'/',props:null})
  
  const navigate = ({path,props=null}) => setPage({path,props})

  const Component = pages[page.path]
  
  const [store, setStore] = useState({
    WINE_PATH:false,
    CURRENT:false,
    WINE:false
  })

  return (
    <div className='min-h-screen h-full bg-[#1F2228] text-[#C4CBDA] grid grid-cols-6'>
      <div className="bg-[#24272E] col-span-2 flex flex-col min-h-screen">
        <Link onClick={()=>navigate({path:'/'})} active={page.path==="/"} >Main</Link>
        <Link onClick={()=>navigate({path:'/wine'})} active={page.path==="/wine"} >Wine</Link> 
      </div>
      <main className='p-2.5 text-sm col-span-4'>
        <Component navigate={navigate} store={store} setStore={setStore} props={page.props}/>
      </main>
      <ToastContainer theme="dark" />
    </div>
  )
}



const Link = ({ children, onClick, active }) => {
  return <span onClick={onClick} className={`p-2.5 hover:bg-black hover:text-white transition-all ease-in duration-300 border-l-2 border-transparent  ${active && 'bg-black !border-[#C4CBDA]'}`}>{children}</span>
}


export default App
