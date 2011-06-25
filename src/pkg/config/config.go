package config

import (
    "fmt"
    "io/ioutil"
    "json"
)

type StorageConfig map[string]string

type Config struct {
    Listen string
    Storage StorageConfig
}

func readFile(path string) []byte {
    data, err := ioutil.ReadFile(path)
    if err != nil {
        panic(fmt.Sprintf("Error reading config: %s", err.String()))
    }
    return data
}

func NewFromFile(path string) (config Config) {
    if err := json.Unmarshal(readFile(path), &config); err != nil {
        panic(fmt.Sprintf("Error parsing config: %s", err.String()))
    }
    return
}
