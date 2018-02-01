// Code generated by schema-generate. DO NOT EDIT.

package schema

// ContainerOverrides
type ContainerOverrides struct {
	Limits   *ResourceOverride `json:"limits,omitempty"`
	Requests *ResourceOverride `json:"requests,omitempty"`
}

// DatacenterOpsConfiguration Configuration for Sourcegraph Datacenter Ops. Defines all config parameters that can be used to change the operational parameters of Sourcegraph Datacenter (e.g., for scaling).
type DatacenterOpsConfiguration struct {
	AlertmanagerConfig      string                          `json:"alertmanagerConfig,omitempty"`
	AlertmanagerURL         string                          `json:"alertmanagerURL,omitempty"`
	AuthProxyIP             string                          `json:"authProxyIP,omitempty"`
	AuthProxyPassword       string                          `json:"authProxyPassword,omitempty"`
	DeploymentOverrides     map[string]*DeploymentOverrides `json:"deploymentOverrides,omitempty"`
	ExperimentIndexedSearch bool                            `json:"experimentIndexedSearch,omitempty"`
	GitoliteIP              string                          `json:"gitoliteIP,omitempty"`
	GitserverCount          int                             `json:"gitserverCount,omitempty"`
	GitserverDiskSize       string                          `json:"gitserverDiskSize,omitempty"`
	GitserverSSH            map[string]string               `json:"gitserverSSH,omitempty"`
	HttpNodePort            int                             `json:"httpNodePort,omitempty"`
	HttpsNodePort           int                             `json:"httpsNodePort,omitempty"`
	LangGo                  bool                            `json:"langGo,omitempty"`
	LangJava                bool                            `json:"langJava,omitempty"`
	LangJavaScript          bool                            `json:"langJavaScript,omitempty"`
	LangPHP                 bool                            `json:"langPHP,omitempty"`
	LangPython              bool                            `json:"langPython,omitempty"`
	LangSwift               bool                            `json:"langSwift,omitempty"`
	LangTypeScript          bool                            `json:"langTypeScript,omitempty"`
	NodeSSDPath             string                          `json:"nodeSSDPath,omitempty"`
	PhabricatorIP           string                          `json:"phabricatorIP,omitempty"`
	StorageClass            string                          `json:"storageClass,omitempty"`
	UseRBAC                 bool                            `json:"useRBAC,omitempty"`
}

// DeploymentOverrides
type DeploymentOverrides struct {
	Containers map[string]*ContainerOverrides `json:"containers,omitempty"`
	Replicas   *int                           `json:"replicas,omitempty"`
}

// GitHubConnection
type GitHubConnection struct {
	Certificate                 string   `json:"certificate,omitempty"`
	InitialRepositoryEnablement bool     `json:"initialRepositoryEnablement,omitempty"`
	PreemptivelyClone           bool     `json:"preemptivelyClone,omitempty"`
	Repos                       []string `json:"repos,omitempty"`
	RepositoryPathPattern       string   `json:"repositoryPathPattern,omitempty"`
	RepositoryQuery             []string `json:"repositoryQuery,omitempty"`
	Token                       string   `json:"token"`
	Url                         string   `json:"url,omitempty"`
}

// GitLabConnection
type GitLabConnection struct {
	Certificate                 string   `json:"certificate,omitempty"`
	InitialRepositoryEnablement bool     `json:"initialRepositoryEnablement,omitempty"`
	ProjectQuery                []string `json:"projectQuery,omitempty"`
	RepositoryPathPattern       string   `json:"repositoryPathPattern,omitempty"`
	Token                       string   `json:"token"`
	Url                         string   `json:"url"`
}

// Langservers
type Langservers struct {
	Address  string `json:"address,omitempty"`
	Language string `json:"language,omitempty"`
}

// Links
type Links struct {
	Blob       string `json:"blob,omitempty"`
	Commit     string `json:"commit,omitempty"`
	Repository string `json:"repository,omitempty"`
	Tree       string `json:"tree,omitempty"`
}

// OpenIDConnectAuthProvider Configures the OpenID Connect authentication provider for SSO.
type OpenIDConnectAuthProvider struct {
	ClientID           string `json:"clientID"`
	ClientSecret       string `json:"clientSecret"`
	Issuer             string `json:"issuer"`
	OverrideToken      string `json:"overrideToken,omitempty"`
	RequireEmailDomain string `json:"requireEmailDomain,omitempty"`
}

// OpenSearchSettings
type OpenSearchSettings struct {
	SearchUrl string `json:"searchUrl"`
}

// Phabricator
type Phabricator struct {
	Repos []Repos `json:"repos,omitempty"`
	Token string  `json:"token,omitempty"`
	Url   string  `json:"url,omitempty"`
}

// Repos
type Repos struct {
	Callsign string `json:"callsign"`
	Path     string `json:"path"`
}

// Repository
type Repository struct {
	Links *Links `json:"links,omitempty"`
	Path  string `json:"path"`
	Type  string `json:"type,omitempty"`
	Url   string `json:"url"`
}

// ResourceOverride
type ResourceOverride struct {
	Cpu    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
}

// SAMLAuthProvider Configures the SAML authentication provider for SSO.
type SAMLAuthProvider struct {
	IdentityProviderMetadataURL string `json:"identityProviderMetadataURL"`
	ServiceProviderCertificate  string `json:"serviceProviderCertificate"`
	ServiceProviderPrivateKey   string `json:"serviceProviderPrivateKey"`
}

// SMTPServerConfig The SMTP server used to send transactional emails (such as email verifications, reset-password emails, and notifications).
type SMTPServerConfig struct {
	Authentication string `json:"authentication"`
	Domain         string `json:"domain,omitempty"`
	Host           string `json:"host"`
	Password       string `json:"password,omitempty"`
	Port           int    `json:"port"`
	Username       string `json:"username,omitempty"`
}

// SearchSavedQueries
type SearchSavedQueries struct {
	Description         string   `json:"description"`
	Key                 string   `json:"key"`
	Notify              bool     `json:"notify,omitempty"`
	NotifyOrganizations []string `json:"notifyOrganizations,omitempty"`
	NotifySlack         bool     `json:"notifySlack,omitempty"`
	NotifyUsers         []string `json:"notifyUsers,omitempty"`
	Query               string   `json:"query"`
	ShowOnHomepage      bool     `json:"showOnHomepage,omitempty"`
}

// SearchScope
type SearchScope struct {
	Description string `json:"description,omitempty"`
	Id          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Value       string `json:"value"`
}

// Settings Configuration settings for users and organizations on Sourcegraph Server.
type Settings struct {
	SearchRepositoryGroups map[string][]string  `json:"search.repositoryGroups,omitempty"`
	SearchSavedQueries     []SearchSavedQueries `json:"search.savedQueries,omitempty"`
	SearchScopes           []SearchScope        `json:"search.scopes,omitempty"`
}

// SiteConfiguration Configuration for a Sourcegraph Server site.
type SiteConfiguration struct {
	AdminUsernames                 string                     `json:"adminUsernames,omitempty"`
	AppURL                         string                     `json:"appURL,omitempty"`
	AuthAllowSignup                bool                       `json:"auth.allowSignup,omitempty"`
	AuthOpenIDConnect              *OpenIDConnectAuthProvider `json:"auth.openIDConnect,omitempty"`
	AuthProvider                   string                     `json:"auth.provider,omitempty"`
	AuthSaml                       *SAMLAuthProvider          `json:"auth.saml,omitempty"`
	AuthUserIdentityHTTPHeader     string                     `json:"auth.userIdentityHTTPHeader,omitempty"`
	AuthUserOrgMap                 map[string][]string        `json:"auth.userOrgMap,omitempty"`
	AutoRepoAdd                    bool                       `json:"autoRepoAdd,omitempty"`
	CorsOrigin                     string                     `json:"corsOrigin,omitempty"`
	DisableAutoGitUpdates          bool                       `json:"disableAutoGitUpdates,omitempty"`
	DisablePublicRepoRedirects     bool                       `json:"disablePublicRepoRedirects,omitempty"`
	DisableTelemetry               bool                       `json:"disableTelemetry,omitempty"`
	EmailAddress                   string                     `json:"email.address,omitempty"`
	EmailSmtp                      *SMTPServerConfig          `json:"email.smtp,omitempty"`
	ExecuteGradleOriginalRootPaths string                     `json:"executeGradleOriginalRootPaths,omitempty"`
	GitMaxConcurrentClones         int                        `json:"gitMaxConcurrentClones,omitempty"`
	GitOriginMap                   string                     `json:"gitOriginMap,omitempty"`
	Github                         []GitHubConnection         `json:"github,omitempty"`
	GithubClientID                 string                     `json:"githubClientID,omitempty"`
	GithubClientSecret             string                     `json:"githubClientSecret,omitempty"`
	GithubEnterpriseAccessToken    string                     `json:"githubEnterpriseAccessToken,omitempty"`
	GithubEnterpriseCert           string                     `json:"githubEnterpriseCert,omitempty"`
	GithubEnterpriseURL            string                     `json:"githubEnterpriseURL,omitempty"`
	GithubPersonalAccessToken      string                     `json:"githubPersonalAccessToken,omitempty"`
	Gitlab                         []GitLabConnection         `json:"gitlab,omitempty"`
	GitoliteHosts                  string                     `json:"gitoliteHosts,omitempty"`
	GitoliteRepoBlacklist          string                     `json:"gitoliteRepoBlacklist,omitempty"`
	HtmlBodyBottom                 string                     `json:"htmlBodyBottom,omitempty"`
	HtmlBodyTop                    string                     `json:"htmlBodyTop,omitempty"`
	HtmlHeadBottom                 string                     `json:"htmlHeadBottom,omitempty"`
	HtmlHeadTop                    string                     `json:"htmlHeadTop,omitempty"`
	HttpToHttpsRedirect            bool                       `json:"httpToHttpsRedirect,omitempty"`
	Langservers                    []Langservers              `json:"langservers,omitempty"`
	LicenseKey                     string                     `json:"licenseKey,omitempty"`
	LightstepAccessToken           string                     `json:"lightstepAccessToken,omitempty"`
	LightstepProject               string                     `json:"lightstepProject,omitempty"`
	MaxReposToSearch               int                        `json:"maxReposToSearch,omitempty"`
	NoGoGetDomains                 string                     `json:"noGoGetDomains,omitempty"`
	OidcClientID                   string                     `json:"oidcClientID,omitempty"`
	OidcClientSecret               string                     `json:"oidcClientSecret,omitempty"`
	OidcEmailDomain                string                     `json:"oidcEmailDomain,omitempty"`
	OidcOverrideToken              string                     `json:"oidcOverrideToken,omitempty"`
	OidcProvider                   string                     `json:"oidcProvider,omitempty"`
	OpenSearch                     *OpenSearchSettings        `json:"openSearch,omitempty"`
	Phabricator                    []Phabricator              `json:"phabricator,omitempty"`
	PhabricatorURL                 string                     `json:"phabricatorURL,omitempty"`
	PrivateArtifactRepoID          string                     `json:"privateArtifactRepoID,omitempty"`
	PrivateArtifactRepoPassword    string                     `json:"privateArtifactRepoPassword,omitempty"`
	PrivateArtifactRepoURL         string                     `json:"privateArtifactRepoURL,omitempty"`
	PrivateArtifactRepoUsername    string                     `json:"privateArtifactRepoUsername,omitempty"`
	RepoListUpdateInterval         int                        `json:"repoListUpdateInterval,omitempty"`
	ReposList                      []Repository               `json:"repos.list,omitempty"`
	SamlIDProviderMetadataURL      string                     `json:"samlIDProviderMetadataURL,omitempty"`
	SamlSPCert                     string                     `json:"samlSPCert,omitempty"`
	SamlSPKey                      string                     `json:"samlSPKey,omitempty"`
	SearchScopes                   []SearchScope              `json:"searchScopes,omitempty"`
	SecretKey                      string                     `json:"secretKey,omitempty"`
	Settings                       *Settings                  `json:"settings,omitempty"`
	SiteID                         string                     `json:"siteID,omitempty"`
	SsoUserHeader                  string                     `json:"ssoUserHeader,omitempty"`
	TlsCert                        string                     `json:"tlsCert,omitempty"`
	TlsKey                         string                     `json:"tlsKey,omitempty"`
	UpdateChannel                  *string                    `json:"update.channel,omitempty"`
	UseJaeger                      bool                       `json:"useJaeger,omitempty"`
}
