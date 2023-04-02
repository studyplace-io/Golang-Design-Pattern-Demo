package proxy_pattern

// Server 服务 接口对象
type Server interface {
	HandleRequest(string, string) (int, string)
}

// Application 真正的应用对象
type Application struct {
	name  string
	other string
}

type ApplicationOption func(*Application)

func WithApplicationName(name string) ApplicationOption {
	return func(application *Application) {
		application.name = name
	}
}
func WithApplicationOther(other string) ApplicationOption {
	return func(application *Application) {
		application.other = other
	}
}

// NewApplication 加入选项模式
func NewApplication(opts ...ApplicationOption) *Application {
	in := &Application{
		name:  defaultName,
		other: defaultOther,
	}

	for _, opt := range opts {
		opt(in)
	}
	return in
}

func (a *Application) GetApplicationName() string {
	return a.name
}

func (a *Application) GetApplicationOther() string {
	return a.other
}

// HandleRequest 真正的处理http请求
func (a *Application) HandleRequest(url, method string) (int, string) {

	if url == "/app/status" && method == "GET" {
		return 200, "Ok"
	}

	if url == "/create/user" && method == "POST" {
		return 201, "User Created"
	}
	return 404, "Not Ok"
}