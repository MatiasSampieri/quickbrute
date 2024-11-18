package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Flags struct {
	BatchSize  int
	ConfigFile string
	Port       int
	OutFile    string
}

type Config struct {
	Request  Request          `yaml:"request"`
	Params   map[string]Param `yaml:"params,omitempty"`
	Criteria Criteria         `yaml:"criteria,omitempty"`
	Helpers  []string         `yaml:"helpers,omitempty"`
	Logger   Logger
}

func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func handleFlags() (*Flags, bool) {
	var flags Flags

	flag.IntVar(&flags.BatchSize, "b", 500, "Set amount of parallel requests")
	flag.IntVar(&flags.Port, "p", 7575, "Set port for inter-node communication")
	flag.StringVar(&flags.OutFile, "o", "out.log", "Set the output file for criteria results")
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Missing config file! Running in helper mode")
		return &flags, true
	}

	flags.ConfigFile = args[0]

	return &flags, false
}

func loadParams(config *Config) []string {
	var paramNames []string

	if config.Params == nil {
		return paramNames
	}

	paramNames = make([]string, len(config.Params))
	index := 0
	for name := range config.Params {
		paramNames[index] = name
		index++
	}

	return paramNames
}
