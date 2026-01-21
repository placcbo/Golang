package utils

import (
	"testing"
)

func TestAdd(t *testing.T) {

	tests := []struct {
		a       int
		b       int
		want    int
		wantErr bool
	}{
		// expect no error
		{10, 2, 11, false},
		{0, 0, 0, false},

		// expect an error
		{-3, 2, 0, true},
		{5, -3, 0, true},
	}
	for _, tt := range tests {
		// call function
		result, err := Add(tt.a, tt.b)
		if tt.wantErr && err == nil {
			t.Fatal("expected an error, but got none")
		}
		if !tt.wantErr && err != nil {
			t.Fatal("did not expect an error, but got one")
		}

		if result != tt.want {
			t.Fatalf("expected %d, got %d", tt.want, result)
		}

	}

}
