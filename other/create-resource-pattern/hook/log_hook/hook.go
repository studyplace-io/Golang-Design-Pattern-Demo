package log_hook

import (
	"github.com/practice/Design-Patterns-practice/other/create-resource-pattern/resource"
	"log"
)

// LoggingHook 是日志记录的钩子实现
type LoggingHook struct {
	Logger *log.Logger
}

// BeforeCreate 在创建资源之前执行的钩子函数
func (h *LoggingHook) BeforeCreate(resource *resource.Resource) error {
	h.Logger.Printf("Creating resource with ID: %s\n", resource.ID)
	return nil
}

// AfterCreate 在创建资源之后执行的钩子函数
func (h *LoggingHook) AfterCreate(resource *resource.Resource) error {
	h.Logger.Printf("Resource with ID %s created successfully\n", resource.ID)
	return nil
}
