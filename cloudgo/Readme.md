# CloudGo
cloudgo是第二次作业的成果，具体的代码和老师的类似，但是并没有直接的复制粘贴，而是自己在阅读文档的基础上完成的，下面将讲述一些相关的东西
## 框架使用情况
[gorilla/mux](github.com/gorilla/mux)、[codegangsta/negroni]("github.com/codegangsta/negroni")、[unrolled/render](github.com/unrolled/render)
以上的框架使用并没有什么特别的理由，仅仅是因为在老师的课件看到了就拿来用了，仅此而已
## 具体服务端代码编写流程
创建新服务器->添加路由->运行服务器
以上的流程是我在完成代码编写的基础上总结出来的，所谓的创建新服务器，也就是利用negroni创建出服务器的过程，而创建的服务器需要处理不同的路径，因此可以采用mux来替换默认的mux，从而扩展功能，最后也就是运行，指定端口运行即可
## 测试
### curl
### ab