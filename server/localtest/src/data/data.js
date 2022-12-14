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
            from:"peer_hoge1",
            to:"peer_hoge2",
        }
    }
}

export const DelFocusData = {
    type:"DEL_FOCUS",
    info:{
        focus:{
            from:"peer_hoge1",
            to:"peer_hoge2",
        }
    }
}

export const DelAllFocusData = {
    type:"DEL_ALL_FOCUS",
    info:{
        focus:{
            from:"peer_hoge1"
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
            peer_id:"peer_hoge1",
            type:"happy"
        }
    }
}

export const SetSoundData = {
    type:"SET_SOUND",
    info:{
        sound:{
            type:"hoge_sound"
        }
    }
}

export const SetRestRoomData1 = {
    type:"SET_REST_ROOM",
    info:{
        rest:{
            peer_id:"peer_hoge1",
            is_rest_room:true
        }
    }
}

export const SetRestRoomData2 = {
    type:"SET_REST_ROOM",
    info:{
        rest:{
            peer_id:"peer_hoge1",
            is_rest_room:false
        }
    }
}
