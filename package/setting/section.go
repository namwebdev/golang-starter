package setting

type Config struct {
	MySql MySqlSetting `mapstructure:"mysql"`
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
