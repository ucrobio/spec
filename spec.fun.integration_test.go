package spec

import (
	"fmt"
	"testing"
)

func TestSpecFunIntegrationSuite(t *testing.T) {
	expected := []string{
		"fun integration suite: before all: recovered from: panic 1",
		"fun integration suite: before all: recovered from: panic 2",
		"before each: recovered from: panic 7",
		"before each: recovered from: panic 8",
		"fun integration suite: first context: context with tests: test 1: err 1",
		"fun integration suite: first context: context with tests: test 1: err 2",
		"after each: recovered from: panic 10",
		"after each: recovered from: panic 9",
		"before each: recovered from: panic 7",
		"before each: recovered from: panic 8",
		"fun integration suite: first context: context with tests: test 2: recovered from: panic 11",
		"after each: recovered from: panic 10",
		"after each: recovered from: panic 9",
		"before each: recovered from: panic 7",
		"before each: recovered from: panic 8",
		"after each: recovered from: panic 10",
		"after each: recovered from: panic 9",
		"fun integration suite: first context: context with tests: test 4: empty test",
		"before each: recovered from: panic 7",
		"before each: recovered from: panic 8",
		"after each: recovered from: panic 10",
		"after each: recovered from: panic 9",
		"fun integration suite: after all: recovered from: panic 3",
		"fun integration suite: after all: recovered from: panic 4",
	}

	received := []string{}

	Run(
		Define(
			"fun integration suite",

			BeforeAll(func() { panic("panic 1") }, func() { panic("panic 2") }),
			AfterAll(func() { panic("panic 3") }, func() { panic("panic 4") }),

			Describe(
				"first context",

				Describe(
					"context without tests",

					BeforeEach(func() { panic("panic 5") }),
					AfterEach(func() { panic("panic 6") }),
				),

				Describe(
					"context with tests",

					BeforeEach(func() { panic("panic 7") }),
					BeforeEach(func() { panic("panic 8") }),
					AfterEach(func() { panic("panic 9") }),
					AfterEach(func() { panic("panic 10") }),

					It(
						"test 1",
						func() error { return fmt.Errorf("err 1") },
						func() error { return fmt.Errorf("err 2") },
					),

					It(
						"test 2",
						func() error { panic("panic 11") },
					),

					It(
						"test 3",
						func() error { return nil },
					),

					It(
						"test 4",
					),
				),
			),
		),
		func(err error) { received = append(received, err.Error()) },
	)

	for index := range expected {
		if expected[index] != received[index] {
			t.Error("expected:", expected[index])
			t.Error("received:", expected[index])
			t.Error("index:", index)
			t.Fail()
		}
	}
}
