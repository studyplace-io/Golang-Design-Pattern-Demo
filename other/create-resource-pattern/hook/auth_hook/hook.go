package auth_hook

import (
	"fmt"
	"github.com/practice/Design-Patterns-practice/other/create-resource-pattern/resource"
)

// AuthHook 是身份验证的钩子实现
type AuthHook struct {
	Username string
	Password string
}

// BeforeCreate 在创建资源之前执行的钩子函数
func (h *AuthHook) BeforeCreate(resource *resource.Resource) error {
	// 模拟身份验证逻辑
	if h.Username != "admin" || h.Password != "password" {
		return fmt.Errorf("authentication failed")
	}
	return nil
}

// AfterCreate 在创建资源之后执行的钩子函数
func (h *AuthHook) AfterCreate(resource *resource.Resource) error {
	fmt.Printf("Resource with ID %s created successfully\n", resource.ID)
	return nil
}
