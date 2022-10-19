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
	config.Vault.EnvName = os.Getenv("VAULT_ENV_NAME")
	config.Vault.EngineName = os.Getenv("VAULT_ENGINE_NAME")
	config.Vault.MountPath = os.Getenv("VAULT_MOUNT_PATH")

	return &config, nil
}
