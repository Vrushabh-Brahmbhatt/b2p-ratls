package ratls

import "time"

const (
	EKMLabel              = "EXPORTER-ratls/v1"
	EKMLength             = 32
	ConnectPath           = "/ratls/connect"
	// MetadataTokenEndpoint is served by the CS launcher via Unix socket.
	// The socket is mounted at /run/container_launcher/teeserver.sock inside the container.
	MetadataTokenEndpoint    = "http://localhost/v1/token"
	MetadataTokenSocketPath  = "/run/container_launcher/teeserver.sock"
	GCPCSIssuer           = "https://confidentialcomputing.googleapis.com"
	WellKnownPath         = "/.well-known/openid-configuration"
	TDXHWModelPrefix      = "GCP_INTEL_TDX"
	SEVHWModelPrefix      = "GCP_AMD_SEV"
	CSSwName              = "CONFIDENTIAL_SPACE"
)

// AttestationBundle is what GPU CS returns to Buffer TEE over /ratls/connect.
type AttestationBundle struct {
	OIDCToken string `json:"oidc_token"`
}

// VerificationOptions holds pinned values Buffer TEE uses during verification.
type VerificationOptions struct {
	Audience            string
	ExpectedImageDigest string
	TokenLeeway         time.Duration
}

// TokenRequestBody is the JSON body for POST to MetadataTokenEndpoint.
type TokenRequestBody struct {
	Audience  string   `json:"audience"`
	Nonces    []string `json:"nonces"`
	TokenType string   `json:"token_type"`
}
