## 开发文档
### 后端相关
#### 架构
* 后端采用微服务架构，具体微服务如下：
  * auth：为其他微服务提供Captcha、Email、SMS验证接口。
  * gameapi：负责提供与游戏服务器交互的接口。
  * account：只负责Web应用的账号相关。
    * 依赖于auth微服务。
  * usergame：负责让Web用户与游戏服务器交互。
    * 依赖于gameapi和account和auth微服务。
  * test：负责处理入服问卷相关。
    * 依赖于auth和account微服务。
  * dash：负责提供玩家数据查询服务。
    * 依赖于account微服务。
### HTTP协议注释
* 有一些规则是不断重复说明的，在这一部分列出。
#### 请求头带JWT
* 请求头加上`Authorization: Bearer 234f62.236f3245.345vy3`
#### 服务器ID 枚举类型
* 互通服是`paper`，生电服是`fabric`，基岩服是`bedrock`，饥荒服是`dst`，星露谷物语服是`stardew`，泰拉瑞亚服是`terraria`
#### 需要管理员权限
* 要求提供JWT的同时，用户的个人信息的`admin`为true
#### 需要验证邮箱/手机/人机
* 先申请验证码，再在请求体里面带上`captchaId/captchaCode`，`emailId/emailCode`，`phoneId/phoneCode`
#### 里面带`:`的路径
* 表示那个路径是变量，例如`/account/set/:email`可以是`/account/set/nerakolo@outlook.com`，变量的意义在下面列出。
* `email`表示邮箱，`telephone`表示手机号。
### HTTP协议（子域名为api.mcax.cn）
* 请求体都是application/json类型（少数如上传文件等除外），而且对象的键，大小写皆可。例如，如果文档写的键是`Username`，那么`username`或`UserName`或`USERNAME`皆可，但`user_name`就不行。
* 响应体默认是此结构，后文只解释`data`部分：
```json
{
    "message": "请求成功",
    "error": "错误描述",
    "data": 可能是任何类型
}
```
* 如果一个接口没有提及响应体，那么响应体一般是：
```json
{
    "message": "请求成功/请求失败：原因",
    "error": "错误描述",
    "data": null
}
```
* 如果一个接口没有提及请求体，那么一般不需要请求体。
#### 获取Captcha验证码 GET
* 路径：`/auth/captcha`
* 响应头：`X-Captcha-Id: jc8u9wty8jcw90t35`、`Content-Type: image/png`
#### 获取邮件验证码 GET
* 路径：`/auth/email/1285607932@qq.com`
#### 获取手机验证码 GET
* 需要JWT，需要3任意币（短信也要钱）
* 路径：`/verify/phone/:telephone`
#### 注销账号 POST
* 路径：`/account/signout`
* 请求头带JWT
* 需要验证邮箱
#### 登录 POST
* 路径：`/account/login`
* 请求体：
```json
{
    "Account": "邮箱/手机号/用户名",
    "Password": "用户密码"
}
```
* 响应体data部分：
```json
"Bearer 用户凭证"
```
#### 注册 POST
* 路径：`/account/login`
* 请求体：
```json
{
    "captchaId": "人机验证ID",
    "captchaCode": "人机验证码",
    "emailId": "邮箱账号",
    "emailCode": "邮箱验证码",
    "Username": "用户名",
    "Password": "用户密码"
}
```
* 响应体data部分：
```json
"Bearer 用户凭证"
```
#### 获取用户信息 GET
* 路径：`/account/get/userinfo`
* 请求头带JWT
* 响应体data部分：
```json
{
    "userId": 123,
    "username": "Nerakolo",
    "avatar": "filename.jpg",
    "profile": "Ciallo!",
    "admin": true,
    "money": 100,
    "email": "nerakolo@outlook.com",
    "telephone": "12312341234",
    "bedrockName": "Nerakolox",
    "javaName": "Nerakolo",
    "dstName": "nerakolo"
}
```
#### 获取用户设置 GET
* 路径：`/account/get/settings`
* 请求头带JWT
* 响应体data部分：
```json
[
    {
        "name": "enableMfa",
        "comment": "启用MFA验证",
        "value": true
    },
    {
        "name": "mfaUseEmail",
        "comment": "开启则用Email作为MFA方式，关闭则为SMS",
        "value": false
    }
]
```
#### 更新用户设置 POST
* 路径：`/account/set/settings`
* 请求头带JWT
* 请求体：
```json
{
    "Name": "enableMfa",
    "Value": true
}
```
#### 每日签到 GET
* 路径：`/account/checkin`
* 请求头带JWT
#### 查看签到历史 GET
* 路径：`/account/get/checkin`
* 请求头带JWT
* 响应体data部分：
```json
[
    {"Date": 1, "Status": true},
    {"Date": 2, "Status": false}
]
```
#### 绑定手机 POST
* 路径：`/account/bind/phone`
* 请求头带JWT
* 请求体：
```json
{
    "phoneId": "手机号码",
    "phoneCode": "短信验证码"
}
```
#### 解绑手机 POST
* 路径：`/account/unbind/phone`
* 请求头带JWT
* 请求体：
```json
手机验证码
```
#### 绑定邮箱 POST
* 路径：`/account/bind/email`
* 请求头带JWT
* 请求体：
```json
{
    "emailId": "1285607932@qq.com",
    "emailCode": "114514"
}
```
#### 解绑邮箱 POST
* 路径：`/account/unbind/email`
* 请求头带JWT
* 请求体：
```json
邮箱验证码
```
#### 获取特定玩家的统计信息 GET
* 路径：`/dash/player/:player_name`
* 查询字符串参数：`server=服务器ID`
* 响应体data部分：见`docs/result.json`
#### 获取特定统计信息的排行榜 GET
* 路径：`/dash/:stat`，`:stat`包括`mined`,`picked_up`,`crafted`,`broken`,`play_time`,`deaths`,`mob_kills`,`damage_dealt`,`drop`
* 查询字符串参数：`server=paper`，见上文“服务器ID”
* 响应体data部分：
```json
[
    {
        "Score": 100,
        "Member": "Nerakolo"
    },
    {
        "Score": 90,
        "Member": "Bestcb233"
    }
]
```
#### 立即备份服务器 POST
* 路径：`/game/backup`
* 需要管理员权限
* 请求体：
```json
{
    "Server": "服务器ID，如'fabric'"
}
```
#### 向服务器发送命令 POST
* 路径：`/game/command`
* 除了部分命令外，其他命令都需要管理员权限
* 请求体：
```json
{
    "Server": "服务器ID",
    "Command": "/list"
}
```
* 响应体data部分：
```json
"当前在线1人：Nerakolo"
```
#### 查看公会列表 GET
* 路径：`/guild/get/guilds`
#### 查看我的公会 GET
* 路径：`/guild/get/myguild`
* 请求头带JWT
#### 查看某个公会 GET
* 路径：`/guild/get/guild?id=10`
* URL参数：
  * `id`：公会ID
#### 申请加入公会 POST
#### 退出/解散公会 POST
#### 向公会捐赠 POST
* 路径：`/guild/donate`
* 请求头带JWT
* 要求公会角色至少为2（正式成员）
* 请求体：
```json
金额（数字）
```
#### 从公会提取捐赠 POST
* 路径：`/guild/withdraw`
* 请求头带JWT
* 要求公会角色至少为3（管理员）
* 请求体：
```json
金额（数字）
```
#### 升级公会 POST
* 路径：`/guild/upgrade`
* 请求头带JWT
* 要求公会角色至少为3（管理员）
#### 设置其他成员角色 POST
* 路径：`/guild/set/role`
* 请求头带JWT
* 管理员可以任命正式成员，会长可以任命管理员，公会角色至少为3
* 请求体：
```json
{
    "userIds": [1, 2, 3], // 用户ID列表
    "role": 2 // 角色，1申请者2会员3管理员
}
```
### HTTP协议（子域名为static.mcax.cn）
* 访问`https://static.mcax.cn/`将得到一个文件系统。
#### 响应体格式（JSON）：
```json
[
    {
        "name": "文件夹或文件名",
        "type": "文件类型，包括directory和file",
        "mtime": "最后修改时间",
        "size": "文件大小，文件夹不存在此项"
    },
    {
        "name":"wiki",
        "type":"directory",
        "mtime":"Wed, 02 Oct 2024 02:54:38 GMT"
    },
    {
        "name":"1.png",
        "type":"file",
        "mtime":"Wed, 11 Sep 2024 11:50:41 GMT",
        "size":121659
    }
]
```
#### 文件夹的意义
* /carousel 代表轮播图目录，里面全是图片，图片名代表图片展示的标题文字。
* /dash 代表玩家数据统计项。
* /gallery 代表服务器画廊。里面的每一个目录代表一个相册，里面的metadata.json代表相册的元数据。
* /wiki 代表维基百科目录，里面有若干个md文件，和一个metadata.json代表元数据，其中index.md是默认的。
