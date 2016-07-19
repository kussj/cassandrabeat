package beater

type CassandraConfig struct {
	Period *int64
	Table  []string
}

type ConfigSettings struct {
	Input CassandraConfig
}
