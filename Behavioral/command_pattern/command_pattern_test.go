package command_pattern

import "testing"

func TestCommandPattern(t *testing.T) {

	loginCmd := NewLoginCommand(new(UserService))
	scoreCmd := NewScoreCommand(new(ScoreService))
	userInvoker := new(Invoker)
	userInvoker.AddCommand(loginCmd, scoreCmd)

	userInvoker.Do("test", "11111")

}

