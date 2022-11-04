package livelint_test

import (
	"testing"

	"github.com/bespinian/livelint/internal/livelint"
	"github.com/matryer/is"
	v1 "k8s.io/api/core/v1"
)

func TestCheckForgottenCMDInDockerfile(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name       string
		container  v1.Container
		userAnswer bool

		wantFailed  bool
		wantMessage string
	}{
		{
			name:        "success - container has command",
			container:   v1.Container{Command: []string{"start_service"}},
			wantMessage: "Your container has defined a command",
		},
		{
			name:        "success - user replied yes",
			userAnswer:  true,
			wantMessage: "Your container has defined a command",
		},
		{
			name:        "failure - user replied no",
			userAnswer:  false,
			wantFailed:  true,
			wantMessage: "You forgot the CMD or ENTRYPOINT instruction in the Dockerfile",
		},
	}
	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			is := is.New(t)

			ui := &UserInteractionMock{
				AskYesNoFunc: func(_ string) bool {
					return tc.userAnswer
				},
			}

			ll := livelint.New(nil, nil, ui)

			result := ll.CheckForgottenCMDInDockerfile(tc.container)
			is.Equal(tc.wantFailed, result.HasFailed)
			is.Equal(tc.wantMessage, result.Message)
		})
	}
}
