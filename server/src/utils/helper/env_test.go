package helper

import "testing"

func TestEnv_GetEnvOrDefault(t *testing.T) {
	tests := []struct {
		name       string
		envPath    string
		defaultEnv string
		want       string
	}{
		{
			name:       "環境変数hogeに値hogehogeがセットされている場合、hogehogeが返る",
			envPath:    "hoge",
			defaultEnv: "fuga",
			want:       "hogehoge",
		},
		{
			name:       "環境変数hoge2に値がセットされていない場合、defaultEnvのfugaが返る",
			envPath:    "hoge2",
			defaultEnv: "fuga",
			want:       "fuga",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 環境変数hogeにhogehogeをセットする
			t.Setenv("hoge", "hogehoge")
			got := GetEnvOrDefault(tt.envPath, tt.defaultEnv)
			if got != tt.want {
				t.Errorf("TestEnv_GetEnvOrDefault Error : want %s, but got %s", tt.want, got)
			}
		})
	}
}
