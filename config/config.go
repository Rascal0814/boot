package config

import (
	"github.com/Rascal0814/boot"
	"github.com/Rascal0814/boot/orm"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"path"
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
func LoadConfig(confPath string) (*Config, error) {
	var conf = new(Config)
	var err error = nil
	c := config.New(
		config.WithSource(
			file.NewSource(getConfigName(confPath)),
		),
	)
	defer func() { _ = c.Close() }()

	if err := c.Load(); err != nil {
		panic(err)
	}

	if err := c.Scan(conf); err != nil {
		panic(err)
	}

	connectDB, err := orm.ConnectDB(conf.Data.Database.Driver, conf.Data.Database.Source)
	if err != nil {
		return nil, err
	}

	db, err = gorm.Open(connectDB, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}
	return conf, nil
}

func DefaultDB(c *Config) *gorm.DB {
	return db
}
