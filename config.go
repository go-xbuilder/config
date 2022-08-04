package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Resources map[string]ResourceValue

func (r Resources) In(name string) bool {
	if _, ok := r[name]; ok {
		return true
	}
	return false
}

func (r Resources) Value(name string) ResourceValue {
	if v, ok := r[name]; ok {
		return v
	}
	return ResourceValue{}
}

type Where map[string]string

func (w Where) In(key string) bool {
	if _, ok := w[key]; ok {
		return true
	}
	return false
}

type OrderFields []string

func (of OrderFields) In(key string) bool {
	for _, k := range of {
		if key == k {
			return true
		}
	}
	return false
}

type ResourceValue struct {
	Where       Where                        `yaml:"where"`
	Convert     map[string]map[string]string `json:"convert"`
	OrderFields OrderFields                  `yaml:"orderFields"`
}

type Config struct {
	valid bool
	Error error
	Alias struct {
		EmptyValue    string `yaml:"emptyValue"`
		NotEmptyValue string `yaml:"notEmptyValue"`
		NullValue     string `yaml:"nullValue"`
		NotNullValue  string `yaml:"notNullValue"`
	} `yaml:"alias"`
	Resources Resources `json:"resources"`
}

func NewConfig(filename string) Config {
	var c Config
	bytes, err := os.ReadFile(filename)
	if err == nil {
		if err = yaml.Unmarshal(bytes, &c); err != nil {
			c.Error = err
			log.Printf("table config file unmarshal error: %s", err)
		} else {
			c.valid = true
		}
	} else {
		log.Printf("read table config file error: %s", err.Error())
		c.Error = err
	}

	return c
}

func (c Config) Resource(name string) (rv ResourceValue) {
	return c.Resources.Value(name)
}
