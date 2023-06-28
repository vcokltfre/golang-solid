package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/vcokltfre/golang-solid/src/api"
)

func init() {
	godotenv.Load()

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logrus.Info("Initialised logging.")
}

func main() {
	bind, ok := os.LookupEnv("BIND")
	if !ok {
		bind = ":8080"
	}

	if err := api.Start(bind); err != nil {
		logrus.Fatal(err)
	}
}
