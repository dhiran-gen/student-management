package driver

import (
	"errors"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnectToMySQL(t *testing.T) {
	testCases := []struct {
		caseId    int
		input     MySQLConfig
		expectErr error
	}{
		// Success case
		{
			caseId: 1,
			input: MySQLConfig{
				Driver:   "mysql",
				Host:     "localhost",
				User:     "vips",
				Password: "1234",
				Port:     "3306",
				Db:       "users",
			},
			expectErr: nil,
		},
		// Error case
		{
			caseId: 2,
			input: MySQLConfig{
				Driver:   "postgres",
				Host:     "localhost",
				User:     "asd",
				Password: "1234",
				Port:     "3309",
				Db:       "sass",
			},
			expectErr: errors.New("cannot connect to sql server"),
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+strconv.Itoa(tc.caseId), func(t *testing.T) {
			_, err := ConnectToMySQL(tc.input)
			if err != nil && tc.expectErr != nil && err.Error() != tc.expectErr.Error() {
				t.Errorf("TestCase[%v] Expected: \t%v \nGot: \t%v\n", tc.caseId, tc.expectErr, err)
			}
		})
	}
}
