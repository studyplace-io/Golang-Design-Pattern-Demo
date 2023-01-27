package command_pattern

import "fmt"

/*
	把请求封装为命令对象，把发出命令与执行命令的责任分开，可以传递给不同对象。
	1. 命令的发送者与执行者解耦。
	2. 可以进行命令对列的实现
	3. 方便记录执行过程，结合装饰器能进行
 */


// 业务代码

type UserService struct {
}


func (*UserService) Login(name, pass string) error {
	if name == "test" && pass == "1111" {
		fmt.Printf("%s登入成功 \n", name)
		return nil
	}

	fmt.Printf("%s登入失败，用户或密码错误\n", name)
	return fmt.Errorf("error username or pass")
}


type ScoreService struct {}

func (s *ScoreService) Score(name string) error {
	fmt.Printf("给%s赠送积分服务\n", name)
	return nil
}


// 上下解耦

// 命令模式

type ICommand interface {
	Exec(args ...interface{}) error
}


type LoginCommand struct {
	*UserService
}

func NewLoginCommand(userService *UserService) *LoginCommand {
	return &LoginCommand{UserService: userService}
}

func (lc *LoginCommand) Exec(args ...interface{}) error {
	if len(args) != 2 {
		panic("args error")
	}
	err := lc.Login(args[0].(string), args[1].(string))

	return err
}

type ScoreCommand struct {
	*ScoreService
}

func NewScoreCommand(scoreService *ScoreService) *ScoreCommand {
	return &ScoreCommand{ScoreService: scoreService}
}

func (sc *ScoreCommand) Exec(args ...interface{}) error {
	if len(args) < 1 {
		panic("args error")
	}
	err := sc.Score(args[0].(string))

	return err

}

type Invoker struct {
	cmds []ICommand
}

func (iv *Invoker) Do(args ...interface{}) {
	for _, cmd := range iv.cmds {
		if err := cmd.Exec(args...); err != nil {
			break // 如果当中有错误，直接失败
		}
	}
}

func (iv *Invoker) AddCommand(cmds ...ICommand) {
	iv.cmds = append(iv.cmds, cmds...)
}