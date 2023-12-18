package jwt_test

import (
	"testing"
	"time"

	"github.com/BoyYangZai/go-server-lib/pkg/jwt"
)

func TestGenerateTokenAndParseJwt(t *testing.T) {
	username := "testuser"
	// 生成 Token
	token, err := jwt.GenerateToken(username)
	if err != nil {
		t.Errorf("GenerateToken failed: %v", err)
		return
	}
	print(token, 111)
	// 解析 Token
	claims, err := jwt.ParseJwt(token)
	if err != nil {
		t.Errorf("ParseJwt failed: %v", err)
		return
	}

	// 验证 Token 中的用户名是否正确
	if claims.Username != username {
		t.Errorf("Username mismatch. Expected: %s, Got: %s", username, claims.Username)
		return
	}

	// 验证 Token 是否在有效期内
	now := time.Now()
	if now.After(claims.ExpiresAt.Time) || now.Before(claims.IssuedAt.Time) || now.Before(claims.NotBefore.Time) {
		t.Errorf("Token is not within valid time range")
		return
	}
}

func TestInvalidToken(t *testing.T) {
	invalidToken := "invalid_token"

	// 尝试解析无效的 Token
	_, err := jwt.ParseJwt(invalidToken)
	if err == nil {
		t.Errorf("Expected error for invalid token, but got nil")
		return
	}
}
