package header

// HTTP/HTTPS security headers.
const (
	StrictTransportSecurity = "Strict-Transport-Security"  // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security
	ContentSecurityPolicy   = "Content-Security-Policy"    // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy
	CrossOriginOpenerPolicy = "Cross-Origin-Opener-Policy" // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cross-Origin-Opener-Policy
	ReferrerPolicy          = "Referrer-Policy"            // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy
	ContentTypeOptions      = "X-Content-Type-Options"     // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options
	XSSProtection           = "X-XSS-Protection"           // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection
	FrameOptions            = "X-Frame-Options"            // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options
	ForwardedProto          = "X-Forwarded-Proto"          // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-Proto
)

// Standard security policies.
const (
	PolicyDeny               = "DENY"
	PolicyNoSniff            = "nosniff"
	PolicyBlockXSS           = "1; mode=block"
	PolicySameOrigin         = "same-origin"
	PolicyFrameAncestorsNone = "frame-ancestors 'none';"
)

// Security header default policies.
var (
	DefaultContentSecurityPolicy = PolicyFrameAncestorsNone
	DefaultFrameOptions          = PolicyDeny
)