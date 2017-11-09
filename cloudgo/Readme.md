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
`curl -v http://localhost:3000/hello/tangxz`
```
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 3000 (#0)
> GET /hello/tangxz HTTP/1.1
> Host: localhost:3000
> User-Agent: curl/7.52.1
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=UTF-8
< Date: Thu, 09 Nov 2017 10:18:08 GMT
< Content-Length: 23
< 
{
  "name": "tangxz"
}
* Curl_http_done: called premature == 0
* Connection #0 to host localhost left intact

```
由上面的结果可以看到，返回的json值的确显示了传入的姓名，同时以更加人性化的方式展现给我们
### ab
#### 参考链接
[简书](http://www.jianshu.com/p/43d04d8baaf7)、[官方文档](https://httpd.apache.org/docs/2.4/programs/ab.html)
`ab -n 100 -c 10 http://localhost:3000/hello/tangxz`
```
This is ApacheBench, Version 2.3 <$Revision: 1757674 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient).....done


Server Software:        
Server Hostname:        localhost
Server Port:            3000

Document Path:          /hello/tangxz
Document Length:        23 bytes

Concurrency Level:      10
Time taken for tests:   0.066 seconds
Complete requests:      100
Failed requests:        0
Total transferred:      14600 bytes
HTML transferred:       2300 bytes
Requests per second:    1511.88 [#/sec] (mean)
Time per request:       6.614 [ms] (mean)
Time per request:       0.661 [ms] (mean, across all concurrent requests)
Transfer rate:          215.56 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    2   4.3      1      19
Processing:     0    4   3.8      2      20
Waiting:        0    3   3.6      2      18
Total:          0    6   5.5      4      20

Percentage of the requests served within a certain time (ms)
  50%      4
  66%      8
  75%      8
  80%      9
  90%     20
  95%     20
  98%     20
  99%     20
 100%     20 (longest request)
```
从上面的结果来看，我们服务器的平均响应时间为6.614ms，每秒可以处理1511个请求