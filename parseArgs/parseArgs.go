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

func getGitLabUrl(s string) (string, error) {
	switch strings.HasPrefix(s, "https") {
	case false:
		return s, errors.New("please use in a secure gitlab url")
	}
	return s, nil
}

// TODO add a funtion that will allow us to process an API token from some sort of Vault.
//func getVaultToken(s string) (string, error) {
//	switch strings.HasPrefix(s, "https") {
//	case false:
//		return s, errors.New("please put in a secure hashicorp vault address")
//	}
//	response, err := http.Get(s)
//	if err != nil {
//		log.Fatalf("%s\n", err)
//	}
//	responseData, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		log.Fatalf("%s\n", err)
//	}
//	return string(responseData), nil
//}

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
//func setExportDirectory(s string) (string, error) {
//	return s, nil
//}

// TODO add a function that will allow the user to change the output of the report.
//func setOutputFormat(s string) (string, error) {
//	return s, nil
//}

// "list_users", "list_active_users", "list_blocked_users", "list_external_users", "list_users_using_2FA", "list_groups", "list_group_projects"
func executeReport(r string, u string, t string) {
	switch r {
	case "list_users":
		reports.ListUsers(u, t)
		//if err != nil {
		//	log.Fatalf("%s\n", err)
		//}
		//fmt.Println(result)
	}
}

func displayHelp() {
	var help string = `
	Usage: gitlab-reporter [command]

	An application that helps call the GitLab API and then
	exports the output into various formats.

	The available commands for execution are listed below:

	Commands:
	--gitlab_url, -gitlab_url, gitlab_url, -gu			Passthrough your GitLab Url
	--vault_address, -vault_address, vault_address -va		Passthrough your Hashicorp Vault Address
	--api_token, -api_token, api_token, -at				Passthrough your GitLab API Token
	--report, -report, report, -r					Passthrough the Report you wish to call
	--export_dir, -export_dir, export_dir, -ed			Allows you to select the export folder.
	--ouput_format, -output_format, output_format, -of 		Allows you to choose the output format of the report
	--help, -help, help, -h 					Shows the Help output
	--version, -version, version, -v				Shows the Version output
	`
	fmt.Printf("%s\n", help)
}

func displayVersion() {
	var versionNumber string = `
GitLab Reporter
Version: 0.01
Build: alpha_build
`
	fmt.Printf("%s\n", versionNumber)
}

func ParseArgs() {
	resp := make(map[string]string)
	args, err := getArgs(os.Args)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	for i, v := range args {
		switch v {
		case "--gitlab_url", "-gitlab_url", "gitlab_url", "-gu":
			response, err := getGitLabUrl(args[i+1])
			if err != nil {
				log.Fatalf("%s\n", err)
			}
			resp["Url"] = response
		// TODO add a funtion that will allow us to process an API token from some sort of Vault.
		//case "--vault_address", "-vault_address", "vault_address", "-va":
		//	response, err := getVaultToken(args[i+1])
		//	if err != nil {
		//		log.Fatalf("%s\n", err)
		//	}
		//	resp["Address"] = response
		case "--api_token", "-api_token", "api_token", "-at":
			response, err := apiToken(args[i+1])
			if err != nil {
				log.Fatalf("%s\n", err)
			}
			resp["Token"] = response
		case "--report", "-report", "report", "-r":
			response, err := validReportCheck(args[i+1])
			if err != nil {
				log.Fatalf("%s\n", err)
			}
			resp["Report"] = response
		//TODO add a function that will allow the user to change the directory where the report gets saved.
		//case "--export_dir", "-export_dir", "export_dir", "-ed":
		//	response, err := setExportDirectory(args[i+1])
		//	if err != nil {
		//		log.Fatalf("%s\n", err)
		//	}
		//	resp["Directory"] = response
		//TODO add a function that will allow the user to change the output of the report.
		//case "--ouput_format", "-output_format", "output_format", "-of":
		//	response, err := setOutputFormat(args[i+1])
		//	if err != nil {
		//		log.Fatalf("%s\n", err)
		//	}
		//	resp["Format"] = response
		case "--help", "-help", "help", "-h":
			displayHelp()
		case "--version", "-version", "version", "-v":
			displayVersion()
		}
	}
	if v, found := resp["Report"]; found {
		executeReport(v, resp["Url"], resp["Token"])
	}
}
