package create_resource_pattern

import (
	"fmt"
	"github.com/practice/Design-Patterns-practice/other/create-resource-pattern/hook/auth_hook"
	"github.com/practice/Design-Patterns-practice/other/create-resource-pattern/hook/log_hook"
	"github.com/practice/Design-Patterns-practice/other/create-resource-pattern/hook/validator_hook"
	"log"
	"regexp"
	"testing"
)

func TestResourceManager(t *testing.T) {
	rm := NewResourceManager()

	// 创建一个身份验证钩子对象
	authHook := &auth_hook.AuthHook{
		Username: "admin",
		Password: "password",
	}

	// 注册身份验证钩子
	rm.RegisterHook(authHook)

	// 创建一个日志记录钩子对象
	logger := log.New(log.Writer(), "", log.LstdFlags)
	loggingHook := &log_hook.LoggingHook{
		Logger: logger,
	}

	// 注册日志记录钩子
	rm.RegisterHook(loggingHook)

	// 创建一个电子邮件验证钩子对象
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	emailValidator := &validator_hook.EmailValidatorHook{
		EmailRegex: emailRegex,
	}

	rm.RegisterHook(emailValidator)

	err := rm.CreateResource("1", "Resource 1", "john_doe11@example.com")
	if err != nil {
		fmt.Println("Error creating resource:", err)
		return
	}

	err = rm.CreateResource("2", "Resource 2", "john_doe@example.com")
	if err != nil {
		fmt.Println("Error创建资源:", err)
		return
	}

	resources := rm.ListResources()
	fmt.Println("Resources:")
	for _, resource := range resources {
		fmt.Printf("ID: %s, Name: %s\n", resource.ID, resource.Name)
	}

	resource, err := rm.GetResourceByID("1")
	if err != nil {
		fmt.Println("Error getting resource:", err)
		return
	}

	err = rm.DeleteResourceByID("1")
	if err != nil {
		fmt.Println("Error deleting resource:", err)
		return
	}

	fmt.Println("Resource 1 Name:", resource.Name)
}
