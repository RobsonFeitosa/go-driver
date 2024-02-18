package cmd

import (
	"bytes"
	"encoding/json"
	"log"
	"os"

	"github.com/RobsonFeitosa/go-driver/internal/users"
	"github.com/RobsonFeitosa/go-driver/pkg/requests"
	"github.com/spf13/cobra"
)

func create() *cobra.Command {
	var (
		name  string
		login string
		pass  string
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Cria uma novo usuário",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" || login == "" || pass == "" {
				log.Println("Nome, login e senha são obrigatório")
				os.Exit(1)
			}

			folder := users.User{
				Name:     name,
				Login:    login,
				Password: pass,
			}

			var body bytes.Buffer
			err := json.NewEncoder(&body).Encode(folder)
			if err != nil {
				log.Printf("%v", err)
				os.Exit(1)
			}

			_, err = requests.Post("/users", &body)
			if err != nil {
				log.Printf("%v", err)
				os.Exit(1)
			}

			log.Println("Usuário criada com sucesso!")
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Nome da pasta")
	cmd.Flags().StringVarP(&login, "login", "l", "", "Login da pasta")
	cmd.Flags().StringVarP(&pass, "pass", "p", "", "Senha da pasta")

	return cmd
}
