package auth

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go-online-shop/external/database"
	"go-online-shop/infra/response"
	"go-online-shop/internal/config"
	"testing"
)

var svc service

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	repo := newRepository(db)
	svc = newService(repo)
}

func TestRegister_Success(t *testing.T) {
	req := RegisterRequestPayload{
		Email:    fmt.Sprintf("%v@test.com", uuid.NewString()),
		Password: "password123",
	}
	err := svc.register(context.Background(), req)
	require.Nil(t, err)
}

func TestRegister_Fail(t *testing.T) {
	t.Run("error email already used", func(t *testing.T) {
		// preparation for duplicate email
		email := fmt.Sprintf("%v@test.com", uuid.NewString())
		req := RegisterRequestPayload{
			Email:    email,
			Password: "password123",
		}
		err := svc.register(context.Background(), req)
		require.Nil(t, err)
		// end preparation

		err = svc.register(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailAlreadyUsed, err)
	})
}
