package config

type Config struct {
	Port             int `mapstructure:"port"`
	TraceIdCacheTime int `mapstructure:"traceid_cache_time"`
}
