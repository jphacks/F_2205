import { 
    NewMemberData1,
    NewMemberData2,
    SetFocusData,
    DelFocusData,
    DelAllFocusData,
    SetScreenShotData,
    SetEffectData,
    SetSoundData,
} from './data/data'

import { Button } from './components/button'
import { Title } from './components/title'

export const Event = ({roomId})=> {
    const WS_LOCAL_API_URL = `ws://localhost:8080/ws/${roomId}`
    const conn = new WebSocket(WS_LOCAL_API_URL)

    conn.onopen = (e)=>{
        console.log("ws connect")
    }

    conn.onclose = (e)=>{
        console.log("ws connect close...")
    }

    conn.onmessage = (e)=>{
        console.log("on message : ",JSON.parse(e.data))
    }

    const sendData = (data)=>{
        conn.send(JSON.stringify(data))
    }


    return (
        <div className="mx-24 my-8">
            <Title title={"FocusEvent"} />
            <div className='mb-8'>
                <div className='flex'>
                    <Button data={NewMemberData1} name={"memberデータ追加(hoge1)"} sendData={sendData} />
                    <Button data={NewMemberData2} name={"memberデータ追加(hoge2)"} sendData={sendData} />
                </div>
                <div className='flex'>
                    <Button data={SetFocusData} name={"hoge1とhoge2をfocusする"} sendData={sendData} />
                    <Button data={DelFocusData} name={"hoge1とhoge2のfocusを削除する"} sendData={sendData} />
                    <Button data={DelAllFocusData} name={"hoge1のすべてのfocusを削除する"} sendData={sendData} />
                </div>
            </div>

            <Title title={"ScreenShotEvent"} />
            <div className='flex mb-8'>
                <Button data={SetScreenShotData} name={"スクリーンショットイベント開始"} sendData={sendData} /> 
            </div>

            <Title title={"EffectEvent"} />
            <div className='flex mb-8'>
                <Button data={SetEffectData} name={"エフェクトイベント開始"} sendData={sendData} /> 
            </div>

            <Title title={"SoundEvent"} />
            <div className='flex mb-8'>
                <Button data={SetSoundData} name={"hoge1サウンドを設定する"} sendData={sendData} /> 
            </div>

        </div>
    );
}
