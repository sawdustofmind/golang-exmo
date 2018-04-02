package tests

import (
	"os"

	exmo "github.com/asxcandrew/golang-exmo/v1"
)

var (
	client *exmo.Client
)

func init() {
	key := os.Getenv("EXMO_API_KEY")
	secret := os.Getenv("EXMO_API_SECRET")
	client = exmo.NewClient().Auth(key, secret)
}
