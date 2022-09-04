package database

import (
	"testing"

	"github.com/nakaaaa/go-clean-architecture/go/internal/infrastructure/config"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestOpen(t *testing.T) {
	config, err := config.NewDBConfig()
	if err != nil {
		t.Errorf("fail to config.NewConfig(): err=%v", err)
		return
	}
	type args struct {
		dsn  string
		opts []gorm.Option
	}
	tests := []struct {
		name    string
		args    args
		want    *gorm.DB
		wantErr bool
	}{
		{
			name: "正常系:DB接接を返す",
			args: args{
				dsn: config.DSN,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Open(tt.args.dsn, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				assert.NotNil(t, got)
			}
		})
	}
}
