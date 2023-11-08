package dpp

// Config 配置项结构体
type Config struct {
	// ...
}

// LoadConfFromFile 从配置文件中加载配置
func LoadConfFromFile(filename string) *Config {
	return &Config{}
}

// Server server 程序
type Server struct {
	Config *Config
}

// NewServer Server 构造函数
func NewServerV1(conf *Config) *Server {
	return &Server{
		// 隐式创建依赖项
		Config: conf,
	}
}

// NewServer Server 构造函数
func NewServerV0() *Server {
	return &Server{
		// 隐式创建依赖项
		Config: LoadConfFromFile("./config.toml"),
	}
}
