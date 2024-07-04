package main

import (
	"reflect"
	"testing"
)

func TestParseCommentsToWorkflowNames(t *testing.T) {
	tests := []struct {
		name        string
		commentBody string
		expected    map[string]struct{}
	}{
		{
			name:        "workflow with spaces in quotes",
			commentBody: `/rerun-workflow "my workflow"`,
			expected:    map[string]struct{}{"my workflow": {}},
		},
		{
			name:        "single word in quotes",
			commentBody: `/rerun-workflow "my"`,
			expected:    map[string]struct{}{"my": {}},
		},
		{
			name:        "single word without quotes",
			commentBody: `/rerun-workflow my`,
			expected:    map[string]struct{}{"my": {}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseCommentsToWorkflowNames(tt.commentBody)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("parseCommentsToWorkflowNames() = %v, want %v", got, tt.expected)
			}
		})
	}
}
