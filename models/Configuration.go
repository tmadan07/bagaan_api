package models

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
)

type CryptoEngineSetting struct {
	Host string
	Port int
	Size int
}

type ConnectionString struct {
	Host              string
	Port              int
	User              string
	Password          string
	Dbname            string
	Max_connections   int
	Max_open_conns    int
	Max_idle_conns    int
	Conn_max_lifetime int
}

type Configuration struct {
	CryptoEngine         CryptoEngineSetting // Crypto engine host, port and pool size
	CBS                  ConnectionString    // database connection setting to CMS + CBS database
	TokenSymmetricKey    string
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
	ServerAddress        string
}

//define config as global variable

var once sync.Once

var (
	configuration Configuration
)

func GetConfiguration() Configuration {

	once.Do(func() { // <-- atomic, does not allow repeating

		log.Info("Loading config")

		configuration = Configuration{} // <-- thread safe
		err := gonfig.GetConf("config.yaml", &configuration)
		if err != nil {
			log.Panic(err)
			return
		}

	})

	return configuration
}
