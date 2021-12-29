package docs

type OidcCert struct {
	Keys [] OidcCertKey `json:"keys"`
}

type OidcCertKey struct {
	Kty string `json:"kty"`
	N string `json:"n"`
	E string `json:"e"`
	X5c string `json:"x5c"`
}
