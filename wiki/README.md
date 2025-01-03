## 接口文档
### GET /list 获取文档列表
响应体data部分：
```json
[
  {
    "id": 1,
    "path": "文档路径",
    "title": "文档标题"
  }
]
```
### GET /get/:id 获取特定文档的内容
响应体data部分：
```json
"文档内容"
```
### POST /add 添加一个文档
* 要求JWT且具有管理员权限
请求体：
```json
{
  "path": "文档路径",
  "title": "文档标题",
  "content": "文档内容"
}
```
### POST /edit 修改一个已有的文档
* 要求JWT且具有管理员权限
* id是必须提供的，path和title和content是可选的，只更新提供的数据
请求体：
```json
{
  "id": 1,
  "path": "文档路径",
  "title": "文档标题",
  "content": "文档内容"
}
```
