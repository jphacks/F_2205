package generate

import "testing"

func Test_RandomStr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		length        int
		wantStrLength int
		wantErr       error
	}{
		{
			name:          "正常に動いた場合、指定した文字数の文字列を返す",
			length:        4,
			wantStrLength: 4,
			wantErr:       nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MakeRandomStrFromLetters(tt.length)
			gotLength := len(got)
			if gotLength != tt.wantStrLength {
				t.Errorf("Test_RandomStr Error : want %v, but got %v", tt.wantStrLength, gotLength)
			}
			if err != nil && tt.wantErr != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("Test_RandomStr Error : want %v, but got %v", tt.wantErr, err)
			}
		})
	}
}
