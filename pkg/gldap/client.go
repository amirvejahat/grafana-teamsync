package gldap

import (
	"crypto/tls"
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

type LdapClient struct {
	LdapServer         string
	BindDN             string
	BaseDN             string
	FilterDN           string
	ServerName         string
	InsecureSkipVerify bool
	UseSSL             bool
	SkipTLS            bool
	Conn               *ldap.Conn
	ClientCertificates []tls.Certificate
	LdapPassword       string
	loginUsername      string
	loginPassword      string
}

func (lc *LdapClient) CreateLdapConnection() error {
	if lc.Conn != nil {
		return nil
	}
	var l *ldap.Conn
	var err error
	if !lc.UseSSL {
		l, err = ldap.Dial("tcp", lc.LdapServer)
		if err != nil {
			return err
		}
	} else {
		config := &tls.Config{
			InsecureSkipVerify: lc.InsecureSkipVerify,
			ServerName:         lc.ServerName,
		}

		if lc.ClientCertificates != nil && len(lc.ClientCertificates) > 0 {
			config.Certificates = lc.ClientCertificates
		}
		l, err = ldap.DialTLS("tcp", lc.LdapServer, config)
		if err != nil {
			return err
		}
	}
	lc.Conn = l
	return nil
}

func (lc *LdapClient) BindAndSearch() (*ldap.SearchResult, error) {
	lc.Conn.Bind(lc.BindDN, lc.LdapPassword)
	searchReq := ldap.NewSearchRequest(
		lc.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		lc.FilterDN,
		[]string{},
		nil,
	)
	result, err := lc.Conn.Search(searchReq)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	for _, entry := range result.Entries {
		entry.PrettyPrint(1)
	}
	return nil, nil
}
