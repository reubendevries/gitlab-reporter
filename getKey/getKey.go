package getKey

type Environment struct {

	// The address of this service
	MyAddress string `env:"MY_ADDRESS"                    default:":8080"                        description:"Listen to http traffic on this tcp address"             long:"my-address"`

	// Vault address, approle login credentials, and secret locations
	VaultAddress             string `env:"VAULT_ADDRESS"                 default:"localhost:8200"               description:"Vault address"                                          long:"vault-address"`
	VaultApproleRoleID       string `env:"VAULT_APPROLE_ROLE_ID"         required:"true"                        description:"AppRole RoleID to log in to Vault"                      long:"vault-approle-role-id"`
	VaultApproleSecretIDFile string `env:"VAULT_APPROLE_SECRET_ID_FILE"  default:"/tmp/secret"                  description:"AppRole SecretID file path to log in to Vault"          long:"vault-approle-secret-id-file"`
	VaultAPIKeyPath          string `env:"VAULT_API_KEY_PATH"            default:"kv-v2/data/api-key"           description:"Path to the API key used by 'secure-sevice'"            long:"vault-api-key-path"`
	VaultAPIKeyField         string `env:"VAULT_API_KEY_FIELD"           default:"api-key-field"                description:"The secret field name for the API key"                  long:"vault-api-key-descriptor"`
	VaultDatabaseCredsPath   string `env:"VAULT_DATABASE_CREDS_PATH"     default:"database/creds/dev-readonly"  description:"Temporary database credentials will be generated here"  long:"vault-database-creds-path"`

	// A service which requires a specific secret API key (stored in Vault)
	SecureServiceAddress string `env:"SECURE_SERVICE_ADDRESS"        required:"true"                        description:"3rd party service that requires secure credentials"     long:"secure-service-address"`
}

func GetKey() {

}
