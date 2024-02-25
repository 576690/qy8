package config

import "strconv"

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Config   string `yaml:"config"` //高级配置
	DB       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"` //日志等级
}

func (m Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config
}
