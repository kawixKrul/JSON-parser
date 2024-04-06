package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Data struct {
	PolicyName     string `json:"PolicyName"`
	PolicyDocument struct {
		Version   string `json:"Version"`
		Statement []struct {
			Sid      string   `json:"Sid"`
			Effect   string   `json:"Effect"`
			Action   []string `json:"Action"`
			Resource string   `json:"Resource"`
		} `json:"Statement"`
	} `json:"PolicyDocument"`
}

func readFromFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return string(content), nil
}

func validateJsonStr(jsonStr string) (bool, error) {
	var jsonData Data
	err := json.Unmarshal([]byte(jsonStr), &jsonData)
	if err != nil {
		return false, err
	}

	if jsonData.PolicyName == "" || jsonData.PolicyDocument.Version == "" {
		return false, fmt.Errorf("missing required fields: PolicyName or PolicyDocument.Version")
	}

	if len(jsonData.PolicyDocument.Statement) == 0 {
		return false, fmt.Errorf("missing required field: PolicyDocument.Statement")
	}

	if jsonData.PolicyDocument.Statement[0].Resource == "" {
		return false, fmt.Errorf("missing required field: PolicyDocument.Statement[0].Resource")
	}

	return !(jsonData.PolicyDocument.Statement[0].Resource == "*"), nil
}
