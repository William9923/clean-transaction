package conf

type Config struct {
	Log   logConfig   `json:"log"`
	MySQL mysqlConfig `json:"mysql"`
}

var cfg Config = Config{}

func GetConfig() Config {
	return cfg
}

type (
	logConfig struct {
		Level      string `json:"level"`
		FileLevel  string `json:"file_level"`
		OutputFile string `json:"output_file"`
	}

	mysqlConfig struct {
		User    string `json:"user"`
		Pass    string `json:"pass"`
		Host    string `json:"host"`
		Port    string `json:"port"`
		DBName  string `json:"db_name"`
		MaxOpen int    `json:"max_open_conn"`
		MaxIdle int    `json:"max_idle_conn"`
	}
)
