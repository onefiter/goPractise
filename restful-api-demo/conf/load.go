package conf

// Config 应用配置
// 通过封装为一个对象, 来与外部配置进行对接
type Config struct {
	App   *app   `toml:"app"`
	Log   *Log   `toml:"log"`
	MySQL *MySQL `toml:"mysql"`
}

type app struct {
	Name string `toml:"name" env:"APP_NAME"`
	Host string `toml:"host" env:"APP_HOST"`
	Port string `toml:"port" env:"APP_PORT"`
	Key  string `toml:"key" env:"APP_KEY"`
}

// MySQL todo
type MySQL struct {
	Host     string `toml:"host" env:"MYSQL_HOST"`
	Port     string `toml:"port" env:"MYSQL_PORT"`
	UserName string `toml:"username" env:"MYSQL_USERNAME"`
	Password string `toml:"password" env:"MYSQL_PASSWORD"`
	Database string `toml:"database" env:"MYSQL_DATABASE"`
	// 因为使用的MySQL连接池, 需要池做一些规划配置
	// 控制当前程序的MySQL打开的连接数
	MaxOpenConn int `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	// 控制MySQL复用, 比如5, 最多运行5个来复用
	MaxIdleConn int `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	// 一个连接的生命周期, 这个和MySQL Server配置有关系, 必须小于Server配置
	// 一个连接用12h 换一个conn, 保证一定的可用性
	MaxLifeTime int `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	// Idle 连接 最多允许存活多久
	MaxIdleTime int `toml:"max_idle_time" env:"MYSQL_MAX_idle_TIME"`
}

// Log todo
type Log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
}
