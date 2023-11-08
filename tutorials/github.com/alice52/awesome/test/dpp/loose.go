package dpp

// 此时代码与某个logger强耦合
// 应该将依赖项解耦出来, 并且将依赖注入到我们的 App 实例中
// 依赖注入就是指在创建组件的时候接收它的依赖项
//var log = logrus.New()
//type App struct{}
//
//func (a *App) Start() {
//	log.Info("app start ...")
//}

type App struct {
	Logger
}

func (a *App) Start() {
	a.Logger.Info("app start ...")
	// ...
}

// NewApp 构造函数，将依赖项注入
func NewApp(lg Logger) *App {
	return &App{
		Logger: lg, // 使用传入的依赖项完成初始化
	}
}

// Logger 将日志库抽象为接口类型
type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}
