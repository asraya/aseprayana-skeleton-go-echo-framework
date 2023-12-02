package handlers

import (
	m "aseprayana-skeleton-go/middlewares"
	s "aseprayana-skeleton-go/services/transaction"
	"testing"

	"github.com/labstack/echo/v4"
)

func Test_domainHandler_Create(t *testing.T) {
	type fields struct {
		domainService s.TransactionService
		jwt           m.JWTService
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
			b := &domainHandler{
				domainService: tt.fields.domainService,
				jwt:           tt.fields.jwt,
			}
			if err := b.Create(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("domainHandler.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
