package hook

import "github.com/practice/Design-Patterns-practice/other/create-resource-pattern/resource"

// ResourceHook 创建资源时的钩子方法
type ResourceHook interface {
	BeforeCreate(resource *resource.Resource) error
	AfterCreate(resource *resource.Resource) error
}
