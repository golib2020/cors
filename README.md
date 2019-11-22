# cors

[![Go Report](https://goreportcard.com/badge/github.com/golib2020/cors)](https://goreportcard.com/report/github.com/golib2020/cors)

`cors`它简单粗暴的解决了GoWeb跨域问题

使用方式：只要实现了`http.handler`接口即可轻松对接

# go默认web服务使用

```go
package main

import (
    "log"
    "net/http"
    
    "github.com/golib2020/cors"
)

func main() {
    mux := http.NewServeMux()
    
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("cors"))
    })
    
    //c := cors.New(mux) //默认参数为 *
    c := cors.New(
        mux,
        cors.WithOrigin("*"),
        cors.WithMethods("*"),
        cors.WithHeaders("*"),
    )
    
    log.Fatal(http.ListenAndServe(":8080", c))
}
```

# 第三方gorilla/mux

```go
package main

import (
    "log"
    "net/http"
    
    "github.com/golib2020/cors"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
    
    router.Methods("GET").Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("cors"))
    })
    
    //c := cors.New(mux) //默认参数为 *
    c := cors.New(
        router,
        cors.WithOrigin("*"),
        cors.WithMethods("*"),
        cors.WithHeaders("*"),
    )
    
    log.Fatal(http.ListenAndServe(":8080", c))
}
```
