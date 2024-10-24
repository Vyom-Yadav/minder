// SPDX-FileCopyrightText: Copyright 2023 The Minder Authors
// SPDX-License-Identifier: Apache-2.0

package reconcilers

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	stubeventer "github.com/mindersec/minder/internal/events/stubs"
	"github.com/mindersec/minder/internal/reconcilers/messages"
	"github.com/mindersec/minder/pkg/eventer/constants"
)

var (
	testProviderID = uuid.New()
	testProjectID  = uuid.New()
	testRepoID     = uuid.New()
)

func Test_handleRepoReconcilerEvent(t *testing.T) {
	t.Parallel()

	scenarios := []struct {
		name            string
		expectedPublish bool
		expectedErr     bool
		entityID        uuid.UUID
		topic           string
	}{
		{
			name:            "valid event",
			topic:           constants.TopicQueueRefreshEntityByIDAndEvaluate,
			entityID:        testRepoID,
			expectedPublish: true,
			expectedErr:     false,
		},
		{
			// this is the case for gitlab. We test here that the event is published for the repo, but no errors occur
			// in this case the current code will issue the reconcile for the repo, but stop without a fatal error
			// just before reconciling artifacts - we verify that because if we hit the artifacts path, we would have
			// a bunch of other mocks to call
			name:            "event with string as upstream ID does publish",
			topic:           constants.TopicQueueRefreshEntityByIDAndEvaluate,
			entityID:        testRepoID,
			expectedPublish: true,
			expectedErr:     false,
		},
		{
			name:            "event with no upstream ID",
			entityID:        uuid.Nil,
			expectedPublish: false,
			expectedErr:     false,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			msg, err := messages.NewRepoReconcilerMessage(testProviderID, scenario.entityID, testProjectID)
			require.NoError(t, err)
			require.NotNil(t, msg)

			stubEventer := &stubeventer.StubEventer{}

			reconciler, err := NewReconciler(nil, stubEventer, nil, nil, nil)
			require.NoError(t, err)
			require.NotNil(t, reconciler)

			err = reconciler.handleRepoReconcilerEvent(msg)
			if scenario.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			if scenario.expectedPublish {
				require.Equal(t, 1, len(stubEventer.Sent))
				require.Contains(t, stubEventer.Topics, scenario.topic)
			} else {
				require.Equal(t, 0, len(stubEventer.Sent))
			}
		})
	}
}
