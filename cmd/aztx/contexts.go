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
	"os/exec"

	"github.com/ktr0731/go-fuzzyfinder"
)

type Account struct {
	CloudName        string   `json:"cloudName"`
	HomeTenantID     string   `json:"homeTenantId"`
	ID               string   `json:"id"`
	IsDefault        bool     `json:"isDefault"`
	ManagedByTenants []string `json:"managedByTenants"`
	Name             string   `json:"name"`
	State            string   `json:"state"`
	TenantID         string   `json:"tenantId"`
	User             struct {
		Name        string `json:"name"`
		AccountType string `json:"type"`
	} `json:"user"`
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
	var Accounts []Account
	err := json.Unmarshal(d, &Accounts)
	if err != nil {
		panic(err)
	}
	idx, errFind := fuzzyfinder.Find(
		Accounts,
		func(i int) string {
			return string(Accounts[i].Name)
		})
	if errFind != nil {
		panic(errFind)
	}
	SetAzureAccountContext(Accounts[idx].Name)
	fmt.Println("Azure Context;\n", Accounts[idx].Name, "-", Accounts[idx].ID)
}
