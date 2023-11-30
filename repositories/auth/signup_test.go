package repositories

import (
	dto "aseprayana-skeleton-go/dto/auth"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func Test_authRepository_Signup(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		req dto.AuthRegisterRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.AuthRegisterResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &authRepository{
				DB: tt.fields.DB,
			}
			got, err := u.Signup(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("authRepository.Signup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authRepository.Signup() = %v, want %v", got, tt.want)
			}
		})
	}
}
