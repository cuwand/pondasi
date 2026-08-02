package headerHelper

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/cuwand/pondasi/constant"
	"github.com/cuwand/pondasi/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupGinContext() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/", nil)
	return c, w
}

func TestGetGin(t *testing.T) {
	c, _ := setupGinContext()

	t.Run("should return header value when exists", func(t *testing.T) {
		expectedValue := "test-client-id"
		c.Request.Header.Set(constant.X_CLIENT_ID, expectedValue)

		result := GetGin(c, constant.X_CLIENT_ID)

		assert.Equal(t, expectedValue, result)
	})

	t.Run("should return empty string when header not exists", func(t *testing.T) {
		result := GetGin(c, "X-NON-EXISTENT")

		assert.Equal(t, "", result)
	})
}

func TestGetGinAndValidate(t *testing.T) {
	t.Run("should return header value when exists", func(t *testing.T) {
		c, _ := setupGinContext()
		expectedValue := "test-timestamp"
		c.Request.Header.Set(constant.X_TIMESTAMP, expectedValue)

		result := GetGinAndValidate(c, constant.X_TIMESTAMP)

		assert.Equal(t, expectedValue, result)
	})

	t.Run("should panic when header is empty", func(t *testing.T) {
		c, _ := setupGinContext()

		assert.Panics(t, func() {
			GetGinAndValidate(c, constant.X_SIGNATURE)
		})
	})

	t.Run("should panic with correct message when header missing", func(t *testing.T) {
		c, _ := setupGinContext()

		defer func() {
			if r := recover(); r != nil {
				assert.Contains(t, r, "Header X-CUSTOM-HEADER is required")
			}
		}()

		GetGinAndValidate(c, "X-CUSTOM-HEADER")
	})
}

func TestSetGin(t *testing.T) {
	c, _ := setupGinContext()

	t.Run("should set header value", func(t *testing.T) {
		key := constant.X_CORE_TRACE_ID
		value := "trace-123456"

		SetGin(c, key, value)

		assert.Equal(t, value, c.Request.Header.Get(key))
	})

	t.Run("should overwrite existing header value", func(t *testing.T) {
		key := constant.X_CLIENT_ID
		c.Request.Header.Set(key, "old-value")

		SetGin(c, key, "new-value")

		assert.Equal(t, "new-value", c.Request.Header.Get(key))
	})
}

func TestGet(t *testing.T) {
	t.Run("should return header value when exists", func(t *testing.T) {
		header := http.Header{}
		header.Set("Content-Type", "application/json")

		result := Get(header, "Content-Type")

		assert.Equal(t, "application/json", result)
	})

	t.Run("should return empty string when header not exists", func(t *testing.T) {
		header := http.Header{}

		result := Get(header, "X-Non-Existent")

		assert.Equal(t, "", result)
	})

	t.Run("should be case insensitive", func(t *testing.T) {
		header := http.Header{}
		header.Set("X-Custom-Header", "test-value")

		result := Get(header, "x-custom-header")

		assert.Equal(t, "test-value", result)
	})
}

func TestGetUserAudit(t *testing.T) {
	t.Run("should panic when header is missing", func(t *testing.T) {
		c, _ := setupGinContext()
		os.Setenv("AUDIT_KEY", "1234567890123456")
		defer os.Unsetenv("AUDIT_KEY")

		assert.Panics(t, func() {
			GetUserAudit(c)
		})
	})

	t.Run("should panic when AUDIT_KEY env not set", func(t *testing.T) {
		c, _ := setupGinContext()
		c.Request.Header.Set(constant.X_CORE_UA, "some-encrypted-value")
		os.Unsetenv("AUDIT_KEY")

		assert.Panics(t, func() {
			GetUserAudit(c)
		})
	})
}

func TestGenerateUserAudit(t *testing.T) {
	t.Run("should generate encrypted string for user", func(t *testing.T) {
		os.Setenv("AUDIT_KEY", "1234567890123456")
		defer os.Unsetenv("AUDIT_KEY")

		user := models.UserRequest{
			Identity: "user-123",
			Username: "testuser",
			FullName: "Test User",
		}

		result := GenerateUserAudit(user)

		assert.NotEmpty(t, result)
		assert.NotEqual(t, "user-123", result)
	})

	t.Run("should panic when AUDIT_KEY env not set", func(t *testing.T) {
		os.Unsetenv("AUDIT_KEY")

		user := models.UserRequest{
			Identity: "user-123",
		}

		assert.Panics(t, func() {
			GenerateUserAudit(user)
		})
	})
}

func TestUserAuditRoundTrip(t *testing.T) {
	t.Run("should encrypt and decrypt user audit correctly", func(t *testing.T) {
		os.Setenv("AUDIT_KEY", "1234567890ABCDEF")
		defer os.Unsetenv("AUDIT_KEY")

		originalUser := models.UserRequest{
			Identity: "user-456",
			Username: "roundtripuser",
			FullName: "Round Trip User",
		}

		// Generate encrypted audit
		encrypted := GenerateUserAudit(originalUser)
		//encrypted := "E0D5B57FD690985278765A06B54E6D7FA522CD6A50A98B059495BCED3EF8E688DD9307B05AA33C388BBE8C2CA42719758E9126FC79B56E1409510C2C204A7A613E130E1685864B54D4FFC692E0BA76A80609AB2E7ECB20B2A41D81BA27F7D156"

		fmt.Println(strings.ToUpper(encrypted))

		// Setup gin context with encrypted header
		c, _ := setupGinContext()
		c.Request.Header.Set(constant.X_CORE_UA, encrypted)

		// Decrypt and verify
		decryptedUser := GetUserAudit(c)

		fmt.Println("decryptedUser")
		fmt.Println(decryptedUser)

		assert.Equal(t, originalUser.Identity, decryptedUser.Identity)
		assert.Equal(t, originalUser.Username, decryptedUser.Username)
		assert.Equal(t, originalUser.FullName, decryptedUser.FullName)
	})
}
