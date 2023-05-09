package main

import "fmt"

type SafetyPlacerT interface {
	placeSafeties()
}

type RockClimberT struct {
	rocksClimbed int
	sp           SafetyPlacerT
}

// constructor
func newRockClimber(sp SafetyPlacerT) *RockClimberT {
	return &RockClimberT{
		sp: sp,
	}
}

type IceSafetyPlacerT struct {
	// routines to whatever data needed
}

func (sp IceSafetyPlacerT) placeSafeties() {
	fmt.Println("placing my ice safeties...")
}

type NoOpsSafetyPlacerT struct{}

func (sp NoOpsSafetyPlacerT) placeSafeties() {
	// mock implementation for testing purposes
	fmt.Println("mocking my safeties")
}

func (rc *RockClimberT) climbRock() {
	rc.rocksClimbed++
	if rc.rocksClimbed == 10 {
		rc.sp.placeSafeties()
	}
}

func main() {
	/*
		no new implementation details in any routines,
		just dependency inject mocking type

		rc := newRockClimber(NoOpsSafetyPlacerT{})
	*/
	rc := newRockClimber(IceSafetyPlacerT{})
	for i := 0; i < 11; i++ {
		rc.climbRock()
	}
}
