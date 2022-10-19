package vault

import (
	"context"
	"fmt"

	vault "github.com/hashicorp/vault/api"
)

func NewVaultClient(vaultAddress string, vaultToken string) *Vault {
	vaultConfig := &vault.Config{
		Address: vaultAddress,
	}

	client, err := vault.NewClient(vaultConfig)
	if err != nil {
		panic(err)
	}

	client.SetToken(vaultToken)

	vault := &Vault{
		Client: client,
	}
	return vault
}

func (vault *Vault) SaveSecret(ctx context.Context, mountPath string, secretPath string, key string, value interface{}) (*vault.KVSecret, error) {
	secretData := map[string]interface{}{
		key: value,
	}

	output, err := vault.Client.KVv2(mountPath).Put(ctx, secretPath, secretData)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (vault *Vault) ReadSecret(ctx context.Context, mountPath string, secretPath string, key string) (interface{}, error) {
	secret, err := vault.Client.KVv2(mountPath).Get(ctx, secretPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read secret from vault : %v", err)
	}
	value := secret.Data[key]
	if value == nil {
		return nil, fmt.Errorf("not found value of " + key)
	}
	return value, nil
}

func (vault *Vault) DeleteSecret(ctx context.Context, mountPath string, secretPath string, key string) error {
	err := vault.Client.KVv2(mountPath).Delete(ctx, secretPath)
	if err != nil {
		return err
	}
	return nil
}
