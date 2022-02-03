package helpMenu

import (
	"fmt"
	"os"
)

func DefaultHelp() {
	var defaultHelp string = `
	Usage: gitlab-reporter [command]

	An application that helps call the GitLab API and then
	exports a report of the results into various formats.

	The available commands for execution are listed below:

	Commands:
	--gitlab_url, -gitlab_url, gitlab_url, -gu			Passthrough your GitLab Url
	--vault_address, -vault_address, vault_address -va		Passthrough your HashiCorp Vault Address
	--namespace, -namespace, namespace -n 				Passthrough your HashiCorp Vault Namespace
	--path, -path, path, -p						Passthrough your HashiCorp Secret Path
	--secret_key, -secret_key, secret_key, -sk			Passthrough your HashiCorp Secret Key
	--api_token, -api_token, api_token, -at				Passthrough your API Token (Either GitLab API Token or Hashicorp Vault Token)
	--report, -report, report, -r					Passthrough the Report you wish to call
	--export_dir, -export_dir, export_dir, -ed			Allows you to select the export folder.
	--ouput_format, -output_format, output_format, -of 		Allows you to choose the output format of the report
	--help, -help, help, -h 					Shows the Help output
	--version, -version, version, -v				Shows the Version output
	`
	fmt.Printf("%s\n", defaultHelp)
	os.Exit(0)
}

func GitlabUrlHelp() {
	var gitlabUrlHelp string = `
	Usage: 		gitlab-reporter --gitlab_url https://gitlab.example.com
			gitlab-reporter -gitlab_url https://gitlab.example.com
			gitlab-reporter gitlab_url https://gitlab.example.com
			gitlab-reporter -gu https://gitlab.example.com

			One of the required flags. This is so the applications understands
			which gitlab url you will be calling, this allows functionality for
			self-hosted gitlab clusters as well as the cloud-hosted cluster by
			gitlab. 
	`
	fmt.Printf("%s\n", gitlabUrlHelp)
	os.Exit(0)
}

func VaultAddressHelp() {
	var vaultAddressHelp string = `
	Usage: 		gitlab-reporter --vault_address https://vault.example.com
			gitlab-reporter -vault_address https://vault.example.com
			gitlab-reporter vault_address https://vault.example.com
			gitlab-reporter -va https://vault.example.com

			Must be used in conjunction with the --namespace, --path, and --secret_key flags.
	`
	fmt.Printf("%s\n", vaultAddressHelp)
	os.Exit(0)
}

func NamespaceHelp() {
	var namespaceHelp string = `
	Usage: 		gitlab-reporter --namepsace namespace
			gitlab-reporter -namespace namespace
			gitlab-reporter namespace namespace
			gitlab-reporter -n namespace

			Must be used in conjunction with the --vault_address, --path, and --secret_key flags.
	`
	fmt.Printf("%s\n", namespaceHelp)
	os.Exit(0)
}

func PathHelp() {
	var pathHelp string = `
	Usage: 		gitlab-reporter --path pass the path where the secret is contained in HashiCorp Vault
			gitlab-reporter -path pass the path where the secret is contained in HashiCorp Vault
			gitlab-reporter path pass the path where the secret is contained in HashiCorp Vault
			gitlab-reporter -p pass the path where the secret is contained in HashiCorp Vault
			
			Must be used in conjunction with the --vault_address, --namespace, and --secret_key flags.
	`
	fmt.Printf("%s\n", pathHelp)
	os.Exit(0)
}

func SecretKeyHelp() {
	var secretKeyHelp string = `
	Usage:		gitlab-reporter --secret_key pass the Key from HashiCorp Vault Secret
			gitlab-reporter -secret_key pass the Key from HashiCorp Vault Secret
			gitlab-reporter secret_key pass the key from the HashiCorp Vault Secret
			gitlab-reporter -sk pass the key from the HashiCorp Vault Secret

			Must be used in conjunction with the --vault_address, --namespace, and --path.
	`
	fmt.Printf("%s\n", secretKeyHelp)
	os.Exit(0)
}

func ApiTokenHelp() {
	var apiTokenHelp string = `
	Usage: 		gitlab-reporter --api_token pass the secret token from either GitLab or Hashicorp Vault
			gitlab-reporter -api_token pass the secret token from either GitLab or HashiCorp Vault
			gitlab-reporter api_token pass the secret token from either GitLab or HashiCorp Vault
			gitlab-reporter -at pass the secret token from either GitLab or HashiCorp Vault

			Required Field that must be passed through either to use API token to decrypt Hashicorp Vault 
			or to directly read the GitLab API directly.
	`
	fmt.Printf("%s\n", apiTokenHelp)
	os.Exit(0)
}

func ReportHelp() {
	var reportHelp string = `
	Usage: 		gitlab-reporter --report <put your report name here>
			gitlab-reporter -report <put your report name here>
			gitlab-reporter report <put your report name here>
			gitlab-reporter -r <put your report name here>
	`
	fmt.Printf("%s\n", reportHelp)
}

func ExportDirectoryHelp() {
	var exportDirectoryHelp string = `
	Usage: 		gitlab-reporter --export_dir <put path you want your report saved>
			gitlab-reporter -export_dir <put path you want your report saved>
			gitlab-reporter export_dir <put path you want your report saved>
			gitlab-reporter -ed <put path you want your report saved>
	`
	fmt.Printf("%s\n", exportDirectoryHelp)
	os.Exit(0)
}

func OutputFormatHelp() {
	var outputFormatHelp string = `
	Usage: 		gitlab-reporter --output_format <put the format you want your report saved as)
			gitlab-reporter -output_format <put the format you want your report saved as)
			gitlab-reporter output_format <put the format you want your report saved as)
			gitlab-reporter -of <put the format you want your report saved as)
	`
	fmt.Printf("%s\n", outputFormatHelp)
	os.Exit(0)
}

func DisplayVersion() {
	var versionNumber string = `
GitLab Reporter
Version: 0.01
Build: alpha_build
`
	fmt.Printf("%s\n", versionNumber)
	os.Exit(0)
}
