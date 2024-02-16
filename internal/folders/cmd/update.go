package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/RobsonFeitosa/go-driver/internal/folders"
	"github.com/RobsonFeitosa/go-driver/pkg/requests"
	"github.com/spf13/cobra"
)

func update() *cobra.Command {
	var id int32
	var name string

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Atualizar o nome de uma pasta",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" || id <= 0 {
				log.Println("Nome da pasta e ID são obrigatório")
				os.Exit(1)
			}

			folder := folders.Folder{Name: name}

			var body bytes.Buffer
			err := json.NewEncoder(&body).Encode(folder)
			if err != nil {
				log.Printf("%x", err)
				os.Exit(1)
			}

			path := fmt.Sprintf("/folders/%d", id)
			_, err = requests.AuthenticatedPut(path, &body)
			if err != nil {
				log.Printf("%x", err)
				os.Exit(1)
			}

			log.Println("Pasta atualizada com sucesso!")
		},
	}

	cmd.Flags().Int32VarP(&id, "id", "", 0, "ID da pasta")
	cmd.Flags().StringVarP(&name, "name", "n", "", "Nome da pasta")

	return cmd
}
