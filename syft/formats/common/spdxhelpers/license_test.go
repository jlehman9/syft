package spdxhelpers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/anchore/syft/syft/pkg"
)

func Test_License(t *testing.T) {
	tests := []struct {
		name     string
		input    pkg.Package
		expected string
	}{
		{
			name:     "no licenses",
			input:    pkg.Package{},
			expected: NONE,
		},
		{
			name: "no SPDX licenses",
			input: pkg.Package{
				Licenses: []string{
					"made-up",
				},
			},
			expected: "LicenseRef-made-up",
		},
		{
			name: "with SPDX license",
			input: pkg.Package{
				Licenses: []string{
					"MIT",
				},
			},
			expected: "MIT",
		},
		{
			name: "with SPDX license expression",
			input: pkg.Package{
				Licenses: []string{
					"MIT",
					"GPL-3.0",
				},
			},
			expected: "MIT AND GPL-3.0-only",
		},
		{
			name: "cap insensitive",
			input: pkg.Package{
				Licenses: []string{
					"gpl-3.0",
				},
			},
			expected: "GPL-3.0-only",
		},
		{
			name: "debian to spdx conversion",
			input: pkg.Package{
				Licenses: []string{
					"GPL-2",
				},
			},
			expected: "GPL-2.0-only",
		},
		{
			name: "includes valid LicenseRef-",
			input: pkg.Package{
				Licenses: []string{
					"one thing first",
					"two things/#$^second",
					"MIT",
				},
			},
			expected: "LicenseRef-one-thing-first AND LicenseRef-two-things----second AND MIT",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, License(test.input))
		})
	}
}
