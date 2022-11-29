package conf

type Config struct {
	App   appConfig   `json:"app"`
	Log   logConfig   `json:"log"`
	MySQL mysqlConfig `json:"mysql"`
}

type (
	appConfig struct {
		Port         string `json:"port"`
		ReadTimeout  uint   `json:"read_timeout"`
		WriteTimeout uint   `json:"write_timeout"`
	}

	logConfig struct {
		Level      string `json:"level"`
		FileLevel  string `json:"file_level"`
		OutputFile string `json:"output_file"`
	}

	mysqlConfig struct {
		User    string `yaml:"user"`
		Pass    string `yaml:"pass"`
		Host    string `yaml:"host" validate:"required"`
		Port    string `yaml:"port"`
		DBName  string `yaml:"db_name" validate:"required"`
		MaxOpen int    `yaml:"max_open_conn"`
		MaxIdle int    `yaml:"max_idle_conn"`
	}
)
