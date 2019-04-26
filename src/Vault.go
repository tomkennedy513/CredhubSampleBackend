package src

import (
	"fmt"
	"github.com/hashicorp/vault/api"
)

func SetValue(path string, value string) *api.Secret {
	client := getClient()

	payload := map[string]interface{}{
		"data": map[string]interface{}{
			path: value,
		},
		"options": map[string]interface{}{},
	}
	res, err := client.Write(fmt.Sprintf("/secret/data/%s", path), payload)
	if err != nil {
		fmt.Println(err)
	}

	return res
}

func GetValue(path string) *api.Secret {
	client := getClient()

	secret, err := client.Read(fmt.Sprintf("/secret/data/%s", path))
	if err != nil {
		fmt.Println(err)
	}

	return secret
}

func DeleteValue(path string) *api.Secret {
	client := getClient()

	secret, err := client.Delete(fmt.Sprintf("/secret/data/%s", path))
	if err != nil {
		fmt.Println(err)
	}

	return secret
}

func getClient() *api.Logical {
	config := api.Config{
		Address: "http://127.0.0.1:8200",
	}
	client, _ := api.NewClient(&config)
	client.SetToken("s.OypjFJxXN3OlGqcHVZhtY9He")
	return client.Logical()
}

