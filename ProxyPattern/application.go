package ProxyPattern

// 真正的应用对象
type Application struct {
}

// 真正的处理http请求
func (a *Application) HandleRequest(url, method string) (int, string) {

	if url == "/app/status" && method == "GET" {
		return 200, "Ok"
	}

	if url == "/create/user" && method == "POST" {
		return 201, "User Created"
	}
	return 404, "Not Ok"
}
