package config

import (
	"encoding/json"
	"io/ioutil"
)

type Configuration struct {
	Sources           []ConfigSource     `json:"sources"`
	AWSFlowLogConfigs []AWSFlowLogConfig `json:"awsFlowLogConfigs"`
}

type ConfigSource struct {
	IP   string            `json:"ip"`
	SNMP *ConfigSourceSNMP `json:"snmp"`
}

type ConfigSourceSNMP struct {
	Port           int    `json:"port"`
	User           string `json:"user"`
	AuthPassphrase string `json:"auth"`
	PrivPassphrase string `json:"priv"`
}

type SNMPEntry struct {
	Address        string `json:"address"`
	User           string `json:"user"`
	AuthPassphrase string `json:"authPassphrase"`
	PrivPassphrase string `json:"privPassphrase"`
}

type AWSFlowLogConfig struct {
	Region          string `json:"region"`
	AccessKeyID     string `json:"accessKeyID"`
	SecretAccessKey string `json:"secretAccessKey"`
	LogGroupName    string `json:"logGroupName"`
	LogStreamName   string `json:"logStreamName"`
}

func Load(path string) (Configuration, error) {
	conf := Configuration{}
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return conf, err
	}
	return conf, json.Unmarshal(contents, &conf)
}
