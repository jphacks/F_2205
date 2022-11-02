# API仕様書

## 実装済みAPIについて
- AddNewMember
  - roomにmemberを新規追加する
  - 部屋に参加したタイミングで使う

- SetFocus
  - to,fromにmember名をいれて、ユーザー同士を接続状態にする

- DelFocus
  - to,fromにmember名をいれて、ユーザー同士の接続状態を解除する

- DelAllFocus
  - fromにmember名をいれて、そのユーザーの持っている接続状態を解除する
  - from以外のmemberがfromとの接続情報を持っていた場合、その情報も削除する

- DelRoom
  - 部屋の持っている接続情報をすべて削除する
  - 部屋に誰もいなくなったタイミングで使用し、serverのオンメモリキャッシュを消す

<br>

## Focus

### AddNewMember

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

### SetFocus

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

### DelFocus

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

### DelAllFocus

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


### DelRoom

部屋の持っている接続情報をすべて削除する

DELETE /ws/:room

response
```
{
    "ok": "delete room successful"
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