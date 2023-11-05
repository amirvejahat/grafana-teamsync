package gldap

import (
	"crypto/tls"

	"github.com/go-ldap/ldap/v3"
)

type LdapClient struct {
	ldapServer         string
	BindDN             string
	BaseDN             string
	FilterDN           string
	ServerName         string
	InsecureSkipVerify bool
	UseSSL             bool
	SkipTLS            bool
	Conn               *ldap.Conn
	ClientCertificates []tls.Certificate
	loginUsername      string
	loginPassword      string
}

func (lc *LdapClient) Connect() error {
	if lc.Conn != nil {
		return nil
	}
	var l *ldap.Conn
	var err error
	if !lc.UseSSL {
		l, err = ldap.Dial("tcp", lc.ldapServer)
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
		l, err = ldap.DialTLS("tcp", lc.ldapServer, config)
		if err != nil {
			return err
		}
	}
	lc.Conn = l
	return nil
}
