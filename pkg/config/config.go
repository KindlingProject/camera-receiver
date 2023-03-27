package config

type Config struct {
	Port             int `yaml:"port,omitempty"`
	TraceIdCacheTime int `yaml:"traceid_cache_time,omitempty"`
}
