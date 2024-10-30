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
}

type Request struct {
	Method  string            `yaml:"method"`
	URL     string            `yaml:"url"`
	Body    string            `yaml:"body,omitempty"`
	Headers map[string]string `yaml:"headers,omitempty"`
	Params  map[string]string `yaml:"params,omitempty"`
}

type Params struct {
	Type string `yaml:"type"`
	From int    `yaml:"from,omitempty"`
	To   int    `yaml:"to,omitempty"`
	File string `yaml:"file,omitempty"`
}

type Response struct {
	Status  int               `yaml:"status,omitempty"`
	Body    string            `yaml:"body,omitempty"`
	Headers map[string]string `yaml:"headers,omitempty"`
}

type Criteria struct {
	Type     string   `yaml:"type"`
	Response Response `yaml:"reponse"`
}

type Config struct {
	Request  Request           `yaml:"request"`
	Params   map[string]Params `yaml:"params,omitempty"`
	Criteria Criteria          `yaml:"criteria,omitempty"`
	Helpers  []string          `yaml:"helpers,omitempty"`
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
	for name, _ := range config.Params {
		paramNames[index] = name
		index++
	}

	return paramNames
}
