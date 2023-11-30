package services

import (
	dto "aseprayana-skeleton-go/dto/transaction"
	m "aseprayana-skeleton-go/middlewares"
	repositories "aseprayana-skeleton-go/repositories/transaction"
	"reflect"
	"testing"
)

func Test_transactionService_GetAll(t *testing.T) {
	type fields struct {
		TransactionR repositories.TransactionRepository
		jwt          m.JWTService
	}
	type args struct {
		paginationDTO dto.GetAllRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.PaginationResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &transactionService{
				TransactionR: tt.fields.TransactionR,
				jwt:          tt.fields.jwt,
			}
			got, err := s.GetAll(tt.args.paginationDTO)
			if (err != nil) != tt.wantErr {
				t.Errorf("transactionService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transactionService.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
