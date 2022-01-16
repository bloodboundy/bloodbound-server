# bloodbound-server

a board game involves 6-12 players, this is the server side of online version

# 连接

使用任一标准 websocket 客户端即可

# 消息格式

```json
{
  "id": "{UUID}",
  "type": "{type_enum}",
  "{type_enum}": Any
}
```

id 由生成者保证唯一即可，故客户端发的消息 id 和服务端发的消息 id 重了也没问题

后面就不再重复 id 和 type 字段

# 流程

1. websocket 连接
2. 认证
3. 交互
   1. 对局操作
   2. 游戏操作
4. 结束/断开连接

## 统一报错(S->C)

```json
{
  "id": "{UUID}",
  "type": "ERR",
  "ERR": {
    "id": "处理时出错的msg_id",
    "code": 1,
    "msg": "错误原因"
  }
}
```

## 认证操作

### 注册 `type=REGISTER`

```json
{}
```

返回：

```json
Player
```

报错：

| code | msg           |
| ---- | ------------- |
| 1    | generate uuid |

### 登录 `type=LOGIN`

```json
{
  "player_id": "{UUID}"
}
```

返回：

```json
Player
```

## 对局操作(C->S)

### 新建对局 `type=NEW_GAME`

```json
{}
```

返回：

```json
{
  "id": "UUID",
  "type": "NEW_GAME",
  "NEW_GAME": Game
}
```

### 加入对局 `type=JOIN_GAME`

```json
{
  "player_id": "{UUID}",
  "game_id": "{UUID}"
}
```

返回：

- 会广播给当前对局中所有玩家

```json
{
  "game_id": "{UUID}",
  "player_id": "{UUID}",
  "at": "{a unix timestamp}"
}
```

### 离开对局 `type=LEAVE_GAME`

```json
{
  "player_id": "{UUID}",
  "game_id": "{UUID}"
}
```

返回：

- 广播：对局
- 主动离开和被动掉线都算 leave

```json
{
  "game_id": "{UUID}",
  "player_id": "{UUID}",
  "at": "{a unix timestamp}"
}
```

### 重连对局

```json
{
  "player_id": "{UUID}",
  "game_id": "{UUID}"
}
```

- 不填 game_id 的话就自动搜索可重连的对局

返回：

- 广播：对局

```json
{
  "game_id": "{UUID}",
  "player_id": "{UUID}",
  "at": "{a unix timestamp}"
}
```

## 游戏操作（客户端操作）

tbd
