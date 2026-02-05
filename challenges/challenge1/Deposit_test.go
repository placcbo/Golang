package challenge1

import "testing"

func TestDeposit_Success(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"deposit 50 to 100", 50, 150},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc := Account{
				Owner:   "Alice",
				Balance: 100,
			}
			newAccout := acc.Deposit(tt.input)
			if newAccout.Balance != tt.want {
				t.Fatalf("expected %d, got %d", tt.want, newAccout.Balance)
			}

			if acc.Balance != 100 {
				t.Fatalf("original account should remain 100, got %d", acc.Balance)

			}
		})
	}

}
