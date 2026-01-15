package utils

import "testing"

// valid accuracy
func TestWorkerAccuracyAverage_validAccuracy(t *testing.T) {
	result, err := WorkerAccuracyAverage(60, 60, 60)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if result != 60 {
		t.Fatalf("expect average of (60 + 60 +60) / 3 %.2f", result)
	}

}
