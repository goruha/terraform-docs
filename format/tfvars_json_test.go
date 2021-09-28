/*
Copyright 2021 The terraform-docs Authors.
Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.
You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package format

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/terraform-docs/terraform-docs/internal/testutil"
	"github.com/terraform-docs/terraform-docs/print"
)

func TestTfvarsJson(t *testing.T) {
	tests := map[string]struct {
		config print.Config
	}{
		// Base
		"Base": {
			config: testutil.WithSections(),
		},
		"Empty": {
			config: testutil.WithDefaultSections(
				testutil.With(func(c *print.Config) {
					c.ModuleRoot = "empty"
				}),
			),
		},

		// Settings
		"EscapeCharacters": {
			config: testutil.With(func(c *print.Config) {
				c.Settings.Escape = true
			}),
		},
		"SortByName": {
			config: testutil.WithSections(
				testutil.With(func(c *print.Config) {
					c.Sort.Enabled = true
					c.Sort.By = print.SortName
				}),
			),
		},
		"SortByRequired": {
			config: testutil.WithSections(
				testutil.With(func(c *print.Config) {
					c.Sort.Enabled = true
					c.Sort.By = print.SortRequired
				}),
			),
		},
		"SortByType": {
			config: testutil.WithSections(
				testutil.With(func(c *print.Config) {
					c.Sort.Enabled = true
					c.Sort.By = print.SortType
				}),
			),
		},

		// No section
		"NoInputs": {
			config: testutil.WithSections(
				testutil.With(func(c *print.Config) {
					c.Sections.Inputs = false
				}),
			),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			expected, err := testutil.GetExpected("tfvars", "json-"+name)
			assert.Nil(err)

			module, err := testutil.GetModule(&tt.config)
			assert.Nil(err)

			formatter := NewTfvarsJSON(&tt.config)

			err = formatter.Generate(module)
			assert.Nil(err)

			actual, err := formatter.ExecuteTemplate("")

			assert.Nil(err)
			assert.Equal(expected, actual)
		})
	}
}
