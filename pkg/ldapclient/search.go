package ldapclient

import (
	"crypto/tls"
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

type LDAPClient struct {
	LdapServer         string
	BindDN             string
	BaseDN             string
	ServerName         string
	InsecureSkipVerify bool
	UseSSL             bool
	SkipTLS            bool
	Conn               *ldap.Conn
	ClientCertificates []tls.Certificate
	LdapPassword       string
}

type LDAPUser struct {
	DN        string
	Username  string
	FirstName string
	LastName  string
	Email     string
}

type LDAPGroup struct {
	DN        string
	GroupName string
	MemberUid []string
}

func (lc *LDAPClient) CreateLdapConnection() error {
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

func (lc *LDAPClient) bindAndSearch(filter string, attributes []string) (*ldap.SearchResult, error) {
	lc.Conn.Bind(lc.BindDN, lc.LdapPassword)
	searchReq := ldap.NewSearchRequest(
		lc.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		filter,
		attributes,
		nil,
	)
	result, err := lc.Conn.Search(searchReq)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if len(result.Entries) > 0 {
		return result, nil
	}
	return nil, fmt.Errorf("no entries found!")
}

func (lc *LDAPClient) ListLDAPUsers() ([]LDAPUser, error) {

	ldapFilter := "(objectClass=posixAccount)"
	attributes := []string{"dn", "cn", "givenName", "sn", "mail"}
	var ldapEntries *ldap.SearchResult
	ldapEntries, err := lc.bindAndSearch(ldapFilter, attributes)
	if err != nil {
		return nil, err
	}
	ldapUsers := make([]LDAPUser, 0)
	for i := 0; i < len(ldapEntries.Entries); i++ {
		ldapUsers = append(ldapUsers, LDAPUser{
			DN:        ldapEntries.Entries[i].DN,
			Username:  ldapEntries.Entries[i].GetAttributeValue("cn"),
			FirstName: ldapEntries.Entries[i].GetAttributeValue("givenName"),
			LastName:  ldapEntries.Entries[i].GetAttributeValue("sn"),
			Email:     ldapEntries.Entries[i].GetAttributeValue("mail"),
		})
	}
	return ldapUsers, nil
}

func (lc *LDAPClient) ListLDAPGroups() ([]LDAPGroup, error) {
	ldapFilter := "(objectClass=posixGroup)"
	attributes := []string{"dn", "cn", "memberUid"}
	var ldapEntries *ldap.SearchResult
	ldapEntries, err := lc.bindAndSearch(ldapFilter, attributes)
	if err != nil {
		return nil, err
	}
	ldapGroups := make([]LDAPGroup, 0)
	for i := 0; i < len(ldapEntries.Entries); i++ {
		ldapGroups = append(ldapGroups, LDAPGroup{
			DN:        ldapEntries.Entries[i].DN,
			GroupName: ldapEntries.Entries[i].GetAttributeValue("cn"),
			MemberUid: ldapEntries.Entries[i].GetAttributeValues("memberUid"),
		})
	}
	return ldapGroups, nil
}
