package service_test

import (
	"errors"
	"testing"

	"github.com/artfunder/auth-service/service"
	"github.com/artfunder/structs"
	"github.com/gomagedon/expectate"
)

var dummyUser = structs.User{
	ID:        1,
	Firstname: "John",
	Lastname:  "Doe",
	Email:     "johndoe@example.com",
	Username:  "johndoe",
	Password:  "1234",
}

type MockUserGetter struct{}

func (MockUserGetter) GetByUsername(username string) (structs.User, error) {
	if username != dummyUser.Username {
		return structs.User{}, errors.New("No user with that username")
	}
	return dummyUser, nil
}

func (MockUserGetter) GetByEmail(email string) (structs.User, error) {
	if email != dummyUser.Email {
		return structs.User{}, errors.New("No user with that email")
	}
	return dummyUser, nil
}

type LocalLoginTest struct {
	name                 string
	userGetter           service.UserGetter
	inputUsernameOrEmail string
	inputPassword        string
	shouldReturnToken    bool
	expectedError        error
}

var localLoginTests = []LocalLoginTest{
	{
		name:                 "Works with username",
		userGetter:           new(MockUserGetter),
		inputUsernameOrEmail: "johndoe",
		inputPassword:        "1234",
		shouldReturnToken:    true,
		expectedError:        nil,
	},
	{
		name:                 "Works with email",
		userGetter:           new(MockUserGetter),
		inputUsernameOrEmail: "johndoe@example.com",
		inputPassword:        "1234",
		shouldReturnToken:    true,
		expectedError:        nil,
	},
	{
		name:                 "Fails with wrong username",
		userGetter:           new(MockUserGetter),
		inputUsernameOrEmail: "badusername",
		inputPassword:        "1234",
		shouldReturnToken:    false,
		expectedError:        service.ErrUserNotFound,
	},
	{
		name:                 "Fails with wrong email",
		userGetter:           new(MockUserGetter),
		inputUsernameOrEmail: "bademail@example.com",
		inputPassword:        "1234",
		shouldReturnToken:    false,
		expectedError:        service.ErrUserNotFound,
	},
}

func TestLocalLogin(t *testing.T) {
	for _, tc := range localLoginTests {
		t.Run(tc.name, func(t *testing.T) {
			testLocalLogin(t, tc)
		})
	}
}

func testLocalLogin(t *testing.T, tc LocalLoginTest) {
	expect := expectate.Expect(t)

	authService := service.NewAuthService(tc.userGetter)
	token, err := authService.LocalLogin(tc.inputUsernameOrEmail, tc.inputPassword)

	if tc.shouldReturnToken {
		expect(token).NotToBe("")
	} else {
		expect(token).ToBe("")
	}

	expect(err).ToBe(tc.expectedError)
}
