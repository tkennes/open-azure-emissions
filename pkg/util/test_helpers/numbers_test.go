package test_helpers

import (
	"testing"
)

func TestFloatsAlmostEqual(t *testing.T) {
	if !FloatsAlmostEqual(1.0, 1.0e0) {
		t.Errorf("Result 1a was incorrect. Expected true, got false")
	}

	if FloatsAlmostEqual(1.0, 1.0+1e-1) {
		t.Errorf("Result 1b was incorrect. Expected false, got true")
	}

	if FloatsAlmostEqual(1.0, 1.0+1e-2) {
		t.Errorf("Result 2 was incorrect. Expected false, got true")
	}

	if FloatsAlmostEqual(1.0, 1.0+1e-3) {
		t.Errorf("Result 3 was incorrect. Expected false, got true")
	}

	if FloatsAlmostEqual(1.0, 1.0+1e-4) {
		t.Errorf("Result 4 was incorrect. Expected false, got true")
	}

	if FloatsAlmostEqual(1.0, 1.0+1e-5) {
		t.Errorf("Result 5 was incorrect. Expected false, got true")
	}

	if FloatsAlmostEqual(1.0, 1.0+1e-6) {
		t.Errorf("Result 6 was incorrect. Expected false, got true")
	}

	if FloatsAlmostEqual(1.0, 1.0+1e-7) {
		t.Errorf("Result 7 was incorrect. Expected false, got true")
	}

	if FloatsAlmostEqual(1.0, 1.0+1e-8) {
		t.Errorf("Result 8 was incorrect. Expected false, got true")
	}

	if FloatsAlmostEqual(1.0, 1.0+1e-9) {
		t.Errorf("Result 9 was incorrect. Expected false, got true")
	}

	if !FloatsAlmostEqual(1.0, 1.0+1e-10) {
		t.Errorf("Result 10 was incorrect. Expected true, got false")
	}
}
