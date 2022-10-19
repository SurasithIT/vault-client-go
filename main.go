package main

import (
	"context"
	"fmt"
	"poc/vault-client/config"
	"poc/vault-client/vault"

	"github.com/joho/godotenv"
)

type Config struct {
	Test string
}

func main() {
	err := godotenv.Load("./etc/local.env")
	if err != nil {
		fmt.Printf("please consider environment variables: %s", err)
		panic(err)
	}

	config, err := config.New()
	if err != nil {
		panic(err)
	}

	vaultAddress := config.Vault.Address
	vaultToken := config.Vault.Token
	mountPath := config.Vault.MountPath
	secretPath := fmt.Sprintf("%s/%s", config.Vault.EngineName, config.Vault.EnvName)

	vault := vault.NewVaultClient(vaultAddress, vaultToken)
	ctx := context.Background()

	vault.SaveSecret(ctx, mountPath, secretPath, "test-secret", "secret-value")

	secret, err := vault.ReadSecret(ctx, mountPath, secretPath, "test-secret")
	if err != nil {
		panic(err)
	}
	fmt.Println("secret value is", secret)

	vault.DeleteSecret(ctx, mountPath, secretPath, "test-secret")
}
