package setting

type Config struct {
	MySql  MySqlSetting `mapstructure:"mysql"`
	Logger LogSetting   `mapstructure:"logger"`
}

type MySqlSetting struct {
	Host                  string `mapstructure:"host"`
	Port                  int    `mapstructure:"port"`
	User                  string `mapstructure:"user"`
	Password              string `mapstructure:"password"`
	DbName                string `mapstructure:"dbName"`
	MaxIdleConnections    int    `mapstructure:"maxIdleConnections"`
	MaxOpenConnections    int    `mapstructure:"maxOpenConnections"`
	ConnectionMaxLiftTime int    `mapstructure:"connectionMaxLiftTime"`
}

type LogSetting struct {
	LogLevel    string `mapstructure:"log_level"`
	FileLogName string `mapstructure:"file_log_name"`
	MaxSize     int    `mapstructure:"max_size"`
	MaxBackups  int    `mapstructure:"max_backups"`
	MaxAge      int    `mapstructure:"max_age"`
	Compress    bool   `mapstructure:"compress"`
}
