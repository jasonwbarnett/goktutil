package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var principal string
var kvno int
var all, old bool

// ktremoveCmd represents the ktremove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a Service Principal Name (SPN) from a Keytab file",
	Long: `Remove a Service Principal Name (SPN) from a Keytab file

Remove all associated keys:
  gokutil remove --keytab /etc/krb5.keytab --principal HTTP/MYHOST --all

Remove old associated keys:
  gokutil remove --keytab /etc/krb5.keytab --principal HTTP/MYHOST --old

Remove kvno associated keys:
  gokutil remove --keytab /etc/krb5.keytab --principal HTTP/MYHOST --kvno 2
`,
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
	removeCmd.Flags().IntVar(&kvno, "kvno", 0, "Removes all entries for the specified principal whose key version number matches kvno.")
	removeCmd.Flags().BoolVar(&all, "all", false, "Removes all entries for the specified principal.")
	removeCmd.Flags().BoolVar(&old, "old", false, "Removes all entries for the specified principal, except those principals with the highest key version number.")
}

func MyCmd() {
	fmt.Println("ktremove called")
}
