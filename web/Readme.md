# net/http包源码追踪
这次的博客我们将从一个最基本的web服务器开始，逐渐去追踪并了解源代码的处理过程
## 一个简单的web server
```
package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World.")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":3000", nil)
}
```
代码很简单，只提供了根目录的路由函数，这里不再细说，下面继续我们的源代码阅读，去探索实现原理
## http.HandleFunc源代码追踪
```
|---------------|         |--------------------------|         |----------|
|http.HandleFunc| ------> |DefaultServeMux.HandleFunc| ------> |mux.Handle|
|---------------|         |--------------------------|         |----------|
```
```
func (mux *ServeMux) Handle(pattern string, handler Handler) {
	mux.mu.Lock()
	defer mux.mu.Unlock()

	if pattern == "" {
		panic("http: invalid pattern " + pattern)
	}
	if handler == nil {
		panic("http: nil handler")
	}
	if mux.m[pattern].explicit {
		panic("http: multiple registrations for " + pattern)
	}

	if mux.m == nil {
		mux.m = make(map[string]muxEntry)
	}
	mux.m[pattern] = muxEntry{explicit: true, h: handler, pattern: pattern}

	if pattern[0] != '/' {
		mux.hosts = true
	}

	// Helpful behavior:
	// If pattern is /tree/, insert an implicit permanent redirect for /tree.
	// It can be overridden by an explicit registration.
	n := len(pattern)
	if n > 0 && pattern[n-1] == '/' && !mux.m[pattern[0:n-1]].explicit {
		// If pattern contains a host name, strip it and use remaining
		// path for redirect.
		path := pattern
		if pattern[0] != '/' {
			// In pattern, at least the last character is a '/', so
			// strings.Index can't be -1.
			path = pattern[strings.Index(pattern, "/"):]
		}
		url := &url.URL{Path: path}
		mux.m[pattern[0:n-1]] = muxEntry{h: RedirectHandler(url.String(), StatusMovedPermanently), pattern: pattern}
	}
}
```
从上面的代码和追踪过程我们可以看到调用http.HandFunc的过程，最后一个函数的内容如上，即添加路由，由于内容比较简单，这里不再描述，接下来我们来看提到的另外一个函数
## http.ListenAndServe源代码追踪
```
|-------------------|         |---------------------|         |-------------|
|http.ListenAndServe| ------> |server.ListenAndServe| ------> |server.Server|
|-------------------|         |---------------------|         |-------------|
```
```
func (srv *Server) Serve(l net.Listener) error {
	defer l.Close()
	if fn := testHookServerServe; fn != nil {
		fn(srv, l)
	}
	var tempDelay time.Duration // how long to sleep on accept failure

	if err := srv.setupHTTP2_Serve(); err != nil {
		return err
	}

	srv.trackListener(l, true)
	defer srv.trackListener(l, false)

	baseCtx := context.Background() // base is always background, per Issue 16220
	ctx := context.WithValue(baseCtx, ServerContextKey, srv)
	for {
		rw, e := l.Accept()
		if e != nil {
			select {
			case <-srv.getDoneChan():
				return ErrServerClosed
			default:
			}
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				srv.logf("http: Accept error: %v; retrying in %v", e, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return e
		}
		tempDelay = 0
		c := srv.newConn(rw)
		c.setState(c.rwc, StateNew) // before Serve can return
		go c.serve(ctx)
	}
}
```
由上面的代码我们可以看到，服务器为每一个请求新建一个连接，同时开一个goroutine来进行逻辑的处理，我们对于源代码的追踪就到此为止，而具体的一个流程就如上面所说的，当有新的请求到来时，并不会阻塞这一过程，也就提供了高并发的处理
`serverHandler{c.server}.ServeHTTP(w, w.req)`
这里摘抄Server.serve的一部分，我们继续追下去
```
|------------------------|
|serverHandler.ServerHTTP|
|------------------------|
```
```
func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request) {
	handler := sh.srv.Handler
	if handler == nil {
		handler = DefaultServeMux
	}
	if req.RequestURI == "*" && req.Method == "OPTIONS" {
		handler = globalOptionsHandler{}
	}
	handler.ServeHTTP(rw, req)
}
```
我们可以看到，这里会判断handler是否为空，也就是我们传入ListenAndServe的第二个参数，我们传入的为空，因此默认为`DefaultServeMux`，之后再调用`handler.ServeHTTP`，我们这里的`DefaultServeMux`就是一个`Handler`，因此根据里面的规则使用就可以了
## 参考
[astaxie/build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang)
这里强势安利一下这本书，我个人由于学习了一些golang语法的皮毛，因此直接从web那一章开始看的，3.4小节总的就是上面我描述的过程，在书本之后还有更多详细的介绍