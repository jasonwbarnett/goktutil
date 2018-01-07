package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/jasonwbarnett/gokrb5.v555/keytab"
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
	Args: func(cmd *cobra.Command, args []string) error {
		if kvno == 0 && all == false && old == false {
			return errors.New("Must at least one of: --kvno, --all or --all")
		}
		if kvno != 0 && all != false {
			return errors.New("Only pass one of: --kvno or --all")
		}
		if kvno != 0 && old != false {
			return errors.New("Only pass one of: --kvno or --old")
		}
		if all != false && old != false {
			return errors.New("Only pass one of: --all or --old")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		keytabKeyRemoval()
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().StringVar(&principal, "principal", "", "Specifies the principal to be removed from the keytab file.")
	removeCmd.MarkFlagRequired("principal")
	removeCmd.Flags().IntVar(&kvno, "kvno", 0, "Removes all entries for the specified principal whose key version number matches kvno.")
	removeCmd.Flags().BoolVar(&all, "all", false, "Removes all entries for the specified principal.")
	removeCmd.Flags().BoolVar(&old, "old", false, "Removes all entries for the specified principal, except those principals with the highest key version number.")
}

func keytabKeyRemoval() {
	if _, err := os.Stat(keytabPath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Keytab does not exist, %v\n", keytabPath)
		os.Exit(2)
	}
	kt, err := keytab.Load(keytabPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}

	var newKeytab keytab.Keytab
	if kvno != 0 {
		newKeytab = removeKvno(kt)
	}
	if all == true {
		newKeytab = removeAll(kt)
	}
	if old == true {
		newKeytab = removeOld(kt)
	}
	if bytes.Equal(newKeytab.Bytes(), kt.Bytes()) {
		fmt.Fprintf(os.Stdout, "No changes needed to be saved.\n")
	} else {
		fmt.Fprintf(os.Stdout, "Writing changes\n")
		err = ioutil.WriteFile(keytabPath, newKeytab.Bytes(), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(4)
		}
	}
}

func removeOld(kt keytab.Keytab) keytab.Keytab {
	highestKvno := 0
	for _, entry := range kt.Entries {
		if entry.Principal.Name() == principal && entry.KeyVersionNumber() > highestKvno {
			highestKvno = entry.KeyVersionNumber()
		}
	}

	newKeytab := filterKeytabPrincipals(kt, func(ke keytab.Entry) bool {
		if ke.Principal.Name() == principal && ke.KeyVersionNumber() < highestKvno {
			return false
		}
		return true
	})

	return newKeytab
}

func removeAll(kt keytab.Keytab) keytab.Keytab {
	newKeytab := filterKeytabPrincipals(kt, func(ke keytab.Entry) bool {
		if ke.Principal.Name() == principal {
			return false
		}

		return true
	})

	return newKeytab
}

func removeKvno(kt keytab.Keytab) keytab.Keytab {
	newKeytab := filterKeytabPrincipals(kt, func(ke keytab.Entry) bool {
		if ke.Principal.Name() == principal && ke.KeyVersionNumber() == kvno {
			return false
		}

		return true
	})

	return newKeytab
}

func filterKeytabPrincipals(kt keytab.Keytab, filter func(keytab.Entry) bool) keytab.Keytab {
	entries := kt.Entries[:0]
	for _, entry := range kt.Entries {
		if filter(entry) {
			entries = append(entries, entry)
		}
	}

	kt.Entries = entries

	return kt
}
