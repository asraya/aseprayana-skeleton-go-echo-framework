package repositories

import (
	dto "aseprayana-skeleton-go/dto/auth"
	"aseprayana-skeleton-go/entity"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func Test_authRepository_Signin(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		req dto.AuthLoginRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &authRepository{
				DB: tt.fields.DB,
			}
			got, err := u.Signin(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("authRepository.Signin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authRepository.Signin() = %v, want %v", got, tt.want)
			}
		})
	}
}
