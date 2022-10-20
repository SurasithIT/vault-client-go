package config

import (
	"os"
	"poc/vault-client/config/vault"
)

type Config struct {
	Vault vault.Config
}

func New() (*Config, error) {
	var config Config

	config.Vault.Address = os.Getenv("VAULT_ADDRESS")
	config.Vault.Token = os.Getenv("VAULT_TOKEN")
	config.Vault.EngineName = os.Getenv("VAULT_ENGINE_NAME")
	config.Vault.SecretPath = os.Getenv("VAULT_SECRET_PATH")

	return &config, nil
}
