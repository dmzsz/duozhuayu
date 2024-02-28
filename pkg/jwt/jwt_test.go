package jwt_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/dmzsz/duozhuayu/internal/configs"
	"github.com/dmzsz/duozhuayu/internal/constants"
	V1Domains "github.com/dmzsz/duozhuayu/internal/domains/v1"
	"github.com/dmzsz/duozhuayu/pkg/jwt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// 获取当前测试文件所在的目录
	testDir, err := os.Getwd()
	if err != nil {
		panic("Failed to get current working directory: " + err.Error())
	}
	// 获取项目根目录
	projectDir := filepath.Join(testDir, "..", "..", "..", "config") // 可能需要根据实际情况调整路径层级

	viper.SetConfigName(".env") // allow directly reading from .env file
	viper.SetConfigType("env")
	viper.AddConfigPath(projectDir)
	viper.AddConfigPath("internal/configs")
	viper.AddConfigPath("/")
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		panic("failed to parse env to config struct")
	}

	err = viper.Unmarshal(&configs.AppConfig)
	if err != nil {
		panic("failed to parse env to config struct")
	}

	// 执行测试
	exitCode := m.Run()

	// 退出测试
	os.Exit(exitCode)
}

func TestGenerateToken(t *testing.T) {
	fmt.Println("configs.AppConfig.JWTSecret", configs.AppConfig)
	jwt.NewJWT()
	token, err := jwt.GenerateToken("asf-asf-asfdasd-asdfsa", "john", "john.doe@example.com", []V1Domains.RoleDomain{{Id: constants.UserRoleID, Name: "user"}}, jwt.AccessToken)
	fmt.Println(token)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestParseToken(t *testing.T) {
	t.Run("With Valid Token", func(t *testing.T) {
		jwt.NewJWT()
		// configs.AppConfig.SecurityConfig.JWT.AccessKeyTTL = 5

		token, _ := jwt.GenerateToken("asf-asf-asfdasd-asdfsa", "john", "john.doe@example.com", []V1Domains.RoleDomain{{Id: constants.AdminRoleID, Name: "admin"}, {Id: constants.UserRoleID, Name: "user"}}, jwt.AccessToken)
		fmt.Println(token)
		claims, err := jwt.ParseToken(token)
		assert.NoError(t, err)
		assert.Equal(t, "asf-asf-asfdasd-asdfsa", claims.UserId)
		assert.Equal(t, "john", claims.Username)
		assert.Equal(t, "john.doe@example.com", claims.Email)
		fmt.Println(claims.RoleIds)
		assert.Equal(t, []string{constants.AdminRoleID, constants.UserRoleID}, claims.RoleIds)
		assert.True(t, !claims.ExpiresAt.Before(time.Now()))
		assert.Equal(t, configs.AppConfig.SecurityConfig.JWT.Issuer, claims.Issuer)
		assert.True(t, !claims.IssuedAt.After(time.Now()))
	})
}
