package main

import (
	"fmt"

	"backend/args"

	"github.com/jtblin/go-ldap-client"
)

func initLdapClient() *ldap.LDAPClient {
	return &ldap.LDAPClient{
		Base:               args.Args.LdapBase,
		Host:               args.Args.LdapHost,
		Port:               args.Args.LdapPort,
		UseSSL:             args.Args.LdapSSL,
		InsecureSkipVerify: args.Args.LdapInsecureSkip,
		BindDN:             args.Args.LdapDN,
		BindPassword:       args.Args.LdapBindPassword,
		UserFilter:         "(cn=%s)",
		GroupFilter:        "(memberOf=%s)",
		Attributes:         []string{"givenName", "sn", "mail", "uid"},
	}
}

func authentUser(username string, password string) bool {
	// It is the responsibility of the caller to close the connection
	nm.Log(fmt.Sprintf("dn %s", args.Args.LdapDN))
	if args.Args.LdapEnable {
		defer ldapCli.Close()

		ok, _, err := ldapCli.Authenticate(username, password)
		if err != nil {
			nm.Log(fmt.Sprintf("Error authenticating user %s: %+v", username, err))
			return false
		}
		if !ok {
			nm.Log(fmt.Sprintf("Authenticating failed for user %s", username))
			return false
		}
		return ok
	}

	return username == args.Args.AdminUsername && password == args.Args.AdminPassword

}
