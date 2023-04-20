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

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) Login(name, pass string) error {
	if name == "test" && pass == "1111" {
		fmt.Printf("%s登入成功 \n", name)
		return nil
	}

	fmt.Printf("%s登入失败，用户或密码错误\n", name)
	return fmt.Errorf("error username or pass")
}

type ScoreService struct{}

func NewScoreService() *ScoreService {
	return &ScoreService{}
}

func (s *ScoreService) Score(name string) error {
	fmt.Printf("给%s赠送积分服务\n", name)
	return nil
}

// 上下解耦

// 命令模式

// ICommand 所有的命令都要实现此接口
type ICommand interface {
	Exec(args ...interface{}) error
}

// LoginCommand 登入命令
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
	// 调用嵌套中的对象方法
	err := lc.Login(args[0].(string), args[1].(string))

	return err
}

// ScoreCommand 积分命令
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
	// 调用嵌套中的对象方法
	err := sc.Score(args[0].(string))

	return err

}

// Invoker 调用者
type Invoker struct {
	// 重点：命令模式内部需要维护一个切片，用来存储不同命令
	cmds []ICommand
}

func NewInvoker() *Invoker {
	return &Invoker{cmds: make([]ICommand, 0)}
}

// Execute 执行所有命令，传入args 需要注意：传入的个数与顺序，避免报错
func (iv *Invoker) Execute(args ...interface{}) {
	for _, cmd := range iv.cmds {
		if err := cmd.Exec(args...); err != nil {
			// 如果当中有错误，直接失败
			break
		}
	}
}

func (iv *Invoker) AddCommand(cmds ...ICommand) {
	iv.cmds = append(iv.cmds, cmds...)
}
