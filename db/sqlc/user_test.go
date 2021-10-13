package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)
import "github.com/Ruadgedy/simplebank/util"

func createRandomUser(t *testing.T) User {
	hashedPassword,err := util.HashPassword(util.RandomString(6))
	require.NoError(t,err)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	sqlResult, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t,err)
	require.NotEmpty(t,sqlResult)

	testQueries.GetUser()

	//require.Equal(t, arg.Username, user.Username)
	//require.Equal(t, arg.HashedPassword, user.HashedPassword)
	//require.Equal(t, arg.FullName, user.FullName)
	//require.Equal(t, arg.Email, user.Email)
	//require.True(t, user.PasswordChangedAt.IsZero())
	//require.NotZero(t, user.CreatedAt)

	return user
}
