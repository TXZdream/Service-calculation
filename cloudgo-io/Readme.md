# cloudgo-io
1. 支持静态文件访问
`/static/` 这里为静态文件夹，在detail.tmpl引入了css文件，在访问时会用到
2. 支持简单js访问
`/json` 这里返回一个json格式的数据，和课上的相同，支持js访问的api接口
3. 提交表单并返回表格
`/` 这是一个form，提交会跳转到`/login`，在这个界面会展示输入信息的表格
4. 501
`/unknown` 返回501信息，仿照http包内置的代码写了一个