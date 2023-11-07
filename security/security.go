package security

type Security struct {
	APIKey string
}

func NewSecurity(apiKey string) Security {
	return Security{
		APIKey: apiKey,
	}
}
