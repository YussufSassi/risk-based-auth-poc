package captcha

import "github.com/yussufsassi/risk-based-auth-poc/cache"

func VerifyCaptcha(user string, token string, code string) bool {
	if code == "example_captcha" {
		cache.SaveToken(user, token)
		return true
	}

	return false
}
