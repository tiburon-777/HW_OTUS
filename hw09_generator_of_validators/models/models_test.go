// +build generation

package models

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

type Validated interface {
	Validate() ([]ValidationError, error)
}

func TestUserValidation(t *testing.T) {
	requireValidation(t, User{})

	goodUser := User{
		ID:       "0a44d582-9749-11ea-a056-9ff7f30f0608",
		Name:     "John",
		Age:      24,
		Email:    "john@abrams.com",
		Role:     "admin",
		Response: Response{Code: 200},
	}
	requireNoValidationErrors(t, goodUser)

	t.Run("ID length", func(t *testing.T) {
		u := goodUser
		u.ID = "123"

		errs, err := u.Validate()
		require.Nil(t, err)
		requireOneFieldErr(t, errs, "ID")
	})

	t.Run("email regexp", func(t *testing.T) {
		u := goodUser
		u.Email = "isnotvalid@@email"

		errs, err := u.Validate()
		require.Nil(t, err)
		requireOneFieldErr(t, errs, "Email")
	})

	t.Run("age borders", func(t *testing.T) {
		u := goodUser

		for _, a := range []int{18, 34, 50} {
			u.Age = a
			errs, err := u.Validate()
			require.Nil(t, err)
			require.Len(t, errs, 0)
		}

		u.Age = 51
		errs, err := u.Validate()
		require.Nil(t, err)
		requireOneFieldErr(t, errs, "Age")
	})

	t.Run("fail phones slice", func(t *testing.T) {
		badUser := User{
			ID:       "0a44d582-9749-11ea-a056-9ff7f30f0608",
			Name:     "John",
			Age:      24,
			Email:    "john@abrams.com",
			Role:     "admin",
			Phones:   []string{"+12dfwdf343242343", "898298741293", "fdsf"},
			Response: Response{Code: 404},
		}

		errs, err := badUser.Validate()
		require.Nil(t, err)
		requireOneFieldErr(t, errs, "Phones")
	})

	t.Run("pass phones slice", func(t *testing.T) {
		goodUser := User{
			ID:       "0a44d582-9749-11ea-a056-9ff7f30f0608",
			Name:     "John",
			Age:      24,
			Email:    "john@abrams.com",
			Role:     "admin",
			Phones:   []string{"12345678901", "qazxswedcvf", "..........."},
			Response: Response{Code: 500},
		}
		requireNoValidationErrors(t, goodUser)
	})

	t.Run("embeded structure", func(t *testing.T) {
		goodUser := User{
			ID:       "0a44d582-9749-11ea-a056-9ff7f30f0608",
			Name:     "John",
			Age:      24,
			Email:    "john@abrams.com",
			Role:     "admin",
			Phones:   []string{"12345678901", "qazxswedcvf", "............"},
			Response: Response{Code: 500},
		}

		errs, err := goodUser.Validate()
		require.Nil(t, err)
		requireOneFieldErr(t, errs, "Phones")
	})

	t.Run("many errors", func(t *testing.T) {
		u := goodUser
		u.Age = -100
		u.Email = "123"
		u.Role = "unknown"

		errs, err := u.Validate()
		require.Nil(t, err)

		fields := make([]string, 0, len(errs))
		for _, e := range errs {
			fields = append(fields, e.Field)
		}
		require.ElementsMatch(t, fields, []string{"Age", "Email", "Role"})
	})
}

func TestAppValidation(t *testing.T) {
	requireValidation(t, App{})

	goodApp := App{
		Version: "1.1.0",
	}
	requireNoValidationErrors(t, goodApp)

	t.Run("version length", func(t *testing.T) {
		a := goodApp
		a.Version = "0.1"

		errs, err := a.Validate()
		require.Nil(t, err)
		requireOneFieldErr(t, errs, "Version")
	})
}

func TestTokenValidation(t *testing.T) {
	requireNoValidation(t, Token{}, "no validated fields - no Validation() method")
}

func TestResponseValidation(t *testing.T) {
	requireValidation(t, Response{})

	goodResponse := Response{
		Code: http.StatusOK,
		Body: "some body",
	}
	requireNoValidationErrors(t, goodResponse)

	t.Run("code set", func(t *testing.T) {
		r := goodResponse

		for _, c := range []int{200, 404, 500} {
			r.Code = c
			errs, err := r.Validate()
			require.Nil(t, err)
			require.Len(t, errs, 0)
		}

		r.Code = 133
		errs, err := r.Validate()
		require.Nil(t, err)
		requireOneFieldErr(t, errs, "Code")
	})
}

func requireValidation(t *testing.T, v interface{}, msgAndArgs ...interface{}) {
	_, ok := v.(Validated)
	require.True(t, ok, msgAndArgs)
}

func requireNoValidation(t *testing.T, v interface{}, msgAndArgs ...interface{}) {
	_, ok := v.(Validated)
	require.False(t, ok, msgAndArgs)
}

func requireNoValidationErrors(t *testing.T, v Validated) {
	errs, err := v.Validate()
	require.Nil(t, err)
	require.Len(t, errs, 0)
}

func requireOneFieldErr(t *testing.T, errors []ValidationError, fieldName string) {
	require.Len(t, errors, 1)
	require.Equal(t, fieldName, errors[0].Field)
	require.NotNil(t, errors[0].Err)
}
