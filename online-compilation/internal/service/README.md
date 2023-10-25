# Service
一、背景知识

这个是文档的地址，kratos框架的一些使用方法，很清晰`https://go-kratos.dev/docs/`

1.首先，所有的.pb.go的文件都是自动生成的，不能修改，也不需要修改，你可以使用`make all`或是`make api`命令进行自动生成，项目依赖库的下载，可以使用 `go generate ./..`和`go mod tidy`

2.在cmd中除了main函数，还有wire和wire_gen文件，wire是一个依赖注入框架，可以分析代码中的依赖关系，自动生成初始化代码
![img.png](img.png)
只要你的依赖注入没缺少，没写错，就可以生成，点图片上的那个运行就可以了,ProviderSet是每一层注入的那个变量，可以看到里面有哪些东西，分别对应service.go文件，biz.go文件，data.go文件

二、框架

1.helloworld可以不用关心，是自动生成的例子，api/note/note.proto这里面是你可以自己定义的接口，确定好请求方法，请求参数，出参入参，就可以命令行输入`make api`了，你在internal/service/note.go这里面是这些接口的入口位置，需要你自己编写的代码

2.![img_1.png](img_1.png)
你可以从这个注册服务的函数开始看，就知道是怎么一步步调用到service层的，从service层，通过依赖注入进来的noteusecase,就可以调用biz层的业务代码函数实现