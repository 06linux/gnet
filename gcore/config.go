package gcore

// 服务器配置
type ServerConfig struct {

	// 服务器名称
	name string

	// 网络连接方式 （tcp 或者 websocket)
	network string

	// 主机地址（格式： 127.0.0.1:8000)
	addr string

	// 限制最大链接数量
	maxConn int

	// woker最大数量（每一个 worker 对应一个 Goroutine)
	maxWorker int

	// 每一个 work 可以处理 task 的上限
	maxTask int
}

// 转发配置
type TransferConfig struct {
	name    string
	channel string
	addr    string
}

type Config struct {

	// 服务器运行配置
	server *ServerConfig

	// 转发配置
	transfer *TransferConfig
}

func NewConfig() *Config {
	return &Config{
		server: &ServerConfig{
			name:      "net-server",
			network:   "tcp",
			addr:      "127.0.0.1:8000",
			maxConn:   10,
			maxWorker: 3,
			maxTask:   10,
		},
		transfer: nil,
	}
}

// 加载服务器配置
func (config *Config) LoadServer(fileName string) {

}

// 加载转发配置
func (config *Config) LoadTransfer(fileName string) {

}
