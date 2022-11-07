export const AddNewMemberData1 = {
    type:"ADD_NEW_MEMBER",
    info:{
        member:{
            name:"hoge1",
            peer_id:"peer_hoge1"
        }
    }
}

export const AddNewMemberData2 = {
    type:"ADD_NEW_MEMBER",
    info:{
        member:{
            name:"hoge2",
            peer_id:"peer_hoge2"
        }
    }
}

export const SetFocusData = {
    type:"SET_FOCUS",
    info:{
        focus:{
            from:"hoge1",
            to:"hoge2",
        }
    }
}

export const DelFocusData = {
    type:"DEL_FOCUS",
    info:{
        focus:{
            from:"hoge1",
            to:"hoge2",
        }
    }
}

export const DelAllFocusData = {
    type:"DEL_ALL_FOCUS",
    info:{
        focus:{
            from:"hoge1"
        }
    }
}

export const SetScreenShotData = {
    type:"SET_SCREEN_SHOT"
}

export const SetEffectData = {
    type:"SET_EFFECT",
    info:{
        effect:{
            name:"hoge1",
            type:"happy"
        }
    }
}

export const SetSoundData = {
    type:"SET_SOUND",
    info:{
        sound:{
            type:"hoge1"
        }
    }
}

