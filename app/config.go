package app

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"path"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator/v10"

	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

const (
	StateLocal State = "local"
	StateDEV   State = "dev"
	StateProd  State = "prod"
)

type State string

type Configs struct {
	vn            []*viper.Viper
	ConfigPath    string
	HTTPTransport *http.Transport
	Validator     *validator.Validate

	BuildVersion string
	BuildDate    string

	Stage State
	Mssql struct {
		EnableConnect bool          `mapstructure:"enable_connect"`
		DatabaseType  string        `mapstructure:"database_type"`
		Server        string        `mapstructure:"server"`
		Port          int           `mapstructure:"port"`
		Timeout       time.Duration `mapstructure:"timeout"`
		Username      string        `mapstructure:"username"`
		Password      string        `mapstructure:"password"`
		Database      struct {
			Master string `mapstructure:"master"`
		}
		ConnectionMasterDB string `mapstructure:"connection_master_db"`
		LogEnable          bool   `mapstructure:"log_enable"`
	} `mapstructure:"mssql"`

	Scheduler struct{
		EnableJob bool `mapstructure:"enable_job"`
	} `mapstructure:"scheduler"`
}

func (c *Configs) parseState(s string) State {
	switch s {
	case "local", "localhost", "l":
		return StateLocal
	case "dev", "develop", "development", "d":
		return StateDEV
	case "prod", "production", "p":
		return StateProd
	}
	return StateLocal
}

func (c *Configs) InitViperWithStage(s string) error {
	c.Stage = c.parseState(s)
	log.Infof("c.Stage: %s", c.Stage)

	name := fmt.Sprintf("config.%s", c.Stage)
	log.Infof("c.Stage name: %s", name)

	paths := []string{path.Join(c.ConfigPath, "general"), path.Join(c.ConfigPath, "secret")}
	log.Infof("config paths: %s", paths)
	if err := c.viperBinding(paths, name, true); err != nil {
		log.Errorln("binding error:", err)
	}

	return nil
}

var logWriter string
var logLevel log.Level

func (c *Configs) viperBinding(ps []string, cn string, isInit bool) error {
	for i, p := range ps {
		var vp *viper.Viper
		if isInit {
			vp = viper.New()
			vp.AddConfigPath(p)
			vp.SetConfigName(cn)
			vp.WatchConfig()
			vp.OnConfigChange(func(e fsnotify.Event) {
				log.Infoln("config file changed:", e.Name)
				c.viperBinding(ps, cn, false)
				log.Printf("config: %+v", c)
			})
		} else {
			vp = c.vn[i]
		}

		if err := vp.ReadInConfig(); err != nil {
			return err
		}

		c.vn = append(c.vn, vp)
		if err := vp.Unmarshal(&c); err != nil {
			log.Errorln("unmarshal config error:", err)
			return err
		}
	}

	if err := c.binding(); err != nil {
		return err
	}

	return nil
}

func (c *Configs) binding() error {

	c.Mssql.ConnectionMasterDB = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", //"sqlserver://%s:%s@%s:%d?database=%s",
		c.Mssql.Server,
		c.Mssql.Username,
		c.Mssql.Password,
		c.Mssql.Port,
		c.Mssql.Database.Master)
	log.Infof("ConnectionMasterDB: %s", c.Mssql.ConnectionMasterDB)

	tp := http.Transport{
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:          300,
		MaxIdleConnsPerHost:   300,
		IdleConnTimeout:       5 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ResponseHeaderTimeout: 5 * time.Second,
		ExpectContinueTimeout: 5 * time.Second,
	}
	c.HTTPTransport = &tp
	c.Validator = validator.New()

	return nil
}
