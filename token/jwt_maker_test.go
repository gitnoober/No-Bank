package token

import (
	"fmt"
	"testing"
	"time"

	"github.com/gitnoober/No-Bank/utils"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	t.Parallel()
	maker, err := NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)

	username := utils.RandomOwner()
	duration := time.Minute

	// issuedAt := time.Now()
	// expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	fmt.Println(token, "token")
	require.NoError(t, err)
	require.NotEmpty(t, token)

	claims, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, claims)

	require.NotZero(t, claims["id"])
	require.Equal(t, username, claims["user_name"])

	// var tm time.Time
	// switch iat := claims["issued_at"].(type) {
	// case float64:
	// 	tm = time.Unix(int64(iat), 0)
	// case json.Number:
	// 	v, _ := iat.Int64()
	// 	tm = time.Unix(v, 0)
	// }
	// fmt.Print("dpfk", tm)\
	// issued_at, err := time.Parse("2023-03-11T02:14:07.170578+05:30", strconv.ParseTime claims["issued_at"].(string))

	// require.WithinDuration(t, issuedAt, issued_at, time.Second)
	// require.WithinDuration(t, expiredAt, claims["expired_at"], time.Second)

}
