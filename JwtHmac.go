package jwtconf

// JwtHmac struct
type JwtHmac struct {
	Issuer     string
	SigningKey string
	ExpiryMS   int64
}
