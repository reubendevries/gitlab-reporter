package reports

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func ValidReportCheck(s string) (string, error) {
	reports := []string{"list_users", "list_active_users", "list_blocked_users", "list_external_users", "list_users_using_2FA", "list_groups", "list_group_projects"}
	for _, v := range reports {
		if v == s {
			return s, nil
		}
		continue
	}
	return s, errors.New("report name passed is not on list of valid reports")
}

func ExecuteReport(m map[string]string) (string, error) {
	var result string
	switch m["Report"] {
	case "list_users":
		result, err := listUsers(m)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		return result, nil
	case "list_active_users":
		result, err := listActiveUsers(m)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		return result, nil
	case "list_blocked_users":
		result, err := listBlockedUsers(m)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		return result, nil
	case "list_external_users":
		result, err := listExternalUsers(m)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		return result, nil
	case "list_users_using_2FA":
		result, err := listUsersUsing2FA(m)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		return result, nil
	case "list_groups":
		result, err := listGroups(m)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		return result, nil
	case "list_group_projects":
		result, err := listGroupProjects(m)
		if err != nil {
			log.Fatalf("%s\n", err)
			os.Exit(1)
		}
		return result, nil
	default:
		return result, errors.New("report not found, exiting application")
	}
}

func listUsers(m map[string]string) (string, error) {
	var result string
	var bearer = "Bearer " + m["Token"]
	lu := "/api/v4/users"
	req, err := http.NewRequest("GET", m["Url"]+lu, nil)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	h := req.Header
	h.Add("Authorization", bearer)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	defer response.Body.Close()
	totalPages := (response.Header.Get("x-total-pages"))
	pages, err := strconv.Atoi(totalPages)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	for i := 1; i >= pages; i++ {
		q := url.Values{}
		q.Add("pages=", strconv.Itoa(i))
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		result = string(body)
	}
	return result, nil
}

func listActiveUsers(m map[string]string) (string, error) {
	var result string
	return result, nil
}

func listBlockedUsers(m map[string]string) (string, error) {
	var result string
	return result, nil
}

func listExternalUsers(m map[string]string) (string, error) {
	var result string
	return result, nil
}

func listUsersUsing2FA(m map[string]string) (string, error) {
	var result string
	return result, nil
}

func listGroups(m map[string]string) (string, error) {
	var result string
	return result, nil
}

func listGroupProjects(m map[string]string) (string, error) {
	var result string
	return result, nil
}
