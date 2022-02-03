package checkStatus

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

func CheckGitlabStatus(u, t string) error {
	apiEndpoint := u + "/api/v4/version"
	var bearer = "Bearer " + t
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		log.Fatalf("%s\n", err)
		os.Exit(1)
	}
	h := req.Header
	h.Add("Authorization", bearer)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("%s\n", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	switch response.StatusCode {
	case 400:
		return errors.New(response.Status)
	case 401:
		return errors.New(response.Status)
	case 403:
		return errors.New(response.Status)
	case 404:
		return errors.New(response.Status)
	case 405:
		return errors.New(response.Status)
	case 409:
		return errors.New(response.Status)
	case 412:
		return errors.New(response.Status)
	case 422:
		return errors.New(response.Status)
	case 429:
		return errors.New(response.Status)
	case 500:
		return errors.New(response.Status)
	}
	return nil
}

func CheckVaultStatus(a, t string) error {
	apiEndpoint := a + "/v1/secret?help=1"
	fmt.Println(apiEndpoint)
	var bearer = "Bearer " + t
	req, err := http.NewRequest("GET", apiEndpoint, nil)
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
	switch response.StatusCode {
	case 400:
		return errors.New(response.Status)
	case 403:
		return errors.New(response.Status)
	case 404:
		return errors.New(response.Status)
	case 405:
		return errors.New(response.Status)
	case 412:
		return errors.New(response.Status)
	case 429:
		return errors.New(response.Status)
	case 473:
		return errors.New(response.Status)
	case 500:
		return errors.New(response.Status)
	case 502:
		return errors.New(response.Status)
	case 503:
		return errors.New(response.Status)
	}
	return nil
}
