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

type Accounts struct {
	Accounts []Account
}

type Account struct {
	CloudName        string   `json:"cloudName"`
	HomeTenantId     string   `json:"homeTenantId"`
	Id               string   `json:"id"`
	IsDefault        bool     `json:"isDefault"`
	ManagedByTenants []string `json:"managedByTenants"`
	Name             string   `json:"name"`
	State            string   `json:"state"`
	TenantId         string   `json:"tenantId"`
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
	if err != nil {
		panic(err)
	}

	return []byte(out)
}

func DisplayAzureAccountsDisplayName() {
	d := GetAzureAccounts()
	var Accounts []Account
	err := json.Unmarshal(d, &Accounts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", Accounts)
}

func SelectAzureAccountsDisplayName() {
	d := GetAzureAccounts()
	var Accounts []Account
	err := json.Unmarshal(d, &Accounts)
	if err != nil {
		panic(err)
	}
	idx, _ := fuzzyfinder.Find(
		Accounts,
		func(i int) string {
			return fmt.Sprintf("%s", Accounts[i].Name)
		}, fuzzyfinder.WithPreviewWindow(func(i, width, _ int) string {
			if i == -1 {
				return "no results"
			}
			s := fmt.Sprintf("%s is selected", Accounts[i].Name)
			if width < len([]rune(s)) {
				return Accounts[i].Name
			}
			return s
		}))
	fmt.Println(Accounts[idx])
}
