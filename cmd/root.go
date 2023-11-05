/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/amirvejahat/grafana-teamsync/pkg/gldap"
	"github.com/spf13/cobra"
)

var (
	grafanaUrl   string
	ldapUrl      string
	grafanaToken string
	ldapDn       string
	ldapUsername string
	ldapPassword string
	daemon       bool
)

const (
	oraganizationID int64 = 1
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "grafana-teamsync",
	Short:   "Grafana Sync + ldap",
	Long:    `Grafana teamsync keeps your grafana dashboards and teams and folders in sync`,
	Version: "0.1",
	Run: func(cmd *cobra.Command, args []string) {
		// if len(args) == 0 {
		// 	cmd.Help()
		// 	os.Exit(0)
		// }
		// gclient, err := grafana.NewClient(grafanaUrl, grafanaToken, 30, 0, nil)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// gclient.GetAllUsers()
		// gclient.AddTeam("myteam", "xxx@gmail.com", 1)
		//gclient.GetAllFolders()
		//gclient.CreateFolder("mynewFolder")

		//list of users and groups from ldap
		// create users based on ldap accounts
		// create teams based on ldap groups
		// create folders based on ldap groups

		// checking ldap connections
		myLdapClient := &gldap.LDAPClient{
			LdapServer:         "localhost:389",
			BindDN:             "cn=admin,dc=xl,dc=com",
			BaseDN:             "dc=xl,dc=com",
			ServerName:         "myServerName",
			LdapPassword:       "password",
			InsecureSkipVerify: true,
			UseSSL:             false,
			SkipTLS:            true,
			Conn:               nil,
		}
		myLdapClient.CreateLdapConnection()
		result, err := myLdapClient.ListLDAPUsers()
		if err != nil {
			fmt.Println("error happening!", err)
		}
		fmt.Println(result)

	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.grafana-teamsync.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringVar(&grafanaUrl, "grafana-url", "http://localhost:3000", "Grafana url")
	rootCmd.PersistentFlags().StringVar(&grafanaToken, "grafana-token", "glsa_eP4TgbMHFRTRn6Ff4Nw0HgyQjHTzWSmi_6c75c778", "Grafana token")
	rootCmd.PersistentFlags().BoolVarP(&daemon, "daemon", "d", false, "running as daemon")
	// rootCmd.PersistentFlags().StringVar(&ldapUrl, "grafana-token", "", "Grafana token")
	//rootCmd.MarkPersistentFlagRequired("grafana-url")

}
