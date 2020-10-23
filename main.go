package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Repository struct {
	ID       int    `json:"id"`
	Location string `json:"location"`
}

type Repositories []Repository

var baseAPIURL string = "https://gitlab.com/api/v4"
var accessToken string = os.Getenv("ACCESS_TOKEN")
var projectID string = os.Getenv("PROJECT_ID")
var imageTag string = os.Getenv("IMAGE_TAG")
var imageLocation string = os.Getenv("IMAGE_LOCATION")

func main() {
	client := http.Client{}
	registryID, err := getRegistryRepositoryID(imageLocation, &client)
	if err != nil {
		log.Fatal(err)
	}

	err = deleteRegistryRepositoryTag(registryID, &client)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func deleteRegistryRepositoryTag(registryID int, client *http.Client) error {
	url := fmt.Sprintf("%s/projects/%s/registry/repositories/%d/tags/%s", baseAPIURL, projectID, registryID, imageTag)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal("http client could not be created.")
	}
	req.Header.Add("PRIVATE-TOKEN", accessToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("error while deleting image.")
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return nil
	}
	if resp.StatusCode != 200 {
		return errors.New("error while deleting image")
	}
	return nil
}

func getRegistryRepositoryID(imageLocation string, client *http.Client) (int, error) {
	url := fmt.Sprintf("%s/projects/%s/registry/repositories", baseAPIURL, projectID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("http client could not be created.")
	}
	req.Header.Add("PRIVATE-TOKEN", accessToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("error while fetching repositories.")
	}
	defer resp.Body.Close()
	responseBody := Repositories{}
	json.NewDecoder(resp.Body).Decode(&responseBody)
	for _, repository := range responseBody {
		if repository.Location == imageLocation {
			return repository.ID, nil
		}
	}
	return 0, errors.New("no repository id found for given image location")
}
