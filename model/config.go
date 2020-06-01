package model

// Config is the base config structure for the server
type Config struct {
	TCPConfig      TCPConfig
	HTTPConfig     HTTPConfig
	DatabaseConfig DatabaseConfig
	ClusterConfig  ClusterConfig `json:"config"`
}

// TCPConfig contains the config related to the TCP chat server
type TCPConfig struct {
	Port          int
	ListenAddress string
}

// HTTPConfig contains the config related to the TCP chat server
type HTTPConfig struct {
	Port          int
	ListenAddress string
}

// DatabaseConfig contains the config for the databases(s) which the application requires
type DatabaseConfig struct {
	RethinkDBConfig RethinkDBConfig
	MongoDBConfig   MongoDBConfig
}

// RethinkDBConfig contains the config for RethinkDB
type RethinkDBConfig struct {
	Address  string
	Database string
}

// MongoDBConfig contains the config for MongoDB
type MongoDBConfig struct {
	URI string
}

// ClusterConfig contains the config for the cluster, also available to the clients via HTTP
type ClusterConfig struct {
	DirectoryServerConfig DirectoryServerConfig `json:"directory"`
	MediaServerConfig     MediaServerConfig     `json:"media"`
}

// DirectoryServerConfig has the config for the directory server
type DirectoryServerConfig struct {
	Endpoint string `json:"endpoint"`
}

// MediaServerConfig has the config for the media server
type MediaServerConfig struct {
	Endpoint string `json:"endpoint"`
}

// SetDefaults sets the default values for the config structure
func (c *Config) SetDefaults() {
	// Set the TCP Config
	// Default TCP port is 1337
	// By default, the TCP server listens on 127.0.0.1
	c.TCPConfig.Port = 1337
	c.TCPConfig.ListenAddress = "127.0.0.1"

	// Set the HTTP Config
	// Default HTTP port is 3000
	// By default, the HTTP server listens on 127.0.0.1
	c.HTTPConfig.Port = 3005
	c.HTTPConfig.ListenAddress = "127.0.0.1"

	// Set the Database Config

	// Set the RethinkDB config
	c.DatabaseConfig.RethinkDBConfig.Address = "127.0.0.1:28015"
	c.DatabaseConfig.RethinkDBConfig.Database = "paperplane"

	// Set the MongoDB config
	c.DatabaseConfig.MongoDBConfig.URI = "mongodb://127.0.0.1/paperplane"

	// Set the Cluster Config

	// Set the Directory Server's Config
	// By default, the same go binary has the directory server
	// Default endpoint is /directory
	c.ClusterConfig.DirectoryServerConfig.Endpoint = "/directory"

	// Set the Media Server's Config
	// By default, the same go binary has the media server
	// Default endpoint is /media
	c.ClusterConfig.MediaServerConfig.Endpoint = "/media"
}
