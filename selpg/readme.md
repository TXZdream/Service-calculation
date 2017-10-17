# 第一次作业
## selpg
### 作业要求
详情见[链接](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html)
### 说明
1. 和链接上有所不同的是，这里的命令行的参数并不是直接接在flag之后的，而是类似这样的格式`selpg -s 1 -e 2 -l 3 main.go`或者加上等于号，其他的相同
2. log文件夹下为日志文件，请勿删除这个文件夹或下面的文件，否则将无法写出日志
3. `-d`命令后接目标文件的名称，即最后输出处理后的内容的位置
### 结构介绍
- lib/selpg
该目录下定义了`Selpg`这个结构体，所有的操作都是在这个结构体上进行的，内有`Read、Write、Print`函数，专门用来处理不同的flag
- log
该目录是存放日志的目录，相关的日志将会在该目录下的log.txt中输出
- main.go
主程序，负责类的创建以及调用相关的函数