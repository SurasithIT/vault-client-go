package vault

import vault "github.com/hashicorp/vault/api"

type Vault struct {
	Client *vault.Client
}
