// +build !integration

package aztx_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/riweston/aztx/cmd/aztx"
)

func TestReadFile(t *testing.T) {
	value := aztx.ReadAzureProfile("azureProfile.json")
	equals(t, "Test Subscription 1", value.Subscriptions[0].Name)
}

func TestWriteFile(t *testing.T) {
	inputFile := aztx.ReadAzureProfile("azureProfile.json")
	uuid, err := uuid.Parse("9e7969ef-4cb8-4a2d-959f-bfdaae452a3d")
	if err != nil {
		panic(err)
	}
	outFile := aztx.WriteAzureProfile(inputFile, uuid, "azureProfile_test.json")
	ok(t, outFile)
	value := aztx.ReadAzureProfile("azureProfile_test.json")
	equals(t, true, value.Subscriptions[0].IsDefault)
	equals(t, false, value.Subscriptions[1].IsDefault)
	equals(t, false, value.Subscriptions[2].IsDefault)
}
