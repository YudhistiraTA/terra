package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/YudhistiraTA/terra/internal/interface/api/rest/errors"
	"github.com/gin-gonic/gin"
)

func verifySignature(c *gin.Context) {
	signature := c.GetHeader("X-SIGNATURE")
	timestamp := c.GetHeader("X-TIMESTAMP")
	if signature == "" || timestamp == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.ErrorResponse{Message: "Invalid header"})
		return
	}
	endpoint := c.Request.URL.String()
	method := c.Request.Method
	payload := fmt.Sprintf("%s:%s:%s", method, endpoint, timestamp)
	token := strings.Split(c.GetHeader("Authorization"), "Bearer ")
	if len(token) == 2 {
		payload = fmt.Sprintf("%s:%s", payload, token[1])
	}
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}
	if len(bodyBytes) > 0 {
		bodyHash := sha256.Sum256(bodyBytes)
		bodyHex := hex.EncodeToString(bodyHash[:])
		payload = fmt.Sprintf("%s:%s", payload, bodyHex)
	}
	fmt.Println(payload)

	hmacSecret := os.Getenv("HMAC_SECRET")
	if hmacSecret == "" {
		panic("HMAC_SECRET environment variable is not set")
	}
	mac := hmac.New(sha512.New, []byte(hmacSecret))
	mac.Write([]byte(payload))
	expectedMAC := mac.Sum(nil)
	expectedSignature := base64.StdEncoding.EncodeToString(expectedMAC)

	if !hmac.Equal([]byte(signature), []byte(expectedSignature)) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.ErrorResponse{Message: "Invalid signature"})
		return
	}
	c.Next()
}

func VerifySignature() gin.HandlerFunc {
	return verifySignature
}
