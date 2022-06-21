/*
Copyright © 2021 Richard Weston github@riweston.io

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package aztx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
	"github.com/ktr0731/go-fuzzyfinder"
)

const (
	InfoColor    = "\033[0;32m%s\033[0m"
	NoticeColor  = "\033[0;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

type Subscription struct {
	EnvironmentName  string    `json:"environmentName"`
	HomeTenantID     uuid.UUID `json:"homeTenantId"`
	ID               uuid.UUID `json:"id"`
	IsDefault        bool      `json:"isDefault"`
	ManagedByTenants []string  `json:"managedByTenants"`
	Name             string    `json:"name"`
	State            string    `json:"state"`
	TenantID         uuid.UUID `json:"tenantId"`
	User             struct {
		Name        string `json:"name"`
		AccountType string `json:"type"`
	} `json:"user"`
}

type File struct {
	InstallationID uuid.UUID      `json:"installationId"`
	Subscriptions  []Subscription `json:"subscriptions"`
}

func SelectAzureAccountsDisplayName() {
	home, errHome := os.UserHomeDir()
	if errHome != nil {
		panic(errHome)
	}
	azureProfile := home + "/.azure/azureProfile.json"
	d := ReadAzureProfile(azureProfile)
	currentCtx := ReadAzureProfileDefault(d)

	idx, errFind := fuzzyfinder.Find(
		d.Subscriptions,
		func(i int) string {
			return d.Subscriptions[i].Name
		})
	if errFind != nil {
		panic(errFind)
		fuzzyfinder.WithHeader(currentCtx))
	}

	errWrite := WriteAzureProfile(d, d.Subscriptions[idx].ID, azureProfile)
	if errWrite != nil {
		panic(errWrite)
	}
	msg := fmt.Sprintf("Set Context: %s (%s)\n", d.Subscriptions[idx].Name, d.Subscriptions[idx].ID)
	fmt.Printf(InfoColor, msg)
}

func ReadAzureProfile(file string) File {
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, errByte := ioutil.ReadAll(jsonFile)
	if errByte != nil {
		fmt.Println(errByte)
	}
	byteValue = bytes.TrimPrefix(byteValue, []byte("\xef\xbb\xbf"))
	var jsonData File
	errJSON := json.Unmarshal(byteValue, &jsonData)
	if errJSON != nil {
		fmt.Println(err)
	}

	return jsonData
}

func ReadAzureProfileDefault(file File) (subscription string) {
	var subscriptionName string
	var subscriptionId uuid.UUID

	for idx := range file.Subscriptions {
		if file.Subscriptions[idx].IsDefault {
			subscriptionName = file.Subscriptions[idx].Name
			subscriptionId = file.Subscriptions[idx].ID
		}
	}
	return fmt.Sprintf("Current Context: %s (%s)", subscriptionName, subscriptionId)
}

func WriteAzureProfile(file File, id uuid.UUID, outFile string) error {
	for idx := range file.Subscriptions {
		if file.Subscriptions[idx].ID == id {
			file.Subscriptions[idx].IsDefault = true
		} else {
			file.Subscriptions[idx].IsDefault = false
		}
	}

	byteValue, err := json.Marshal(&file)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outFile, byteValue, 0600)
	return err
}
