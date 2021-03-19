package main

import "testing"

func TestSteps8pathsUDDDUDUU(t *testing.T) {
	want := int32(1)

	/* montain draw
	_/\      _
	   \    /
		\/\/
	*/
	got := CountValleysInPath(uint32(8), "UDDDUDUU")

	if got != want {
		t.Errorf("got %d wanted %d", got, want)
	}
}

func TestSteps12pathsUUUDDDDUDDUU(t *testing.T) {
	want := int32(2)

	/* montain draw
	   /\
	  /  \
	_/    \   _    _
	       \/  \  /
		        \/
	*/
	got := CountValleysInPath(uint32(12), "UUUDDDDUDDUU")

	if got != want {
		t.Errorf("got %d wanted %d", got, want)
	}
}

func TestZeroSteps8pathsUADDUDUU(t *testing.T) {

	want := int32(0)

	got := CountValleysInPath(uint32(8), "UADDUDUU")

	if got != want {
		t.Errorf("got %d wanted %d", got, want)
	}
}

func TestZeroSteps8pathsUUDDU8UU(t *testing.T) {
	want := int32(0)

	got := CountValleysInPath(uint32(8), "UUDDU8UU")

	if got != want {
		t.Errorf("got %d wanted %d", got, want)
	}
}
