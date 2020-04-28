package aws

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestResourcesFuncNaming(t *testing.T) {
	var fileNames []string
	err := filepath.Walk(".", func(fileName string, info os.FileInfo, err error) error {
		if strings.HasPrefix(fileName, "resources_") {
			fileNames = append(fileNames, fileName)
		}
		return nil
	})

	if err != nil {
		t.Errorf("Error reading filenames: %v", err)
	}

	for _, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			t.Errorf("Error opening file: %v", err)
		}
		defer file.Close()

		var serviceName string
		var functionNames []string
		scanner := bufio.NewScanner(file)
		regexServiceName := regexp.MustCompile(`"github\.com/aws/aws-sdk-go-v2/service/(.*)"`)
		regexFunctionName := regexp.MustCompile(`func (.*)\(.*\).*\(.*\).*\{`)
		for scanner.Scan() {
			matchesServiceName := regexServiceName.FindStringSubmatch(scanner.Text())
			if len(matchesServiceName) > 0 {
				serviceName = matchesServiceName[1]
			}

			matchesFunctionName := regexFunctionName.FindStringSubmatch(scanner.Text())
			if len(matchesFunctionName) > 0 {
				functionNames = append(functionNames, matchesFunctionName[1])
			}
		}

		if serviceName == "" {
			t.Errorf("Did not find a service name for file: %v", fileName)
		}

		for _, functionName := range functionNames {
			if !strings.HasPrefix(strings.ToLower(functionName), fmt.Sprintf("get%s", serviceName)) {
				t.Errorf("Function %s in file %s has incorrect name", functionName, fileName)
			}
		}
	}
}

func TestResourcesFileNaming(t *testing.T) {
	var fileNames []string
	err := filepath.Walk(".", func(fileName string, info os.FileInfo, err error) error {
		if strings.HasPrefix(fileName, "resources_") {
			fileNames = append(fileNames, fileName)
		}
		return nil
	})

	if err != nil {
		t.Errorf("Error reading filenames: %v", err)
	}

	for _, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			t.Errorf("Error opening file: %v", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		regexServiceName := regexp.MustCompile(`"github\.com/aws/aws-sdk-go-v2/service/(.*)"`)
		for scanner.Scan() {
			matchesServiceName := regexServiceName.FindStringSubmatch(scanner.Text())
			if len(matchesServiceName) > 0 {
				if !strings.HasSuffix(fileName, fmt.Sprintf("%s.go", matchesServiceName[1])) {
					t.Errorf("Filename %s has incorrect name", fileName)
				}
			}
		}
	}
}

func TestResourcesFileVariableAssignmentNaming(t *testing.T) {
	var fileNames []string
	err := filepath.Walk(".", func(fileName string, info os.FileInfo, err error) error {
		if strings.HasPrefix(fileName, "resources_") {
			fileNames = append(fileNames, fileName)
		}
		return nil
	})

	if err != nil {
		t.Errorf("Error reading filenames: %v", err)
	}

	for _, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			t.Errorf("Error opening file: %v", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		regexVariableAssignment := regexp.MustCompile(`	(.*) := (.*)\(client\)`)
		for scanner.Scan() {
			matchesVariableAssignment := regexVariableAssignment.FindStringSubmatch(scanner.Text())
			if len(matchesVariableAssignment) > 0 {
				functionName0 := strings.ToLower(fmt.Sprintf("get%s", matchesVariableAssignment[1]))
				functionName1 := strings.ToLower(matchesVariableAssignment[2])
				functionNames := strings.ReplaceAll(functionName0, ", ", "and")
				if !(functionName0 == functionName1) && !(functionNames == functionName1) {
					t.Errorf("Variable %s has incorrect name", matchesVariableAssignment[1])
				}
			}
		}
	}
}
