package config

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	RPCURL          string `json:"rpc_url"`
	WSURL           string `json:"ws_url"`
	ContractAddress string `json:"contract_address"`
}

func Load() (Config, error) {
	path := os.Getenv("WHEN_CONFIG_PATH")
	if path == "" {
		path = "config.json"
	}

	cfg := Config{}
	data, err := os.ReadFile(path)
	if err == nil {
		if err := json.Unmarshal(data, &cfg); err != nil {
			return Config{}, err
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return Config{}, err
	}

	if v := os.Getenv("WHEN_RPC_URL"); v != "" {
		cfg.RPCURL = v
	}
	if v := os.Getenv("WHEN_WS_URL"); v != "" {
		cfg.WSURL = v
	}
	if v := os.Getenv("WHEN_CONTRACT_ADDRESS"); v != "" {
		cfg.ContractAddress = v
	}
	if cfg.RPCURL == "" {
		cfg.RPCURL = "http://localhost:8545"
	}
	if cfg.WSURL == "" {
		cfg.WSURL = "ws://127.0.0.1:8545"
	}
	return cfg, nil
}
