package main

import (
	"log"

	authCmd "github.com/RobsonFeitosa/go-driver/internal/auth/cmd"
	filesCmd "github.com/RobsonFeitosa/go-driver/internal/files/cmd"
	folderCmd "github.com/RobsonFeitosa/go-driver/internal/folders/cmd"
	usersCmd "github.com/RobsonFeitosa/go-driver/internal/users/cmd"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{}

func main() {
	authCmd.Register(RootCmd)
	filesCmd.Register(RootCmd)
	folderCmd.Register(RootCmd)
	usersCmd.Register(RootCmd)

	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
