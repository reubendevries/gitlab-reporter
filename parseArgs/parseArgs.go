package parseArgs

import (
	"errors"
	"fmt"
	"gitlab-reporter/reports"
	"log"
	"os"
	"strings"
)

func getArgs(s []string) ([]string, error) {
	var args []string
	for i := 0; i < len(os.Args); i++ {
		if i != 0 {
			args = append(args, os.Args[i])
		}
	}
	return args, nil
}

// TODO add a funtion that will allow us to process an API token from some sort of Vault.
func getVaultToken(s, t string) (string, error) {
	var vaultToken string
	return vaultToken, nil
}

func apiToken(s string) (string, error) {
	return s, nil
}

func validReportCheck(s string) (string, error) {
	reports := []string{"list_users", "list_active_users", "list_blocked_users", "list_external_users", "list_users_using_2FA", "list_groups", "list_group_projects"}
	for _, v := range reports {
		if v == s {
			return s, nil
		}
		continue
	}
	return s, errors.New("report name passed is not on list of valid reports")
}

// TODO add a function that will allow the user to change the directory where the report gets saved.
func setExportDirectory(s string) (string, error) {
	return s, nil
}

// TODO add a function that will allow the user to change the output of the report.
func setOutputFormat(s string) (string, error) {
	return s, nil
}

func parseToken(m map[string]string) {
	if v, found := m["Address"]; found {
		token, err := getVaultToken(v, m["Token"])
		if err != nil {
			log.Fatalf("%s", err)
		}
		m["Token"] = token
		executeReport(m)
	} else {
		executeReport(m)
	}
}
func executeReport(m map[string]string) (string, error) {
	var result string
	switch m["Report"] {
	case "list_users":
		result, err := reports.ListUsers(m)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		return result, nil
	case "list_active_users":
		result, err := reports.ListActiveUsers(m)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		return result, nil
	case "list_blocked_users":
		result, err := reports.ListBlockedUsers(m)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		return result, nil
	case "list_external_users":
		result, err := reports.ListExternalUsers(m)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		return result, nil
	case "list_users_using_2FA":
		result, err := reports.ListUsersUsing2FA(m)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		return result, nil
	case "list_groups":
		result, err := reports.ListGroups(m)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		return result, nil
	case "list_group_projects":
		result, err := reports.ListGroupProjects(m)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		return result, nil
	default:
		return result, errors.New("report not found, exiting application")
	}
}

func defaultHelp() {
	var defaultHelp string = `
	Usage: gitlab-reporter [command]

	An application that helps call the GitLab API and then
	exports a report of the results into various formats.

	The available commands for execution are listed below:

	Commands:
	--gitlab_url, -gitlab_url, gitlab_url, -gu			Passthrough your GitLab Url
	--vault_address, -vault_address, vault_address -va		Passthrough your Hashicorp Vault Address
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
func gitlabUrlHelp() {
	var gitlabUrlHelp string = `
	Usage: 	gitlab-reporter --gitlab_url https://gitlab.example.com
			gitlab-reporter -gitlab_url https://gitlab.examplecom
			gitlab-reporter gitlab_url https://gitlab.examplecom
			gitlab-reporter -gu https://gitlab.example.com
	`
	fmt.Printf("%s\n", gitlabUrlHelp)
	os.Exit(0)
}

func vaultAddressHelp() {
	var vaultAddressHelp string = `
	Usage: 	gitlab-reporter --vault_address https://vault.example.com
			gitlab-reporter -vault_address https://vault.example.com
			gitlab-reporter vault_address https://vault.example.com
			gitlab-reporter -va https://vault.example.com
	`
	fmt.Printf("%s\n", vaultAddressHelp)
	os.Exit(0)
}

func apiTokenHelp() {
	var apiTokenHelp string = `
	Usage: 	gitlab-reporter --api_token <put your secret token here>
			gitlab-reporter -api_token <put your secret token here>
			gitlab-reporter api_token <put your secret token here>
			gitlab-reporter -at <put your secret token here>
	`
	fmt.Printf("%s\n", apiTokenHelp)
	os.Exit(0)
}

func reportHelp() {
	var reportHelp string = `
	Usage: 	gitlab-reporter --report <put your report name here>
			gitlab-reporter -report <put your report name here>
			gitlab-reporter report <put your report name here>
			gitlab-reporter -r <put your report name here>
	`
	fmt.Printf("%s\n", reportHelp)
}

func exportDirectoryHelp() {
	var exportDirectoryHelp string = `
	Usage: 	gitlab-reporter --export_dir <put path you want your report saved>
			gitlab-reporter -export_dir <put path you want your report saved>
			gitlab-reporter export_dir <put path you want your report saved>
			gitlab-reporter -ed <put path you want your report saved>
	`
	fmt.Printf("%s\n", exportDirectoryHelp)
	os.Exit(0)
}

func outputFormatHelp() {
	var outputFormatHelp string = `
	Usage: 	gitlab-reporter --output_format <put the format you want your report saved as)
			gitlab-reporter -output_format <put the format you want your report saved as)
			gitlab-reporter output_format <put the format you want your report saved as)
			gitlab-reporter -of <put the format you want your report saved as)
	`
	fmt.Printf("%s\n", outputFormatHelp)
	os.Exit(0)
}

func displayVersion() {
	var versionNumber string = `
GitLab Reporter
Version: 0.01
Build: alpha_build
`
	fmt.Printf("%s\n", versionNumber)
	os.Exit(0)
}

func ParseArgs() {
	data := make(map[string]string)
	args, err := getArgs(os.Args)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--gitlab_url", "-gitlab_url", "gitlab_url", "-gu":
			if strings.Contains(args[i+1], "help") {
				gitlabUrlHelp()
			} else {
				data["Url"] = args[i]
			}
		case "--vault_address", "-vault_address", "vault_address", "-va":
			fmt.Println(args[i])
			fmt.Println(args[i+1])
			if strings.Contains(args[i+2], "help") {
				vaultAddressHelp()
			} else {
				data["Address"] = args[i]
			}
		case "--api_token", "-api_token", "api_token", "-at":
			fmt.Println(args[i])
			fmt.Println(args[i+1])
			if strings.Contains(args[i+1], "help") {
				apiTokenHelp()
			} else {
				data["Token"] = args[i]
			}
		case "--report", "-report", "report", "-r":
			fmt.Println(args[i])
			fmt.Println(args[i+1])
			response, err := validReportCheck(args[i])
			if err != nil {
				log.Fatalf("%s\n", err)
			}
			if strings.Contains(args[i+1], "help") {
				reportHelp()
			} else {
				data["Report"] = response
			}
		//TODO add a function that will allow the user to change the directory where the report gets saved.
		case "--export_dir", "-export_dir", "export_dir", "-ed":
			fmt.Println(args[i])
			fmt.Println(args[i+1])
			if strings.Contains(args[i+1], "help") {
				exportDirectoryHelp()
			} else {
				data["Directory"] = args[i]
			}
		//TODO add a function that will allow the user to change the output of the report.
		case "--ouput_format", "-output_format", "output_format", "-of":
			fmt.Println(args[i])
			fmt.Println(args[i+1])
			if strings.Contains(args[i+1], "help") {
				outputFormatHelp()
			} else {
				data["Format"] = args[i+1]
			}
		case "--version", "-version", "version", "-v":
			displayVersion()
		case "--help", "-help", "help", "-h":
			defaultHelp()
		}
	}
	parseToken(data)
}
