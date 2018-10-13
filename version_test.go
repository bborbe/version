package version_test

import (
	"fmt"
	"testing"

	"github.com/bborbe/version"
)

func TestLess(t *testing.T) {
	testcases := []struct {
		name           string
		versionA       version.Version
		versionB       version.Version
		expectedResult bool
	}{
		{
			versionA:       "1",
			versionB:       "1",
			expectedResult: false,
		},
		{
			versionA:       "1",
			versionB:       "2",
			expectedResult: true,
		},
		{
			versionA:       "2",
			versionB:       "1",
			expectedResult: false,
		},
		{
			versionA:       "1.1",
			versionB:       "1.1",
			expectedResult: false,
		},
		{
			versionA:       "1.2",
			versionB:       "1.1",
			expectedResult: false,
		},
		{
			versionA:       "1.1",
			versionB:       "1.2",
			expectedResult: true,
		},
		{
			versionA:       "a",
			versionB:       "a",
			expectedResult: false,
		},
		{
			versionA:       "a",
			versionB:       "b",
			expectedResult: true,
		},
		{
			versionA:       "b",
			versionB:       "a",
			expectedResult: false,
		},
		{
			versionA:       "a.a",
			versionB:       "a.a",
			expectedResult: false,
		},
		{
			versionA:       "a.b",
			versionB:       "a.a",
			expectedResult: false,
		},
		{
			versionA:       "a.a",
			versionB:       "a.b",
			expectedResult: true,
		},
		{
			versionA:       "1-beta",
			versionB:       "1",
			expectedResult: true,
		},
		{
			versionA:       "1",
			versionB:       "1-beta",
			expectedResult: false,
		},
		{
			versionA:       "1.1",
			versionB:       "1",
			expectedResult: false,
		},
		{
			versionA:       "1",
			versionB:       "1.1",
			expectedResult: true,
		},
		{
			versionA:       "v1.9.9-beta.0",
			versionB:       "v1.12.1",
			expectedResult: true,
		},
		{
			versionA:       "v1.12.1",
			versionB:       "v1.9.9-beta.0",
			expectedResult: false,
		},
		// TODO
		//{
		//	versionA:       "b9",
		//	versionB:       "b10",
		//	expectedResult: true,
		//},
		//{
		//	versionA:       "b10",
		//	versionB:       "b9",
		//	expectedResult: false,
		//},
	}
	for _, tc := range testcases {
		t.Run(fmt.Sprintf("%s < %s", tc.versionA, tc.versionB), func(t *testing.T) {
			result := tc.versionA.Less(tc.versionB)
			if tc.expectedResult != result {
				t.Errorf("expected result %v but got %v", tc.expectedResult, result)
			}
		})
	}
}
