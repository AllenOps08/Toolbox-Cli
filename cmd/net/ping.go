/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra" // Importing the cobra library for creating the CLI
)

var urlPath string
// Creating a variable to store the client with a timeout of 2 seconds
var client = http.Client{
   Timeout: time.Second * 2,
}

// Defining the ping function which takes a domain as an argument and returns the status code and an error
func ping(domain string) (int, error) {
   // Constructing the URL
   url := "http://" + domain
   // Creating a new request with the method "HEAD"
   req, err := http.NewRequest("HEAD", url, nil)
   if err != nil {
   	return 0, err
   }
   // Sending the request with the client and storing the response
   resp, err := client.Do(req)
   if err != nil {
   	return 0, err
   }
   // Closing the response body
   resp.Body.Close()
   // Returning the status code
   return resp.StatusCode, nil
}

// Creating the ping command
var pingCmd = &cobra.Command{
   Use:   "ping",
   Short: "This pings a remote URL and returns the response",
   Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
   // Defining the run logic for the ping command
   Run: func(cmd *cobra.Command, args []string) {
   	// Getting the URL from the flag
   	urlPath := cmd.Flag("url").Value.String()
   	// Calling the ping function and storing the result
   	resp, err := ping(urlPath)
   	if err != nil {
   		fmt.Println(err)
   	} else {
   		fmt.Println(resp)
   	}
   },
}

// Initializing the command and adding the ping command to it
func init() {
   pingCmd.Flags().StringVar(&urlPath, "url", "u", "The url to parse")
   if err := pingCmd.MarkFlagRequired("url"); err != nil {
   	fmt.Println(err)
   }
   NetCmd.AddCommand(pingCmd)
   // Here you will define your flags and configuration settings.

   // Cobra supports Persistent Flags which will work for this command
   // and all subcommands, e.g.:
   // pingCmd.PersistentFlags().String("foo", "", "A help for foo")

   // Cobra supports local flags which will only run when this command
   // is called directly, e.g.:
   // pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}