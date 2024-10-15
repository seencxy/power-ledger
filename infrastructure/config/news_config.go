package config

import (
	"PowerLedgerGo/domain/entity"
	"PowerLedgerGo/infrastructure/contract/artifacts"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/v3/client"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/driver/mysql"
	"log"
	"time"

	"github.com/go-god/setting"
	"gorm.io/gorm"
)

const configPath = "./infrastructure/config/config.yaml"

// DBConfig database config
type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Dbname   string
}

// AppConfig app section
// Underline variables need to be annotated with a mapstructure tag
type AppConfig struct {
	Port         int
	PProfPort    int           `mapstructure:"pprof_port"`
	AppName      string        `mapstructure:"app_name"`
	AppEnv       string        `mapstructure:"app_env"`
	AppDebug     bool          `mapstructure:"app_debug"`
	GracefulWait time.Duration `mapstructure:"graceful_wait"`
}

type ContractConfig struct {
	AdvancedVirtualPowerPlantDAOAddress string `mapstructure:"AdvancedVirtualPowerPlantDAOAddress"`
	HomomorphicEncryptionAddress        string `mapstructure:"HomomorphicEncryptionAddress"`
	DeployByPrv                         string `mapstructure:"DeployByPrv"`
}

// configImpl config
type ConfigImpl struct {
	DB       DBConfig
	App      AppConfig
	Contract ContractConfig
}

// NewConfig load config
func NewConfig() *ConfigImpl {
	s := &ConfigImpl{}
	s.load()

	return s
}

// read and parse the configuration file
func (s *ConfigImpl) load() {
	conf := setting.New(setting.WithConfigFile(configPath))
	if err := conf.Load(); err != nil {
		log.Fatalf("read config file err:%s\n", err.Error())
	}

	if err := conf.ReadSection("app", &s.App); err != nil {
		log.Fatalf("read app section err:%s", err.Error())
	}
	if s.App.PProfPort == 0 {
		s.App.PProfPort = s.App.Port + 1000
	}

	if err := conf.ReadSection("db", &s.DB); err != nil {
		log.Fatalf("read db section err:%s", err.Error())
	}

	if err := conf.ReadSection("contract", &s.Contract); err != nil {
		log.Fatalf("read contract section err:%s", err.Error())
	}
}

// InitDB init gorm db
func (s *ConfigImpl) InitDB() *gorm.DB {
	dbConf := s.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbConf.User,
		dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("db open error: ", err)
	}

	// 数据库迁移
	_ = db.AutoMigrate(&entity.UserInfo{})
	_ = db.AutoMigrate(&entity.Trade{})

	return db
}

type ContractInstance struct {
	Client                              *client.Client
	AdvancedVirtualPowerPlantDAOSession artifacts.AdvancedVirtualPowerPlantDAOSession
	HomomorphicEncryptionSession        artifacts.HomomorphicEncryptionSession
}

func (s *ConfigImpl) InitContractInstance() *ContractInstance {
	var instance ContractInstance
	privateKey, _ := hex.DecodeString(s.Contract.DeployByPrv)
	config := &client.Config{
		IsSMCrypto:  false,
		GroupID:     "group0",
		PrivateKey:  privateKey,
		Host:        "127.0.0.1",
		Port:        20200,
		TLSCaFile:   "./infrastructure/config/ca.crt",
		TLSKeyFile:  "./infrastructure/config/sdk.key",
		TLSCertFile: "./infrastructure/config/sdk.crt"}

	contractInstance, err := client.DialContext(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	// 创建AdvancedVirtualPowerPlantDAOSession实例
	AdvancedVirtualPowerPlantDAOInstance, err := artifacts.NewAdvancedVirtualPowerPlantDAO(common.HexToAddress(s.Contract.AdvancedVirtualPowerPlantDAOAddress), contractInstance)
	if err != nil {
		log.Fatal(err)
	}

	instance.AdvancedVirtualPowerPlantDAOSession = artifacts.AdvancedVirtualPowerPlantDAOSession{
		Contract:     AdvancedVirtualPowerPlantDAOInstance,
		CallOpts:     *contractInstance.GetCallOpts(),
		TransactOpts: *contractInstance.GetTransactOpts(),
	}

	// 创建HomomorphicEncryption实例
	HomomorphicEncryptionInstance, err := artifacts.NewHomomorphicEncryption(common.HexToAddress(s.Contract.HomomorphicEncryptionAddress), contractInstance)
	if err != nil {
		log.Fatal(err)
	}

	instance.HomomorphicEncryptionSession = artifacts.HomomorphicEncryptionSession{
		Contract:     HomomorphicEncryptionInstance,
		CallOpts:     *contractInstance.GetCallOpts(),
		TransactOpts: *contractInstance.GetTransactOpts(),
	}

	instance.Client = contractInstance
	return &instance
}

// AppConfig returns app config
func (s *ConfigImpl) AppConfig() *AppConfig {
	return &s.App
}
