package handlers

import (
	s "aseprayana-skeleton-go/services/auth"

	"testing"

	"github.com/labstack/echo/v4"
)

func Test_domainHandler_Signup(t *testing.T) {
	type fields struct {
		serviceAuth s.AuthService
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &domainHandler{
				serviceAuth: tt.fields.serviceAuth,
			}
			if err := u.Signup(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("domainHandler.Signup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
