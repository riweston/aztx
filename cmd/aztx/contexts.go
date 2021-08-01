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

type Subscription struct {
	CloudName        string    `json:"cloudName"`
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
	home, _ := os.UserHomeDir()
	azureProfile := home + "/.azure/azureProfile.json"
	d := ReadAzureProfile(azureProfile)

	idx, errFind := fuzzyfinder.Find(
		d.Subscriptions,
		func(i int) string {
			return d.Subscriptions[i].Name
		})
	if errFind == nil {
		panic(errFind)
	}

	WriteAzureProfile(d, d.Subscriptions[idx].ID, azureProfile)
	fmt.Print(d.Subscriptions[idx].Name, "\n", d.Subscriptions[idx].ID, "\n")
}

func ReadAzureProfile(file string) File {
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	byteValue = bytes.TrimPrefix(byteValue, []byte("\xef\xbb\xbf"))
	var jsonData File
	errJson := json.Unmarshal(byteValue, &jsonData)
	if errJson != nil {
		fmt.Println(err)
	}

	return jsonData
}

func WriteAzureProfile(file File, id uuid.UUID, outFile string) error {
	for idx, _ := range file.Subscriptions {
		if file.Subscriptions[idx].ID == id {
			file.Subscriptions[idx].IsDefault = true
		}
	}

	byteValue, err := json.Marshal(&file)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outFile, byteValue, 0666)
	return err
}
