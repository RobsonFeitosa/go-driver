package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/RobsonFeitosa/go-driver/internal/users"
	"github.com/RobsonFeitosa/go-driver/pkg/requests"
	"github.com/spf13/cobra"
)

func update() *cobra.Command {
	var id int32
	var name string

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Atualizar o nome do usuário",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" || id <= 0 {
				log.Println("Nome do usuário e ID são obrigatório")
				os.Exit(1)
			}

			folder := users.User{Name: name}

			var body bytes.Buffer
			err := json.NewEncoder(&body).Encode(folder)
			if err != nil {
				log.Printf("%x", err)
				os.Exit(1)
			}

			path := fmt.Sprintf("/users/%d", id)
			_, err = requests.AuthenticatedPut(path, &body)
			if err != nil {
				log.Printf("%x", err)
				os.Exit(1)
			}

			log.Println("Usuário atualizado com sucesso!")
		},
	}

	cmd.Flags().Int32VarP(&id, "id", "", 0, "ID do usuário")
	cmd.Flags().StringVarP(&name, "name", "n", "", "Nome do usuário")

	return cmd
}
