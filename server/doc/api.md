# API仕様書

## Room
### POST /room

新しい部屋を作成する

ここで作成した部屋以外ではイベント(Focus,Effect,...)を実行できない

response

```
{
    "data": {
        "id": "1111"
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


## Focus

### GET /ws/:room_id (AddNewMember)

roomにmemberを新規追加する

focus情報にもmemberを新規追加する

send
```
{
    "type": "ADD_NEW_MEMBER",
    "info": {
        "member": {
            "name": "hoge1",
            "peer_id": "peer_hoge1"
        }
    }
}
```

response
```
{
    "event_type": "ADD_NEW_MEMBER",
    "members": [
        {
            "peer_id": "peer_hoge1",
            "name": "hoge1"
        }
    ],
    "focus_members": [
        {
            "name": "hoge1",
            "connects": []
        }
    ],
    "effect_member": {
        "name": "",
        "type": ""
    },
    "sound": {
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
{
    "type": "SET_FOCUS",
    "info": {
        "focus": {
            "from": "hoge1",
            "to": "hoge2"
        }
    }
}
```

response
```
{
    "event_type": "SET_FOCUS",
    "members": [
        {
            "peer_id": "peer_hoge1",
            "name": "hoge1"
        },
        {
            "peer_id": "peer_hoge2",
            "name": "hoge2"
        }
    ],
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
    },
    "sound": {
        "type": ""
    }
}
```

<br>

### GET /ws/:room_id (DelFocus)

to,fromにmember名をいれて、ユーザー同士の接続状態を解除する

send
```
{
    "type": "DEL_FOCUS",
    "info": {
        "focus": {
            "from": "hoge1",
            "to": "hoge2"
        }
    }
}
```


response
```
{
    "event_type": "DEL_FOCUS",
    "members": [
        {
            "peer_id": "peer_hoge1",
            "name": "hoge1"
        },
        {
            "peer_id": "peer_hoge2",
            "name": "hoge2"
        }
    ],
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
    },
    "sound": {
        "type": ""
    }
}
```

<br>

### GET /ws/:room_id (DelAllFocus)

fromにmember名をいれて、そのユーザーの持っている接続状態を解除する

send
```
{
    "type": "DEL_ALL_FOCUS",
    "info": {
        "focus": {
            "from": "hoge1"
        }
    }
}
```


response
```
{
    "event_type": "DEL_ALL_FOCUS",
    "members": [
        {
            "peer_id": "peer_hoge1",
            "name": "hoge1"
        },
        {
            "peer_id": "peer_hoge2",
            "name": "hoge2"
        }
    ],
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
    },
    "sound": {
        "type": ""
    }
}
```

<br>

## ScreenShot

### SetScreenShot

スクリーンショットの開始を通知する

send
```
{
    "type": "SET_SCREEN_SHOT"
}
```


response
```
{
    "event_type": "SET_SCREEN_SHOT",
    "members": [
        {
            "peer_id": "peer_hoge1",
            "name": "hoge1"
        },
        {
            "peer_id": "peer_hoge2",
            "name": "hoge2"
        }
    ],
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
    },
    "sound": {
        "type": ""
    }
}
```

<br>



## Effect

### SetEffect

どのメンバーが何のエフェクトを使ったのかを通知する

send
```
{
    "type": "SET_EFFECT",
    "info": {
        "effect": {
            "name": "hoge1",
            "type": "happy"
        }
    }
}
```


response
```
{
    "event_type": "SET_EFFECT",
    "members": [
        {
            "peer_id": "peer_hoge1",
            "name": "hoge1"
        },
        {
            "peer_id": "peer_hoge2",
            "name": "hoge2"
        }
    ],
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
    },
    "sound": {
        "type": ""
    }
}
```

<br>

## Sound

### SetSound
どの音楽を流すか通知する

send
```
{
    "type": "SET_SOUND",
    "info": {
        "sound": {
            "type": "hoge1"
        }
    }
}
```

response
```
{
    "event_type": "SET_SOUND",
    "members": [
        {
            "peer_id": "peer_hoge1",
            "name": "hoge1"
        },
        {
            "peer_id": "peer_hoge2",
            "name": "hoge2"
        }
    ],
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
    },
    "sound": {
        "type": "hoge1"
    }
}
```