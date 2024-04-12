package config

import "time"

type AppConfig struct {
	Server         Server
	DB             DB
	Cache          Cache
	RabbitMQ       RabbitMQ
	AdapterService AdapterService
	Firebase       Firebase
}

type Server struct {
	KeyID               string
	Name                string
	AppVersion          string
	RestAPIPort         string
	GrpcPort            string
	BaseURL             string
	DebugMode           bool
	ReadTimeout         time.Duration
	WriteTimeout        time.Duration
	SSL                 bool
	CtxDefaultTimeout   int
	CSRF                bool
	Debug               bool
	MaxCountRequest     int           // max count of connections
	ExpirationLimitTime time.Duration //  expiration time of the limit
}

type Cache struct {
	Redis Redis
}

type DB struct {
	Mongodb Mongodb
}

type AdapterService struct {
	AuthService AuthService
	UserService UserService
}

type Firebase struct {
	ServiceAccountKeyFilePath string
}

type Redis struct {
	Address  string
	Port     int
	Password string
	DB       int
}

type Mongodb struct {
	DbName          string
	Username        string
	Password        string
	Connection      string
	ConnectTimeout  time.Duration
	MaxConnIdleTime int
	MinPoolSize     uint64
	MaxPoolSize     uint64
}

type RabbitMQ struct {
	Connection  string
	ServiceName string

	OrderTransactionExchange string
	OrderCommitRoutingKey    string
}

type NotifyEvent struct {
	Connection        string
	Exchange          string
	PublishRoutingKey string
	Queue             string
}

type AuthService struct {
	BaseURL     string
	InternalKey string
}

type UserService struct {
	UserURL     string
	InternalKey string
}
