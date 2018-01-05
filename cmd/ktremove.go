package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var principal string
var kvno, all, old bool

// ktremoveCmd represents the ktremove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a Service Principal Name (SPN) from a Keytab file",
	Long: `Remove a Service Principal Name (SPN) from a Keytab file

gokutil remove --keytab /etc/krb5.keytab --
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		MyCmd()
		// fmt.Println("Jason")
		// kt, err := keytab.Load("/etc/krb5.keytab")
		// if err != nil {
		// 	panic(fmt.Sprintf("%v", err))
		// }
		// fmt.Printf("%+v", kt)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().StringVar(&principal, "principal", "", "Specifies the principal to be removed from the keytab file.")
	removeCmd.Flags().BoolVar(&kvno, "kvno", false, "Removes all entries for the specified principal whose key version number matches kvno.")
	removeCmd.Flags().BoolVar(&all, "all", false, "Removes all entries for the specified principal.")
	removeCmd.Flags().BoolVar(&old, "old", false, "Removes all entries for the specified principal, except those principals with the highest key version number.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ktremoveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ktremoveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func MyCmd() {
	fmt.Println("ktremove called")
}
