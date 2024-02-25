// Copyright 2024 Stacklok, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package reminder_test

import (
	"testing"

	"github.com/stacklok/minder/internal/config/reminder"
	"github.com/stretchr/testify/assert"
)

func TestValidateConfig(t *testing.T) {
	tests := []struct {
		name   string
		config *reminder.Config
		errMsg string
	}{
		{
			name: "ValidValues",
			config: &reminder.Config{
				RecurrenceConfig: reminder.RecurrenceConfig{
					Interval:             "1h",
					BatchSize:            100,
					MaxPerProject:        10,
					MinProjectFetchLimit: 5,
					MinElapsed:           "1h",
				},
			},
		},
		{
			name: "InvalidInterval",
			config: &reminder.Config{
				RecurrenceConfig: reminder.RecurrenceConfig{
					Interval:             "1x",
					BatchSize:            100,
					MaxPerProject:        10,
					MinProjectFetchLimit: 5,
					MinElapsed:           "1h",
				},
			},
			errMsg: "invalid interval: 1x",
		},
		{
			name: "InvalidMinElapsed",
			config: &reminder.Config{
				RecurrenceConfig: reminder.RecurrenceConfig{
					Interval:             "1h",
					BatchSize:            100,
					MaxPerProject:        10,
					MinProjectFetchLimit: 5,
					MinElapsed:           "1x",
				},
			},
			errMsg: "invalid min_elapsed: 1x",
		},
		{
			name: "InvalidBatchSize",
			config: &reminder.Config{
				RecurrenceConfig: reminder.RecurrenceConfig{
					Interval:             "1h",
					BatchSize:            10,
					MaxPerProject:        10,
					MinProjectFetchLimit: 5,
					MinElapsed:           "1h",
				},
			},
			errMsg: "batch_size 10 cannot be less than max_per_project(10)*min_project_fetch_limit(5)=50",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := reminder.ValidateConfig(tt.config)
			if tt.errMsg != "" {
				assert.EqualError(t, err, tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
