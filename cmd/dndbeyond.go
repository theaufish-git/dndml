/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/theaufish-git/dndml/internal/dndml"
)

// dndbeyondCmd represents the dndbeyond command
var dndbeyondCmd = &cobra.Command{
	Use:   "dndbeyond",
	Short: "Converts structured yaml files into source blocks suitable for pasting into dndbeyond.",
	Long:  "Converts structured yaml files into source blocks suitable for pasting into dndbeyond.",
	Run:   dndbeyond,
}

func dndbeyond(cmd *cobra.Command, args []string) {
	parser := dndml.NewUnmarshaller()

	objs, err := unmarshal(cmd, args, parser)
	if err != nil {
		log.Fatal("cannot parse files:", err)
	}

	for k, v := range objs {
		log.Printf("---\n%s => %+v\n", k, v)
	}
}
