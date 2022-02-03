package getKey

import (
	"fmt"
	"log"

	"github.com/hashicorp/vault/api"
)

func GetVaultSecret(a, t, n, p, k string) (string, error) {
	config := &api.Config{
		Address: a,
	}
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	client.SetToken(t)
	secret, err := client.Logical().Read(p)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	m, ok := secret.Data["data"].(map[string]string)
	if !ok {
		fmt.Printf("%T %#v\n", secret.Data["data"], secret.Data["data"])
	}
	return m[k], nil
}
