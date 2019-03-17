# moment@interface
# https://moment.siiva.com

###  一.账户
#### 1.手机+密码登录

> POST /login

`request body`

```json
{
  "phone":"18888888888",
  "passwd":"123456"
}
```

`成功返回`

```json
{
  "access_token": "2c0dbf809448ce0801e39a5309c697fa",
  "token_type": "Bearer"
}
```

#### 2.获取登录验证码

> GET /sendcode/:phone

`成功返回`

```
{
  "info": "已发送"
}
```

#### 3.微信登录

> GET /login/wechat/:code

`成功返回`

```json
{
  "access_token": "2c0dbf809448ce0801e39a5309c697fa",
  "token_type": "Bearer"
}
```

#### 4.手机+验证码登录

> POST /login/code

`request body`

```json
{
  "phone":"18888888888",
  "code":"5593"
}
```

`成功返回`

```json
{
  "access_token": "2c0dbf809448ce0801e39a5309c697fa",
  "token_type": "Bearer"
}
```


###  二.活动
#### 1.创建活动

> POST /activity

`request body`

```json
{
  "group":"可选项",
  "activity_name": "半场篮球",
  "address": "北京市",
  "mark": "3v3半场篮球"
}
```

`成功返回`

```
{
  "activity_id": "1540193045",
  "siteIds": [
    "1540193045ea"
  ]
}
```

#### 2.更新活动

> POST /activity/update/:activity_id

`request body`

```json
{
  "group":"可选项",
  "activity_name": "半场篮球",
  "address": "北京市",
  "mark": "3v3半场篮球"
}
```

`成功返回`

```
{
  "info": "更新成功"
}
```

#### 3.活动list

> GET /activity/list

`成功返回`

```json
[
  {
    "activity_id": "1539762286",
    "activity_name": "半场篮球",
    "address": "北京市",
    "mark": "3v3半场篮球",
    "create_time": "2018-10-17 15:44"
},
    ...
]
```

###  三.场地
#### 1.添加场地

> POST /site

`request body`

```json
{
  "activity_id":"623324123",
  "siteName": "左半场",
  "siteType": "moment",
  "group": "可选"
}
```

`成功返回`

```
{
  "info":"添加成功"
}
```






