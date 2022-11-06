# API仕様書

## Focus

### GET /ws/:room_id (AddNewMember)

roomにmemberを新規追加する

fromさんを接続情報に新規追加する

send
```
const data = {
    type:"NEW_MEMBER",
    info:{
        focus:{
            from:"hoge1"
        }
    }
}
```

response
```
{
    "event_type": "NEW_MEMBER",
    "focus_members": [
        {
            "name": "hoge1",
            "connects": []
        }
    ],
    "effect_member": {
        "name": "",
        "type": ""
    }
}
```

<br>

### GET /ws/:room_id (SetFocus)

to,fromにmember名をいれて、ユーザー同士を接続状態にする

fromさんとtoさんが存在するものと仮定する

send
```        
const data = {
    type:"SET_FOCUS",
    info:{
        focus:{
            from:"hoge1",
            to:"hoge2",
        }
    }
}
```

response
```
{
    "event_type": "SET_FOCUS",
    "focus_members": [
        {
            "name": "hoge1",
            "connects": [
                {
                    "name": "hoge2"
                }
            ]
        },
        {
            "name": "hoge2",
            "connects": [
                {
                    "name": "hoge1"
                }
            ]
        }
    ],
    "effect_member": {
        "name": "",
        "type": ""
    }
}
```

<br>

### GET /ws/:room_id (DelFocus)

to,fromにmember名をいれて、ユーザー同士の接続状態を解除する

send
```
const data = {
    type:"DEL_FOCUS",
    info:{
        focus:{
            from:"hoge1",
            to:"hoge2",
        }
    }
}
```


response
```
{
    "event_type": "DEL_FOCUS",
    "focus_members": [
        {
            "name": "hoge1",
            "connects": []
        },
        {
            "name": "hoge2",
            "connects": []
        }
    ],
    "effect_member": {
        "name": "",
        "type": ""
    }
}
```

<br>

### GET /ws/:room_id (DelAllFocus)

fromにmember名をいれて、そのユーザーの持っている接続状態を解除する

send
```
const data = {
    type:"DEL_ALL_FOCUS",
    info:{
        focus:{
            from:"hoge1"
        }
    }
}
```


response
```
{
    "event_type": "DEL_ALL_FOCUS",
    "focus_members": [
        {
            "name": "hoge1",
            "connects": []
        },
        {
            "name": "hoge2",
            "connects": []
        }
    ],
    "effect_member": {
        "name": "",
        "type": ""
    }
}
```

<br>

### POST /room

新しい部屋を作成する

ここで作成した部屋以外ではイベント(Focus,Effect,...)を実行できない

response

```
{
    "data": {
        "id": "cp4p"
    }
}
```

<br>

### DELETE /room/:room_id

サーバーで管理している部屋の情報を削除する

部屋を閉じるタイミングでたたく

正常

response
```
{
    "ok": "delete hub of roomId successful"
}
```

<br>

指定したidの部屋がなかった場合

response
```
{
    "error": "Hubs.CheckAndDeleteHubOfRoomId Error : roomId not found in Hubs"
}
```

<br>


## ScreenShot

### SetScreenShot

スクリーンショットの開始を通知する

send
```
const data = {
    type:"SET_SCREEN_SHOT"
}
```


response
```
{
    "event_type": "SET_SCREEN_SHOT",
    "focus_members": [
        {
            "name": "hoge1",
            "connects": []
        },
        {
            "name": "hoge2",
            "connects": []
        }
    ],
    "effect_member": {
        "name": "",
        "type": ""
    }
}
```

<br>



## Effect

### SetEffect

スクリーンショットの開始を通知する

send
```
const data = {
    type:"SET_EFFECT",
    info:{
        effect:{
            name:"hoge1",
            type:"happy"
        }
    }
}
```


response
```
{
    "event_type": "SET_EFFECT",
    "focus_members": [
        {
            "name": "hoge1",
            "connects": []
        },
        {
            "name": "hoge2",
            "connects": []
        }
    ],
    "effect_member": {
        "name": "hoge1",
        "type": "happy"
    }
}
```

<br>