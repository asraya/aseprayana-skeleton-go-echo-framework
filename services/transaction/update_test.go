package services

import (
	dto "aseprayana-skeleton-go/dto/transaction"
	m "aseprayana-skeleton-go/middlewares"
	repositories "aseprayana-skeleton-go/repositories/transaction"
	"reflect"
	"testing"
)

func Test_transactionService_Update(t *testing.T) {
	type fields struct {
		TransactionR repositories.TransactionRepository
		jwt          m.JWTService
	}
	type args struct {
		req dto.UpdateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.TransactionResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &transactionService{
				TransactionR: tt.fields.TransactionR,
				jwt:          tt.fields.jwt,
			}
			got, err := b.Update(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("transactionService.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transactionService.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
