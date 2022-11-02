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

## 詳しく

### AddNewMember

roomにmemberを新規追加する

fromさんを接続情報に新規追加する

send
```
const data = {
    type:"NEW_MEMBER",
    info:{
        from:"hoge1",
    }
}
```

response
```
{
    "focus": {
        "members": [
            {
                "name": "hoge1",
                "Connects": []
            }
        ]
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
        from:"hoge1",
        to:"hoge11",
    }
}
```

response
```
{
    "focus": {
        "members": [
            {
                "name": "hoge1",
                "Connects": [
                    {
                        "name": "hoge11"
                    }
                ]
            },
            {
                "name": "hoge11",
                "Connects": [
                    {
                        "name": "hoge1"
                    }
                ]
            }
        ]
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
        from:"hoge1",
        to:"hoge11",
    }
}
```


response
```
{
    "focus": {
        "members": [
            {
                "name": "hoge1",
                "Connects": []
            },
            {
                "name": "hoge11",
                "Connects": []
            }
        ]
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
        from:"hoge1",
    }
}
```


response
```
{
    "focus": {
        "members": [
            {
                "name": "hoge1",
                "Connects": []
            },
            {
                "name": "hoge11",
                "Connects": []
            }
        ]
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
