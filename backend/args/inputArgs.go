package args

var (
	// Args is a reference to an instantiation of
	// the configuration that the CLI expects but
	// with some values set.
	// By setting some values in advance we provide
	// default values that the user might provide
	// or not.
	Args = &CliArgs{
		Listen:           8080,
		Harbor:           "https://harbor.example.com/api/",
		User:             "admin",
		JwtKey:           "Not-@-S3cretKey",
		LdapSSL:          true,
		LdapInsecureSkip: false,
		LdapEnable:       false,
		AdminUsername:    "admin",
		AdminPassword:    "Not@SecurePassword",
		GoogleAccessID:   "None",
		PrivateKey:       "None",
	}
)

// CliArgs defines the configuration that the CLI
// expects. By using a struct we can very easily
// aggregate them into an object and check what are
// the expected types.
// If we need to mock this later it's just a matter
// of reusing the struct.
type CliArgs struct {
	Listen           int    `arg:"-l,help:port to listen to,env"`
	Harbor           string `arg:"-u,help:Url to harbor api endpoint,env"`
	User             string `arg:"-n,help:User to login on harbor,env"`
	Password         string `arg:"-p,help:Password to connect on harbor,env"`
	Etcd             string `arg:"-d,help:url to etcd <host>:<port>,env"`
	JwtKey           string `arg:"-j,help:secret key for token generation,env"`
	LdapBase         string `arg:"help:ldapdabse,env"`
	LdapPort         int    `arg:"help:ldap port,env"`
	LdapHost         string `arg:"help:ldap host,env"`
	LdapDN           string `arg:"help:ldap dn of user allowed to query the ldap base,env"`
	LdapBindPassword string `arg:"help:ldap password of dn user,env"`
	LdapSSL          bool   `arg:"help:ldap enable ssl,env"`
	LdapInsecureSkip bool   `arg:"help:ldap skip certificate verification,env"`
	LdapEnable       bool   `arg:"help:enable the ldap authentication,env"`
	AdminUsername    string `arg:"help:Admin user in case of no ldap enable,env"`
	AdminPassword    string `arg:"help:Admin password in case of no ldap enable,env"`
	Regexp           string `arg:"help:Regular expression to filter images from charts. for example exclude postgres-exporter from postgres,env"`
	S3Bucket         string `arg:"help:name of the bucket if none then the convenience script will be downloaded,env"`
	S3GCSKey         string `arg:"help:Service account key for GCS,env"`
	JenkinsBuildAPI  string `arg:"help:path to trigger the build,env"`
	JenkinsURL       string `arg:"help:URL of Jenkins,env"`
	JenkinsUser      string `arg:"help:User of Jenkins (should be a token),env"`
	JenkinsPass      string `arg:"help:Password for Jenkins (should be a token),env"`
	Admins           string `arg:"help:comma separated list of administrators,env"`
	GoogleAccessID   string `arg:"help:access id of GCS if s3 enabled,env"`
	PrivateKey       string `arg:"help:private key of GCS if s3 enabled,env"`
	GoogleClientID   string `arg:"help:google client id of GCS if google auth enable,env"`
	GoogleSecret     string `arg:"help:google secret of GCS if google auth enable,env"`
}
