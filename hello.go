package main

import (
	"fmt"
	"syscall/js"
	"net/http"
)

var headers js.Value

func init() {
	headers := js.Global().Get("Object").New()
	headers.Set("Content-Type", "text/plain")
}

func main() {
	require := js.Global().Get("require")
	http := require.Invoke("http")
	app := http.Get("Server").New()
	app.Call("on", "request", js.NewCallback(request))
	app.Call("listen", 3000, js.NewCallback(listen))
	wait()
}

func request(args []js.Value) {
	res := args[1]
	
	http.Get("https://api.github.com", resp *http.Response, err error)
	resp.Status()
	
	res.Call("writeHead", 200, headers)
	res.Call("write", "Hello World")
	res.Call("end", "\n")
}

func listen(args []js.Value) {
	fmt.Println("listening on port 3000")
}

func wait() {
	done := make(chan bool)
	js.Global().Get("process").Call("on", "SIGTERM", js.NewCallback(func(args []js.Value) {
		done <- true
	}))
	<-done
}
