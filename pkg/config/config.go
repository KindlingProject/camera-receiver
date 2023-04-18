package config

type ReceiverConfig struct {
	Port             int `mapstructure:"port"`
	TraceIdCacheTime int `mapstructure:"traceid_cache_time"`
}

type PrometheusConfig struct {
	Address          string  `mapstructure:"address"`
	QueryInterval    int     `mapstructure:"query_interval"`
	QueryP9xVal      float32 `mapstructure:"query_p9x_val"`
	QueryP9xTable    string  `mapstructure:"query_p9x_table"`
	QueryP9xDuration string  `mapstructure:"query_p9x_duration"`
}
