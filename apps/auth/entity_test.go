package auth

import (
	"github.com/stretchr/testify/require"
	"go-online-shop/infra/response"
	"golang.org/x/crypto/bcrypt"
	"log"
	"testing"
)

func TestValidateAuthEntity(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "test@test.com",
			Password: "123456",
		}

		err := authEntity.Validate()
		require.Nil(t, err)
	})

	t.Run("email is required", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "",
			Password: "123456",
		}

		err := authEntity.Validate()
		require.NotNil(t, response.ErrEmailRequired, err)
	})

	t.Run("email is invalid", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "test.com",
			Password: "123456",
		}

		err := authEntity.Validate()
		require.NotNil(t, response.ErrEmailInvalid, err)
	})

	t.Run("password is required", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "test@test.com",
			Password: "",
		}

		err := authEntity.Validate()
		require.NotNil(t, response.ErrPasswordRequired, err)
	})

	t.Run("password must have minimum 6 characters", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "test@test.com",
			Password: "123",
		}

		err := authEntity.Validate()
		require.NotNil(t, response.ErrPasswordInvalidLength, err)
	})
}

func TestEncryptPassword(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "test@test.com",
			Password: "123456",
		}

		err := authEntity.EncryptPassword(bcrypt.DefaultCost)
		require.Nil(t, err)

		log.Printf("%+v\n", authEntity)
	})
}
