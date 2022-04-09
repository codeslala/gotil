package oauth2

const (
	// auth address is hardcoded.
	authAddress      = "172.19.253.65:80"
	localAuthAddress = "localhost:8025"

	// tokenURl is hardcoded.
	tokenURL      = "https://oauth.xxx.com/v1/oauth2/clientauth"
	localTokenURL = "http://localhost:8023/v1/oauth2/clientauth"
)

var useLocal bool

func UseLocal() {
	useLocal = true
}

func AuthAddress() string {
	if useLocal {
		return localAuthAddress
	}
	return authAddress
}

func TokenURL() string {
	if useLocal {
		return localTokenURL
	}
	return tokenURL
}
