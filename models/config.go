package models

type Config struct {
	Telegram `json:"telegram"`
	Logger   `json:"logeger"`
	Postgres `json:"postgres"`
}

type Telegram struct {
	Token     string `json:"token"`
	ModeDebug bool   `json:"mode_debug"`
}

type Logger struct {
	Level string
}

type Postgres struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
	SSLMode  string `json:"ssl_mode"`
}
