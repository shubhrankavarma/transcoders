package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/amagimedia/transcoders/config"
	vault "github.com/hashicorp/vault/api"
)

// Function to read the config file from HashiCorp Vault
func GetHashiCorpVaultValues(myStruct *config.Properties) {

	// Create a new client
	config := vault.DefaultConfig()

	config.Address = os.Getenv("VAULT_ADDR")

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}

	// Authenticate
	client.SetToken(os.Getenv("VAULT_TOKEN"))

	// For every field in the struct, get the value of hashicorp_vault tag and get the value from vault
	// and set the value to the struct field
	for i := 0; i < reflect.TypeOf(*myStruct).NumField(); i++ {
		field := reflect.TypeOf(*myStruct).Field(i)
		tag := field.Tag.Get("hashicorp")
		fmt.Println(tag)
		if tag != "" {
			secret, err := client.KVv2(os.Getenv("secret")).Get(context.Background(), tag)
			if err != nil {
				log.Fatalf("unable to read secret: %v", err)
			}

			value, ok := secret.Data[tag].(string)
			if !ok {
				log.Fatalf("value type assertion failed: %T %#v", secret.Data[tag], secret.Data[tag])
			}

			reflect.ValueOf(myStruct).Elem().FieldByName(field.Name).SetString(value)
		}
	}
}
