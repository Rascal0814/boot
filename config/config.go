package config

import (
	"github.com/Rascal0814/boot"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"path"
	"strings"
)

var db *gorm.DB

// getConfigName  获取文件名
func getConfigName(confPath string) string {
	c := boot.Loc
	if env := boot.GetEnv(); env != "" {
		c = boot.CType(env)
	}
	return path.Join(confPath, string(c)+".yaml")

}

// LoadConfig 加载配置文件 /config/xxx.yaml
func LoadConfig() (*Config, error) {
	var conf = new(Config)
	var err error = nil
	c := config.New(
		config.WithSource(
			file.NewSource(getConfigName("config")),
		),
	)
	defer func() { _ = c.Close() }()

	if err := c.Load(); err != nil {
		panic(err)
	}

	if err := c.Scan(conf); err != nil {
		panic(err)
	}

	switch strings.ToUpper(conf.Data.Database.Driver) {
	case "MYSQL":
		db, err = gorm.Open(mysql.Open(conf.Data.Database.Source), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	}

	if err != nil {
		return nil, err
	}
	return conf, nil
}

func DefaultDB() *gorm.DB {
	return db
}
