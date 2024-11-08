package config

import (
	"github.com/spf13/viper"
)

/*github.com/spf13/viper: Это популярная библиотека в Go для работы с конфигурационными файлами
и переменными окружения. Viper поддерживает различные форматы файлов (JSON, YAML, TOML и др.)
и упрощает процесс загрузки и использования конфигурационных данных в приложении.*/
type Config struct {
    Database struct {
        Host     string
        Port     int
        User     string
        Password string
        DBName   string
    }
}

func LoadConfig() (*Config, error) {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    viper.AutomaticEnv() //Позволяет Viper считывать переменные окружения и использовать их для переопределения значений из конфигурационного файла.

    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }

    var cfg Config
    if err := viper.Unmarshal(&cfg); err != nil {
        return nil, err
    }

    return &cfg, nil
}
