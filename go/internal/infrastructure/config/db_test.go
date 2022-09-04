package config

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"gorm.io/gorm"
)

func TestOpen(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name    string
		args    args
		want    *gorm.DB
		wantErr bool
	}{
		{
			name: "異常系: 想定しないDSN設定時、errを返却する",
			args: args{
				dsn: "",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()

			got, _ := Open(tt.args.dsn)
			assert.Equal(t, got, tt.want)
		})
	}
}
