package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/amirvejahat/grafana-teamsync/pkg/grafana"
	"github.com/amirvejahat/grafana-teamsync/pkg/ldapclient"
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

		// TODO: add viper config manager
		// if len(args) == 0 {
		// 	cmd.Help()
		// 	os.Exit(0)
		// }
		//list of users and groups from ldap
		// create users based on ldap accounts
		// create teams based on ldap groups
		// create folders based on ldap groups

		// checking ldap connections
		myLdapClient := &ldapclient.LDAPClient{
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
		// establish ldap connection
		myLdapClient.CreateLdapConnection()

		// obtain list of ldap users
		// ldapUsersList, err := myLdapClient.ListLDAPUsers()
		// if err != nil {
		// 	log.Fatal(err)
		// 	return
		// }

		// obtain list of ldap groups
		ldapGroupsList, err := myLdapClient.ListLDAPGroups()
		mar, err := json.MarshalIndent(ldapGroupsList, "", " ")
		fmt.Println(string(mar))
		if err != nil {
			log.Fatal(err)
			return
		}

		// create a grafana client
		gclient, err := grafana.NewClient(grafanaUrl, grafanaToken, 30*time.Second, 3, nil)
		if err != nil {
			log.Fatal(err)
			return
		}

		// create a folder and a team for each group
		for _, group := range ldapGroupsList {
			gclient.CreateFolder(group.GroupName)
			// check whether team exists or not
			gclient.AddTeam(group.GroupName, fmt.Sprintf("%s@gmail.com", group.GroupName), oraganizationID)
			// for _, member := range group.MemberUid {
			// first we need to get the user id.
			// gclient.AddTeamMember(team.ID, )
			//}
		}

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
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.grafana-teamsync.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringVar(&grafanaUrl, "grafana-url", "http://localhost:3000", "Grafana url")
	rootCmd.PersistentFlags().StringVar(&grafanaToken, "grafana-token", "glsa_eP4TgbMHFRTRn6Ff4Nw0HgyQjHTzWSmi_6c75c778", "Grafana token")
	rootCmd.PersistentFlags().BoolVarP(&daemon, "daemon", "d", false, "running as daemon")
	//rootCmd.MarkPersistentFlagRequired("grafana-url")

}
