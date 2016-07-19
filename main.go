package main

import (
        cassandrabeat "github.com/bklise-goomzee/cassandrabeat/beater"

        "github.com/elastic/beats/libbeat/beat"
)

// You can overwrite these, e.g.: go build -ldflags "-X main.Version 1.0.0-beta3"
var Version = ""
var Name = "cassandrabeat"

func main() {
        beat.Run(Name, Version, cassandrabeat.New())
}
