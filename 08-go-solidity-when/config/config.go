package config

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
)

type Config struct {
	RPCURL               string `json:"rpc_url"`
	WSURL                string `json:"ws_url"`
	ContractAddress      string `json:"contract_address"`
	StartBlock           uint64 `json:"start_block"`
	Confirmations        uint64 `json:"confirmations"`
	ReplayIntervalSecond int    `json:"replay_interval_seconds"`
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
	if v := os.Getenv("WHEN_START_BLOCK"); v != "" {
		parsed, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return Config{}, err
		}
		cfg.StartBlock = parsed
	}
	if v := os.Getenv("WHEN_CONFIRMATIONS"); v != "" {
		parsed, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return Config{}, err
		}
		cfg.Confirmations = parsed
	}
	if v := os.Getenv("WHEN_REPLAY_INTERVAL_SECONDS"); v != "" {
		parsed, err := strconv.Atoi(v)
		if err != nil {
			return Config{}, err
		}
		cfg.ReplayIntervalSecond = parsed
	}

	if cfg.RPCURL == "" {
		cfg.RPCURL = "http://localhost:8545"
	}
	if cfg.WSURL == "" {
		cfg.WSURL = "ws://127.0.0.1:8545"
	}
	if cfg.Confirmations == 0 {
		cfg.Confirmations = 1
	}
	if cfg.ReplayIntervalSecond == 0 {
		cfg.ReplayIntervalSecond = 30
	}
	return cfg, nil
}
