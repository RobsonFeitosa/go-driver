package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/RobsonFeitosa/go-driver/internal/users"
	"github.com/RobsonFeitosa/go-driver/pkg/requests"
	"github.com/spf13/cobra"
)

func list() *cobra.Command {
	var id int32

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lista de usuários",
		Run: func(cmd *cobra.Command, args []string) {
			path := "/users"
			if id > 0 {
				path = fmt.Sprintf("/users/%d", id)
			}

			data, err := requests.AuthenticatedGet(path)
			if err != nil {
				log.Printf("%v", err)
				os.Exit(1)
			}

			var u users.User
			err = json.Unmarshal(data, &u)
			if err != nil {
				log.Printf("%v", err)
				os.Exit(1)
			}

			log.Println(u.Name)
			log.Println(u.Login)
			log.Println(u.Password)
		},
	}

	cmd.Flags().Int32VarP(&id, "id", "", 0, "ID da pasta")

	return cmd
}
