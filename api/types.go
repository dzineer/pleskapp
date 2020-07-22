// Copyright 1999-2020. Plesk International GmbH. All rights reserved.

package api

import "git.plesk.ru/~abashurov/pleskapp/types"

type DomainManagement interface {
	CreateDomain(domain string, ipAddresses types.ServerIPAddresses) (*DomainInfo, error)
	AddDomainFeatures(domain string, features []string) error
	RemoveDomain(domain string) error
	GetDomain(domain string) (DomainInfo, error)
	ListDomains() ([]DomainInfo, error)
}

type DomainInfo struct {
	ID             int
	Name           string
	HostingType    string
	ParentDomainID int
	GUID           string
	WWWRoot        string
	Sysuser        string
}

type FTPManagement interface {
	ListDomainFtpUsers(domain string, user types.FtpUser) ([]FTPUserInfo, error)
	CreateFtpUser(domain string, user types.FtpUser) (*FTPUserInfo, error)
	UpdateFtpUser(domain string, user string, userNew types.FtpUser) error
	DeleteFtpUser(domain string, user types.FtpUser) error
}

type FTPUserInfo struct {
	Name           string
	Home           string
	Quota          int
	ParentDomainID int
}

type DatabaseManagement interface {
	ListDatabases() ([]DatabaseInfo, error)
	ListDomainDatabases(domain string) ([]DatabaseInfo, error)
	ListDatabaseServers() ([]DatabaseServerInfo, error)
	CreateDatabase(domain types.Domain, database types.NewDatabase, server types.DatabaseServer) (*DatabaseInfo, error)
	RemoveDatabase(database types.Database) error
	DeployDatabase(database types.Database, dbuser types.DatabaseUser, server types.DatabaseServer, filename string, isWindows bool, sysuser *string) error
}

type DatabaseInfo struct {
	ID               int
	Name             string
	Type             string
	ParentDomainID   int
	DatabaseServerID int
}

type DatabaseServerInfo struct {
	ID        int
	Host      string
	Port      int
	Type      string
	Status    string
	IsDefault bool
	IsLocal   bool
}

type DatabaseUserManagement interface {
	CreateDatabaseUser(database types.Database, dbuser types.NewDatabaseUser) (*DatabaseUserInfo, error)
	RemoveDatabaseUser(dbuser types.DatabaseUser) error
	ListDatabaseUsers(database types.Database) ([]DatabaseUserInfo, error)
}

type DatabaseUserInfo struct {
	ID         int
	Login      string
	DatabaseID int
}

type Authentication interface {
	GetAPIKey(preAuth PreAuth) (string, error)
	GetLoginLink(auth Auth) (string, error)
	RemoveAPIKey(auth Auth) (string, error)
}

type Server interface {
	GetInfo() (ServerInfo, error)
	GetIpAddresses() (types.ServerIPAddresses, error)
}

type ServerInfo struct {
	IsWindows bool
	Version   string
}

type AuthClient interface {
	GetIgnoreSsl() bool
}

type Auth interface {
	GetAddress() string
	GetPort() string
	GetIgnoreSsl() bool
	GetIsWindows() bool
	GetApiKey() string
}

type PreAuth interface {
	GetAddress() string
	GetPort() string
	GetIgnoreSsl() bool
	GetIsWindows() bool
	GetLogin() string
	GetPassword() string
}

type HasApiUrlComponents interface {
	GetAddress() string
	GetPort() string
}

func GetApiUrl(a HasApiUrlComponents, path string) string {
	return "https://" + a.GetAddress() + ":" + a.GetPort() + path
}