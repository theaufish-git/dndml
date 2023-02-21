/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/theaufish-git/dndml/internal/dndml"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dndml",
	Short: "Converts structured yaml files into source blocks suitable for pasting into dndbeyond.",
	Long:  "Converts structured yaml files into source blocks suitable for pasting into dndbeyond.",
	Run:   run,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	log.SetPrefix("")
	log.SetFlags(0)

	//rootCmd.PersistentFlags().StringSliceP("include", "i", []string{"."}, "directories that contain referenced files")
	rootCmd.PersistentFlags().StringSliceP("patterns", "p", []string{"*.yml", "*.yaml"}, "patterns that determine which files should be processed")
	rootCmd.PersistentFlags().StringP("out-dir", "o", ".", "the directory that output files will be written to")
}

func run(cmd *cobra.Command, args []string) {
	patterns, err := cmd.PersistentFlags().GetStringSlice("patterns")
	if err != nil {
		log.Fatal("cannot read patterns:", err)
	}

	fnames := []string{}
	fnameset := map[string]struct{}{}
	for _, root := range args {
		filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			// if error or directory immediately return
			if err != nil || info.IsDir() {
				return err
			}

			for _, p := range patterns {
				// try to match pattern directly
				if matched, err := filepath.Match(p, info.Name()); err != nil {
					return err
				} else if matched {
					if _, ok := fnameset[path]; !ok {
						log.Print(path)
						fnames = append(fnames, path)
						fnameset[path] = struct{}{}
					}

					return nil
				}
			}
			return nil
		})
	}

	parser := dndml.NewParser()
	if err := parser.Parse(fnames...); err != nil {
		log.Fatal("could not parse:", err)
	}

	for k, v := range parser.Objects() {
		log.Printf("---\n%s => %+v\n", k, v)
	}
}
