package sdk

// Boolean is a boolean
var Boolean = JSONSchema{
	Type:   "boolean",
	GoType: "bool",
}

// Object is an object
var Object = JSONSchema{
	Type:   "object",
	GoType: "interface{}",
}

// Integer are integers
var Integer = JSONSchema{
	Type:   "integer",
	GoType: "int",
}

// Bytes are byte strings
var Bytes = JSONSchema{
	Type:        "string",
	GoType:      "[]byte",
	Format:      "bytes",
	DisplayType: "bytes",
}

// Date is a date
var Date = JSONSchema{
	Type:        "string",
	GoType:      "time.Time",
	Format:      "date-time",
	DisplayType: "date",
}

// Float is a floating point number
var Float = JSONSchema{
	Type:   "number",
	GoType: "float64",
}

// Python is python code
var Python = JSONSchema{
	Type:        "string",
	GoType:      "string",
	Format:      "python",
	DisplayType: "python",
}

// Password is a password type
var Password = JSONSchema{
	Type:        "string",
	GoType:      "string",
	Format:      "password",
	DisplayType: "password",
}

// String is a string
var String = JSONSchema{
	Type:   "string",
	GoType: "string",
}

// Array is generic array
var Array = JSONSchema{
	Type:   "array",
	GoType: "[]interface{}",
}

// File is a file
var File = JSONSchema{
	ID:          "file",
	Title:       "File",
	Description: "File Object",
	Type:        "object",
	GoType:      "types.SDKFile",
	BackRef:     true,
	Properties: map[string]JSONSchema{
		"filename": JSONSchema{
			Type:        "string",
			GoType:      "string",
			Title:       "Filename",
			Description: "Name of file",
		},
		"content": JSONSchema{
			Type:        "string",
			GoType:      "string",
			Format:      "bytes",
			Title:       "Content",
			Description: "File contents",
		},
	},
}

// CredentialUsernamePassword is a CC credential type for managing login names and passwords.
var CredentialUsernamePassword = JSONSchema{
	ID:          "credential_username_password",
	Title:       "Credential: Username and Password",
	Description: "A username and password combination",
	Type:        "object",
	GoType:      "types.CredentialUsernamePassword",
	Required:    []string{"username", "password"},
	BackRef:     true,
	Properties: map[string]JSONSchema{
		"username": JSONSchema{
			Type:        "string",
			GoType:      "string",
			Title:       "Username",
			Description: "The username to log in with",
		},
		"password": JSONSchema{
			Title:       "Password",
			Description: "The password",
			Type:        "string",
			GoType:      "string",
			Format:      "password",
			DisplayType: "password",
		},
	},
}

// CredentialAsymmetricKey is a CC credential type for managing access tokens
var CredentialAsymmetricKey = JSONSchema{
	ID:          "credential_asymmetric_key",
	Title:       "Credential: Asymmetric Key",
	Description: "A shared key",
	Type:        "object",
	GoType:      "types.CredentialAsymmetricKey",
	Required:    []string{"privateKey"},
	BackRef:     true,
	Properties: map[string]JSONSchema{
		"privateKey": JSONSchema{
			Title:       "Private Key",
			Description: "The private key",
			Type:        "string",
			GoType:      "string",
			Format:      "password",
			DisplayType: "password",
		},
	},
}

// CredentialToken is a CC credential type for managing a token and an an optional domain
var CredentialToken = JSONSchema{
	ID:          "credential_token",
	Title:       "Credential: Token",
	Description: "A pair of a token, and an optional domain",
	Type:        "object",
	GoType:      "types.CredentialToken",
	Required:    []string{"token"},
	BackRef:     true,
	Properties: map[string]JSONSchema{
		"token": JSONSchema{
			Title:       "Token",
			Description: "The shared token",
			Type:        "string",
			GoType:      "string",
			Format:      "password",
			DisplayType: "password",
		},
		"domain": JSONSchema{
			Type:        "string",
			GoType:      "string",
			Title:       "Domain",
			Description: "The domain for the token",
		},
	},
}

// CredentialSecretKey is a CC credential type for managing access tokens
var CredentialSecretKey = JSONSchema{
	ID:          "credential_secret_key",
	Title:       "Credential: Secret Key",
	Description: "A shared secret key",
	Type:        "object",
	GoType:      "types.CredentialSecretKey",
	Required:    []string{"secretKey"},
	BackRef:     true,
	Properties: map[string]JSONSchema{
		"secretKey": JSONSchema{
			Title:       "Secret Key",
			Description: "The shared secret key",
			Type:        "string",
			GoType:      "string",
			Format:      "password",
			DisplayType: "password",
		},
	},
}
