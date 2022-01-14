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

# 流程

1. websocket 连接
2. 认证
3. 交互
   1. 对局操作
   2. 游戏操作
4. 结束/断开连接

## 认证操作(C->S)

### 注册 `type=REGISTER`

```json
{}
```

### 登录 `type=LOGIN`

```json
{
  "player_id": "{UUID}"
}
```

## 对局操作(C->S)

### 新建对局 `type=NEW_GAME`

```json
{
  "id": "UUID",
  "type": "NEW_GAME",
  "NEW_GAME": {}
}
```

后面就不再重复 id 和 type 字段

### 离开对局 `type=LEAVE_GAME`

```json
{ "game_id": "{UUID}" }
```

### 重连对局

```json
{ "game_id": "{UUID}" }
```

不填 game_id 的话就自动搜索可重连的对局

## 对局消息(S->C)

### 新建对局成功 `type=NEW_GAME`

```json
{
  "id": "UUID",
  "type": "NEW_GAME",
  "NEW_GAME": Game
}
```

Game 结构请看标准对象一节

后面就不再重复 id 和 type 字段了

### 有玩家加入对局 `type=JOIN_GAME`

- 会广播给当前对局中所有玩家
- 自己加入对局的响应包也是这个 type

```json
{
  "game_id": "{UUID}",
  "player_id": "{UUID}",
  "at": "{a unix timestamp}"
}
```

### 有玩家离开 `type=LEAVE_GAME`

- 主动离开和被动掉线都算 leave

```json
{
  "game_id": "{UUID}",
  "player_id": "{UUID}",
  "at": "{a unix timestamp}"
}
```

### 有玩家重连 `type=RECONN_GAME`

```json
{
  "game_id": "{UUID}",
  "player_id": "{UUID}",
  "at": "{a unix timestamp}"
}
```

## 游戏操作（客户端操作）

tbd
