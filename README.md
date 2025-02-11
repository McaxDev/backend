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
    "account": "邮箱/用户名/手机号/基岩版名称/Java版名称",
    "password": "用户密码"
}
```
* 响应体data部分：
```json
"Bearer 用户凭证"
```
#### 注册 POST
* 路径：`/account/signup`
* 请求体：
```json
{
    "captchaId": "人机验证ID",
    "captchaCode": "人机验证码",
    "emailId": "邮箱账号",
    "emailCode": "邮箱验证码",
    "username": "用户名",
    "password": "用户密码"
}
```
* 响应体data部分：
```json
"Bearer 用户凭证"
```
#### 获取我的信息 GET
* 路径：`/account/get/myinfo`
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
    "javaName": "Nerakolo"
}
```
#### 获取其他用户的信息 GET
* 路径（id是用户id）：`/account/get/userinfo?id=1`
* 响应体同上
#### 获取用户设置 GET
* 路径：`/account/get/settings`
* 请求头带JWT
* 响应体data部分：
```json
[
    {
        "id": "UseMFA",
        "name": "启用MFA验证",
        "value": true
    },
    {
        "id": "PubEmail",
        "name": "公开我的邮箱",
        "value": false
    }
]
```
#### 更新用户设置 POST
* 路径：`/account/set/setting`
* 请求头带JWT
* 请求体（index代表该项在"获取"里的位置，从0开始）：
```json
{
    "index": 1,
    "value": true
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
{
    "phoneId": "手机号码",
    "phoneCode": "短信验证码"
}
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
{
    "emailId": "1285607932@qq.com",
    "emailCode": "114514"
}
```
#### 更新用户名 POST
* 路径：`/account/set/username`
* 请求头带JWT
* 请求体：
```json
"用户名字符串"
```
#### 更新密码 POST
* 路径：`/account/set/password`
* 需要验证邮箱
* 请求体：
```json
    "emailId": "asdf@abc.com",
    "emailCode": "123456",
    "password": "114514"
```
#### 更新用户信息 POST
* 路径：`/account/set/userinfo`
* 请求头带JWT
* 请求体：
```json
{
    "avatar": "image.png",
    "profile": "鸡你太美"
}
```
#### 获取黑名单列表 GET
* 路径：`/account/get/blacklist`
* 响应体data部分：
```json
{
    "phone": ["12312341234", "12323452345"],
    "email": ["qwer@abc.com", "asfd@bcd.com"]
}
```
#### 更新黑名单 POST
* 路径：`/account/edit/blacklist`
* 需要管理员权限
* 请求体data部分：
```json
{
    "id": "记录ID，带为编辑已有记录，不带为新增记录",
    "type": "账号类型，如email或phone",
    "value": "值号码",
    "expiry": "解封时间，使用ISO8601格式"
}
```
#### 删除黑名单 DELETE
* 路径：`/account/del/blacklist`
* 需要管理员权限
* 请求体data部分（123代表黑名单记录的ID）：
```json
123
```
#### 获取wiki列表
* 路径：`/wiki/list`
* 响应体data部分：
```json
[
        {
            "category": "atype",
            "id": 1,
            "path": "a",
            "title": "测试A",
        },
        {
            "category": "btype",
            "id": 2,
            "path": "b",
            "title": "测试B",
        }
    ]
```
#### 获取特定wiki信息
* 路径：`/wiki/get?id=?`
* 响应体data部分：
```json
{
        "category": "atype",
        "content": "测试A",
        "createdAt": "2025-01-12T04:58:10Z",
        "html": "测试A",
        "id": 1,
        "path": "a",
        "title": "测试A",
        "updatedAt": "2025-01-12T04:58:10Z"
}
```
#### 编辑特定wiki信息
* 路径：`/wiki/edit`
* 需要管理员权限
* 请求体（带id为编辑，不带id为新增）：
* 不需要传入html字段，html字段由后端从content渲染
```json
{
        "id": 1,
        "category": "atype",
        "content": "测试A",
        "path": "a",
        "title": "测试A",
}
```
#### 删除特定wiki
* 路径：`/wiki/delete`
* 需要管理员权限
* 请求体（数字12代表wiki的id）：
```json
12
```
#### 查看相册列表 GET
* 路径：`/gallery/get/albums`
#### 查看相册的图片列表 GET
* 路径（10是相册ID）：`/gallery/get/images?id=10`
#### 创建相册 POST
* 路径：`/gallery/add/album`
* 请求头带JWT
* 请求体：
```json
{
    "path": "相册url路径",
    "title": "相册标题"
}
```
#### 上传图片 POST
* 路径：`/gallery/add/image`
* 请求头带JWT
* 请求体MIME类型为`multipart/form-data`
* 请求体字段列表：
  * `album`: 代表相册ID，如`10`
  * `image`: 图片文件
  * `title`: 图片标题
  * `description`: 图片描述
#### 修改相册信息 POST
* 对于修改或删除相册，web管理员具有所有相册权限，相册创建者具有他创建的相册的权限，如果相册属于公会，那么公会管理员有该相册的权限
* 路径：`/gallery/set/album`
* 请求头带JWT
#### 修改图片信息 POST
* 路径：`/gallery/set/image`
* 请求头带JWT
#### 删除相册 DELETE
* 路径：`/gallery/del/album`
* 请求头带JWT
* 请求体（里面数字是相册ID）：
```json
3
```
#### 删除图片 DELETE
* 路径：`/gallery/del/images`
* 请求头带JWT
* 请求体（里面数字是图片ID）：
```json
[1, 2, 3]
```
#### 获取特定玩家的统计信息 GET
* 路径：`/game/player/:player_name`
* 查询字符串参数：`server=服务器ID`
* 响应体data部分：见`docs/result.json`
#### 获取特定统计信息的排行榜 GET
* 路径：`/game/rank/:stat`，`:stat`包括`mined`,`picked_up`,`crafted`,`broken`,`play_time`,`deaths`,`mob_kills`,`damage_dealt`,`drop`
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
* 需要管理员权限
* 请求体：
```json
{
    "Server": "服务器ID",
    "Command": "/list"
}
```
#### 发送绑定MCJE的验证码 POST
* 路径：`/game/bind/je`
* 需要请求头带JWT
* 请求体：
```json
"用户名"
```
#### 验证绑定验证码 POST
* 路径：`/game/auth/bind`
* 需要请求头带JWT
* 请求体：
```json
{
    "gamename": "玩家名",
    "authcode": "验证码内容"
}
```
#### 查看在线玩家 GET
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
#### 重命名公会 POST
* 路径：`/guild/rename`
* 请求头带JWT，要求公会角色3或4
* 请求体：
```json
"新公会名"
```
#### 申请加入公会 POST
* 路径：`/guild/join`
* 请求头带JWT，要求公会角色0
* 请求体（123代表目标公会ID）：
```json
123
```
#### 退出公会 POST
* 路径：`/guild/leave`
* 请求头带JWT，要求公会角色1或2或3
#### 审核入会申请 POST
* 路径：`/guild/review`
* 请求头带JWT，要求公会角色2或3
* 请求体（ids代表用户id，agree代表同意或拒绝）：
```json
{
    "ids": [1, 2, 3],
    "agree": true
}
```
#### 任命公会管理员 POST
* 路径：`/guild/appoint`
* 请求头带JWT，要求公会角色为4
* 请求体：
```json
{
    "ids": [1, 2, 3],
    "agree": false
}
```
#### 创建公会 POST
* 路径：`/guild/create`
* 请求头带JWT，要求公会角色为0
* 请求体：
```json
{
    "name": "理塘 - Leetown",
    "profile": "到达世界最高城，理塘"
}
```
#### 转让公会 POST
* 路径：`/guild/transfer`
* 请求头带JWT，要求公会角色为4
* 请求体（5代表新会长id）：
```json
5
```
#### 解散公会 POST
* 路径：`/guild/dissolve`
* 请求头带JWT，要求公会角色为4
#### 向公会捐赠 POST
* 路径：`/guild/donate`
* 请求头带JWT
* 要求公会角色至少为2或3或4
* 请求体（里面的数字是贡献币）：
```json
123
```
#### 从公会提取捐赠 POST
* 路径：`/guild/withdraw`
* 请求头带JWT
* 要求公会角色为3或4
* 请求体（里面的数字是贡献币）：
```json
123
```
#### 升级公会 POST
* 路径：`/guild/upgrade`
* 请求头带JWT
* 要求公会角色为3或4
#### 获取帖子列表 GET
* 路径：`/bbs/get/posts?category=bedrock`
* 响应体data部分：
```json
[
    {
        "id": 1,
        "createdAt": "创建时间",
        "updatedAt": "更新时间",
        "title": "标题",
        "category": "分类",
        "guild": {},
        "user": {},
    }
]
```
#### 获取特定帖子内容 GET
* 路径：`/bbs/get/post?id=1`
* 响应体data部分：
```json
{
    "id": 1,
    "createdAt": "创建时间",
    "updatedAt": "更新时间",
    "title": "标题",
    "category": "分类",
    "source": "源MD内容",
    "content": "编译后html内容",
    "guild": {},
    "user": {},
    "comments": []
}
```
#### 发帖 POST
* 路径：`/bbs/add/post`
* 请求头带JWT
* 请求体（useMd表示是否用markdown解析源内容）：
```json
{
    "category": "分类",
    "title": "标题",
    "content": "源内容",
    "useMd": true
}
```
#### 发评论 POST
* 路径：`/bbs/add/comment`
* 请求头带JWT
* 请求体（useMd表示是否用markdown解析源内容，id是帖子id）：
* attitude表示态度，-1,0,1分别代表反对、无意见，支持。
```json
{
    "id": 10,
    "content": "帖子内容",
    "useMd": false,
    "attitude": 0
}
```
#### 改帖子 POST
* 路径：`/bbs/set/post`
* 请求头带JWT
* 请求体：
* useMd表示是否用markdown解析，只需要提供需要修改的字段
```json
{
    "id": 10,
    "category": "分类",
    "title": "标题",
    "content": "源内容",
    "useMd": false
}
```
#### 改评论 POST
* 路径：`/bbs/set/comment`
* 请求头带JWT
* useMd表示是否用markdown解析，只需要提供需要修改的字段
* 请求体：
```json
    "id": 10,
    "attitude": 0,
    "content": "评论内容",
    "useMd": false
```
#### 删帖子 DELETE
* 路径：`/bbs/del/posts`
* 请求头带JWT
* 请求体（提供帖子ID列表）：
```json
[0, 1, 2]
```
#### 删评论 DELETE
* 路径：`/bbs/del/comments`
* 请求头带JWT
* 请求体（提供评论ID列表）：
```json
[0, 1, 2]
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
