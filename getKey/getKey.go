package getKey

import (
	"fmt"
	"gitlab-reporter/checkStatus"
	"gitlab-reporter/reports"
	"log"
	"os"

	"github.com/hashicorp/vault/api"
)

func ParseToken(m map[string]string) {
	if v, found := m["Address"]; found {
		err := checkStatus.CheckVaultStatus(v, m["Token"])
		if err != nil {
			log.Fatalf("%s\n", err)
			os.Exit(1)
		}
		token, err := getVaultSecret(v, m["Token"], m["Namespace"], m["Path"], m["Key"])
		if err != nil {
			log.Fatalf("%s\n", err)
			os.Exit(1)
		}
		m["Token"] = token
		err = checkStatus.CheckGitlabStatus(m["Url"], m["Token"])
		if err != nil {
			log.Fatalf("%s\n", err)
			os.Exit(1)
		}
		reports.ExecuteReport(m)
	} else {
		err := checkStatus.CheckGitlabStatus(m["Url"], m["Token"])
		if err != nil {
			log.Fatalf("%s\n", err)
			os.Exit(1)
		}
		reports.ExecuteReport(m)
	}
}

func getVaultSecret(a, t, n, p, k string) (string, error) {
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
