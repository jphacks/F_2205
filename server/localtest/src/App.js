import {useState} from 'react'
import axios from 'axios'

import { Title } from './components/title'
import { Event } from './event'

function App() {
    const LOCAL_API_URL    = `http://localhost:8080/room`
    const [roomId, setRoomId] = useState("")

    const createRoom = ()=>{
        axios.post(LOCAL_API_URL)
            .then((res)=>{
                console.log("room API",res.data.data.id)
                setRoomId(res.data.data.id)
            })
    }

    return (
        <div className="mx-36 my-8">
            <p className='text-6xl font-bold text-gray-600 mb-4'>
                API Test
            </p>
            <p className='text-xl text-gray-700 mb-12'>
                検証ツールからconsoleを開き、テストをしてください
            </p>

            <div className=''>
            {
                roomId?
                <>
                <Event roomId={roomId}/>
                </>
                :
                <>
                    <Title title={"CreateRoom"} />
                    <div className='flex mb-8'>
                        <button 
                            className='
                                bg-blue-500 hover:bg-blue-700 
                                text-white font-bold py-2 
                                px-4 rounded-full
                                mx-4 my-2
                            ' 
                            onClick={()=>{createRoom()}}
                        >
                            ルーム作成
                        </button>
                    </div>
                </>
            }
            </div>
        </div>
    );
}

export default App;
