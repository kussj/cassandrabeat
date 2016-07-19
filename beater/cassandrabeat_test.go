package beater

import (
	"testing"
)

func TestGetLatency(t *testing.T) {
	/**
	 * Should return both ready and write latency as floats.
	 * TODO: test for latency of NaN
	*/

	// Arrange
	test_table := "system.local"
	expected_read := 3.349
	expected_write := 1.074
	
	// Act
	actual_read, actual_write := getLatency(test_table)

	// Assert
	if actual_read != expected_read {
		t.Errorf("getLatency(%s): expected_read %f, actual_read %f", test_table, expected_read, actual_read)
	}
	if actual_write != expected_write {
		t.Errorf("getLatency(%s): expected_write %f, actual_write %f", test_table, expected_write, actual_write)
	}
}

