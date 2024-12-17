/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/jagretti/secretly/internal/client"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// replicateCmd represents the replicate command
var replicateCmd = &cobra.Command{
	Use:   "replicate",
	Short: "Starts the replication process",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: replicate,
}

func init() {
	rootCmd.AddCommand(replicateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// replicateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// replicateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func replicate(cmd *cobra.Command, args []string) {
	client, err := client.BuildKubernetesClient()
	if err != nil {
		println(err)
	}
	for {
		secretsClient := client.CoreV1().Secrets("default") // only default namespace for now
		secret, err := secretsClient.Get(context.TODO(), "test", v1.GetOptions{})
		if err != nil {
			fmt.Println("No secret found!")
			println(err)
		} else {
			fmt.Println("Secret found!")
			fmt.Print(secret)
		}
	}
}
