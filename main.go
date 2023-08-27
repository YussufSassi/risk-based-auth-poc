package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/yussufsassi/risk-based-auth-poc/cache"
	"github.com/yussufsassi/risk-based-auth-poc/captcha"
)

func getRequiredFormValues(c *gin.Context, keys ...string) (map[string]string, bool) {
	formValues := make(map[string]string)
	for _, key := range keys {
		value, exists := c.GetPostForm(key)
		if !exists {
			return nil, false
		}
		formValues[key] = value
	}
	return formValues, true
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.Use(static.Serve("/", static.LocalFile("./dist", true)))
	r.POST("/authenticate", authVerify)
	r.POST("/captcha", captchaVerify)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func authVerify(c *gin.Context) {
	requiredKeys := []string{"user", "password", "token"}

	formValues, valid := getRequiredFormValues(c, requiredKeys...)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	user := formValues["user"]
	password := formValues["password"]
	token := formValues["token"]
	if user == "user" && password == "password" {
		tokenExists := cache.TokenExists(user, token)
		if !tokenExists {
			fmt.Println("Not exist")
			cache.SaveToken(user, token)
			c.JSON(http.StatusOK, gin.H{
				"user":    user,
				"token":   token,
				"message": fmt.Sprintf("Successfully authenticated with fingerprint: %s!", token),
			})
			return
		}
		isValidToken := cache.CmpToken(user, token)
		fmt.Println(isValidToken)
		if isValidToken {
			cache.SaveToken(user, token)
			c.JSON(http.StatusOK, gin.H{
				"user":    user,
				"token":   token,
				"message": fmt.Sprintf("Successfully authenticated with fingerprint: %s!", token),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"user":    user,
			"token":   token,
			"message": fmt.Sprintf("Fingerprint mismatch. Your fingerprint (%s), does not match our records. Please validate the captcha and try again.", token),
			"captcha": true,
		})
		return

	}

	c.JSON(http.StatusBadRequest, gin.H{
		"user":    user,
		"token":   token,
		"message": "invalid username or password.",
	})

}

func captchaVerify(c *gin.Context) {
	requiredKeys := []string{"user", "token", "code"}
	formValues, valid := getRequiredFormValues(c, requiredKeys...)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}
	user := formValues["user"]
	token := formValues["token"]
	code := formValues["code"]

	verfiyCatpcha := captcha.VerifyCaptcha(user, token, code)

	if verfiyCatpcha {
		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully verified captcha. Please try to sign in again",
			"user":    user,
			"token":   token,
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Invalid captcha! Try again!",
		"user":    user,
		"token":   token,
	})

}
