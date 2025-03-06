package config

type AppConfig struct {
	DataSource string `json:"dataSource"`
	DataDir    string `json:"dataDir"`
	ServerURL  string `json:"serverUrl"`
	Port       int    `json:"port"`
}
