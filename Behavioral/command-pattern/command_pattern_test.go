package command_pattern

import "testing"

func TestCommandPattern(t *testing.T) {

	us := NewUserService()
	ss := NewScoreService()
	loginCmd := NewLoginCommand(us)
	scoreCmd := NewScoreCommand(ss)

	userInvoker := NewInvoker()
	userInvoker.AddCommand(loginCmd, scoreCmd) // 加入命令后，可以存储著并按需执行。

	// 执行
	userInvoker.Execute("test", "1111")

}

