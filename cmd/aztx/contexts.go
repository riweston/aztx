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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

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
	InstallationID uuid.UUID `json:"installationId"`
	Subscriptions  []Subscription
}

func GetAzureAccounts() []byte {
	binary, errLook := exec.LookPath("az")
	if errLook != nil {
		panic(errLook)
	}

	args := []string{"account", "list", "-o", "json"}

	out, err := exec.Command(binary, args...).CombinedOutput()
	// TODO: This currently breaks when selecting context
	if err != nil {
		panic(err)
	}

	return out
}

func SetAzureAccountContext(accountname string) {
	binary, errLook := exec.LookPath("az")
	if errLook != nil {
		panic(errLook)
	}

	args := []string{"account", "set", "--subscription", accountname}

	_, err := exec.Command(binary, args...).Output()
	if err != nil {
		panic(err.Error())
	}
}

func SelectAzureAccountsDisplayName() {
	d := GetAzureAccounts()
	var Accounts []Subscription
	err := json.Unmarshal(d, &Accounts)
	if err != nil {
		panic(err)
	}
	idx, errFind := fuzzyfinder.Find(
		Accounts,
		func(i int) string {
			return Accounts[i].Name
		})
	if errFind != nil {
		panic(errFind)
	}
	SetAzureAccountContext(Accounts[idx].Name)
	fmt.Print(Accounts[idx].Name, "\n", Accounts[idx].ID, "\n")
}

func ReadAzureProfile(file string) File {
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var jsonData File
	errJson := json.Unmarshal(byteValue, &jsonData)
	if errJson != nil {
		fmt.Println(err)
	}

	return jsonData
}

