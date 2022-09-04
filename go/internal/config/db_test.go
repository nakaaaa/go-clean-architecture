package config

import (
	"reflect"
	"testing"
)

func TestNewDBConfig(t *testing.T) {
	tests := []struct {
		name    string
		want    *DBConfig
		wantErr bool
	}{
		// TODO: なにかしら環境変数を設定すれば動く
		// {
		// 	name: "正常系: DB構成情報を返却する",
		// 	want: &DBConfig{
		// 		DSN: "test",
		// 	},
		// },
		{
			name: "異常系: nilを返却する",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := NewDBConfig()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDBConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
