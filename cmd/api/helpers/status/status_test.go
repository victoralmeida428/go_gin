package status_test

import (
	"abramed_go/cmd/api/helpers"
	"abramed_go/cmd/api/helpers/status"
	"testing"
	"time"
)

type input struct {
	totalRequired, totalOptional, required, optional int
	deadline, report_opt, report_required            time.Time
}
type TestCase struct {
	input    input
	expected status.StatusForm
}

func getMock() []TestCase {
	today := time.Now()
	return []TestCase{
		{
			input: input{
				totalRequired:   4,
				required:        4,
				totalOptional:   16,
				optional:        16,
				deadline:        time.Date(2024, 12, 20, 0, 0, 0, 0, time.Local),
				report_opt:      time.Date(2024, 12, 20, 0, 0, 0, 0, time.Local),
				report_required: time.Date(2024, 12, 20, 0, 0, 0, 0, time.Local),
			},
			expected: status.COMPLETE,
		},
		{
			input: input{
				totalRequired:   4,
				required:        4,
				totalOptional:   16,
				optional:        2,
				deadline:        time.Date(2025, 1, 4, 0, 0, 0, 0, time.Local),
				report_opt:      time.Date(2025, 1, 3, 0, 0, 0, 0, time.Local),
				report_required: time.Date(2025, 1, 3, 0, 0, 0, 0, time.Local),
			},
			expected: status.PARTIAL,
		},
		{
			input: input{
				totalRequired:   4,
				required:        4,
				totalOptional:   16,
				optional:        0,
				deadline:        time.Date(2024, 12, 21, 0, 0, 0, 0, time.Local),
				report_opt:      time.Date(2024, 12, 20, 0, 0, 0, 0, time.Local),
				report_required: time.Date(2024, 12, 20, 0, 0, 0, 0, time.Local),
			},
			expected: status.PARTIAL,
		},
		{
			input: input{
				totalRequired:   4,
				required:        2,
				totalOptional:   16,
				optional:        0,
				deadline:        today.Add(helpers.Day),
				report_opt:      today,
				report_required: today,
			},
			expected: status.WAITING,
		},
		{
			input: input{
				totalRequired:   4,
				required:        4,
				totalOptional:   16,
				optional:        0,
				deadline:        time.Date(2024, 12, 10, 0, 0, 0, 0, time.Local),
				report_opt:      time.Date(2024, 12, 20, 0, 0, 0, 0, time.Local),
				report_required: time.Date(2024, 12, 20, 0, 0, 0, 0, time.Local),
			},
			expected: status.PARTIALBACKDATE,
		},
		{
			input: input{
				totalRequired:   4,
				required:        2,
				totalOptional:   16,
				optional:        0,
				deadline:        time.Date(2024, 12, 10, 0, 0, 0, 0, time.Local),
				report_opt:      time.Date(2024, 12, 20, 0, 0, 0, 0, time.Local),
				report_required: time.Date(2024, 12, 20, 0, 0, 0, 0, time.Local),
			},
			expected: status.PENDING,
		},
		{
			input: input{
				totalRequired:   4,
				required:        4,
				totalOptional:   16,
				optional:        16,
				deadline:        time.Date(2024, 12, 10, 0, 0, 0, 0, time.Local),
				report_opt:      time.Date(2024, 12, 20, 0, 0, 0, 0, time.Local),
				report_required: time.Date(2024, 12, 20, 0, 0, 0, 0, time.Local),
			},
			expected: status.COMPLETEBACKDATE,
		},
	}
}
func TestStatusForm(t *testing.T) {

	mock := getMock()
	for _, data := range mock {
		t.Run(data.expected.ToString(), func(t *testing.T) {
			target := status.GetStatusForm(
				data.input.totalRequired,
				data.input.totalOptional,
				data.input.required,
				data.input.optional,
				data.input.deadline,
				data.input.report_opt,
				data.input.report_required,
			)
			if target != data.expected {
				t.Errorf("input: %+v\nExpected/Target: %s != %s", data.input, data.expected.ToString(), target.ToString())
			}

		})
	}

}

func BenchmarkStatus(b *testing.B) {
	data := getMock()[0]

	for n := 0; n < b.N; n++ {
		status.GetStatusForm(
			data.input.totalRequired,
			data.input.totalOptional,
			data.input.required,
			data.input.optional,
			data.input.deadline,
			data.input.report_opt,
			data.input.report_required,
		)
	}
}
