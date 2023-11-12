package validator_hook

import (
	"fmt"
	"github.com/practice/Design-Patterns-practice/other/create-resource-pattern/resource"
	"regexp"
)

// EmailValidatorHook 是电子邮件验证的钩子实现
type EmailValidatorHook struct {
	EmailRegex *regexp.Regexp
}

// BeforeCreate 在验证用户之前执行的钩子函数
func (h *EmailValidatorHook) BeforeCreate(resource *resource.Resource) error {
	if !h.EmailRegex.MatchString(resource.Email) {
		return fmt.Errorf("invalid email address")
	}
	return nil
}

// AfterCreate 在验证用户之后执行的钩子函数
func (h *EmailValidatorHook) AfterCreate(resource *resource.Resource) error {
	fmt.Printf("User with ID %s validated successfully\n", resource.ID)
	return nil
}
