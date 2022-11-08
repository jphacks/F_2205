# API仕様書

APIの仕様を直接触って確かめたい場合は

以下のコマンドでAPIのテストサーバーを起動する

```shell
cd ./server && make test-api
```

<br>

## Room
### POST /room

新しい部屋を作成する

ここで作成した部屋以外ではイベント(Focus,Effect,...)を実行できない

response
```json
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

response
```json
{
    "ok": "delete room of roomId successful"
}
```

<br>

### GET /ws/:room_id (AddNewMember)

roomにmemberを新規追加する

focus情報にもmemberを新規追加する

またfocusと違いmembersは辞書型なので注意すること

時間があればfocusも辞書型にしたい

send
```json
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
```json
{
    "event_type": "ADD_NEW_MEMBER",
    "members": {
        "peer_hoge1": {
            "name": "hoge1",
            "isRestRoom": false
        }
    },
    "focus_members": [
        {
            "peer_id": "peer_hoge1",
            "connects": []
        }
    ],
    "effect_member": {
        "peer_id": "",
        "type": ""
    },
    "sound": {
        "type": ""
    }
}
```

<br>

## Hub
### DELETE /ws/:room_id
サーバーで管理しているHubの情報を削除する

部屋を閉じるタイミングでたたく

response
```json
{
    "ok": "delete hub of roomId successful"
}
```

<br>


## Focus

### GET /ws/:room_id (SetFocus)

to,fromにmember名をいれて、ユーザー同士を接続状態にする

fromさんとtoさんが存在するものと仮定する

send
```json
{
    "type": "SET_FOCUS",
    "info": {
        "focus": {
            "from": "peer_hoge1",
            "to": "peer_hoge2"
        }
    }
}
```

response
```json
{
    "event_type": "SET_FOCUS",
    "members": {
        "peer_hoge1": {
            "name": "hoge1",
            "isRestRoom": false
        },
        "peer_hoge2": {
            "name": "hoge2",
            "isRestRoom": false
        }
    },
    "focus_members": [
        {
            "peer_id": "peer_hoge1",
            "connects": [
                {
                    "peer_id": "peer_hoge2"
                }
            ]
        },
        {
            "peer_id": "peer_hoge2",
            "connects": [
                {
                    "peer_id": "peer_hoge1"
                }
            ]
        }
    ],
    "effect_member": {
        "peer_id": "",
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
```json
{
    "type": "DEL_FOCUS",
    "info": {
        "focus": {
            "from": "peer_hoge1",
            "to": "peer_hoge2"
        }
    }
}
```


response
```json
{
    "event_type": "DEL_FOCUS",
    "members": {
        "peer_hoge1": {
            "name": "hoge1",
            "isRestRoom": false
        },
        "peer_hoge2": {
            "name": "hoge2",
            "isRestRoom": false
        }
    },
    "focus_members": [
        {
            "peer_id": "peer_hoge1",
            "connects": []
        },
        {
            "peer_id": "peer_hoge2",
            "connects": []
        }
    ],
    "effect_member": {
        "peer_id": "",
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
```json
{
    "type": "DEL_ALL_FOCUS",
    "info": {
        "focus": {
            "from": "peer_hoge1"
        }
    }
}
```


response
```json
{
    "event_type": "DEL_ALL_FOCUS",
    "members": {
        "peer_hoge1": {
            "name": "hoge1",
            "isRestRoom": false
        },
        "peer_hoge2": {
            "name": "hoge2",
            "isRestRoom": false
        }
    },
    "focus_members": [
        {
            "peer_id": "peer_hoge1",
            "connects": []
        },
        {
            "peer_id": "peer_hoge2",
            "connects": []
        }
    ],
    "effect_member": {
        "peer_id": "",
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
```json
{
    "type": "SET_SCREEN_SHOT"
}
```


response
```json
{
    "event_type": "SET_SCREEN_SHOT",
    "members": {
        "peer_hoge1": {
            "name": "hoge1",
            "isRestRoom": false
        },
        "peer_hoge2": {
            "name": "hoge2",
            "isRestRoom": false
        }
    },
    "focus_members": [
        {
            "peer_id": "peer_hoge1",
            "connects": []
        },
        {
            "peer_id": "peer_hoge2",
            "connects": []
        }
    ],
    "effect_member": {
        "peer_id": "",
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
```json
{
    "type": "SET_EFFECT",
    "info": {
        "effect": {
            "peer_id": "peer_hoge1",
            "type": "happy"
        }
    }
}
```


response
```json
{
    "event_type": "SET_EFFECT",
    "members": {
        "peer_hoge1": {
            "name": "hoge1",
            "isRestRoom": false
        },
        "peer_hoge2": {
            "name": "hoge2",
            "isRestRoom": false
        }
    },
    "focus_members": [
        {
            "peer_id": "peer_hoge1",
            "connects": []
        },
        {
            "peer_id": "peer_hoge2",
            "connects": []
        }
    ],
    "effect_member": {
        "peer_id": "peer_hoge1",
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
```json
{
    "type": "SET_SOUND",
    "info": {
        "sound": {
            "type": "hoge_sound"
        }
    }
}
```

response
```json
{
    "event_type": "SET_SOUND",
    "members": {
        "peer_hoge1": {
            "name": "hoge1",
            "isRestRoom": false
        },
        "peer_hoge2": {
            "name": "hoge2",
            "isRestRoom": false
        }
    },
    "focus_members": [
        {
            "peer_id": "peer_hoge1",
            "connects": []
        },
        {
            "peer_id": "peer_hoge2",
            "connects": []
        }
    ],
    "effect_member": {
        "peer_id": "",
        "type": ""
    },
    "sound": {
        "type": "hoge_sound"
    }
}
```

<br>


## RestRoom

### SetRestRoom
誰がお手洗い行ってるかを設定する

peerId がpeer_hoge1の人がお手洗いに行っている状態にする

send
```json
{
    "type": "SET_REST_ROOM",
    "info": {
        "rest": {
            "peer_id": "peer_hoge1",
            "is_rest_room": true
        }
    }
}
```

response
```json
{
    "event_type": "SET_REST_ROOM",
    "members": {
        "peer_hoge1": {
            "name": "hoge1",
            "isRestRoom": true
        },
        "peer_hoge2": {
            "name": "hoge2",
            "isRestRoom": false
        }
    },
    "focus_members": [
        {
            "peer_id": "peer_hoge1",
            "connects": []
        },
        {
            "peer_id": "peer_hoge2",
            "connects": []
        }
    ],
    "effect_member": {
        "peer_id": "",
        "type": ""
    },
    "sound": {
        "type": ""
    }
}
```

<br>

お手洗いから戻ってきた状態にする
send
```json
{
    "type": "SET_REST_ROOM",
    "info": {
        "rest": {
            "peer_id": "peer_hoge1",
            "is_rest_room": false
        }
    }
}
```

response
```json
{
    "event_type": "SET_REST_ROOM",
    "members": {
        "peer_hoge1": {
            "name": "hoge1",
            "isRestRoom": false
        },
        "peer_hoge2": {
            "name": "hoge2",
            "isRestRoom": false
        }
    },
    "focus_members": [
        {
            "peer_id": "peer_hoge1",
            "connects": []
        },
        {
            "peer_id": "peer_hoge2",
            "connects": []
        }
    ],
    "effect_member": {
        "peer_id": "",
        "type": ""
    },
    "sound": {
        "type": ""
    }
}
```