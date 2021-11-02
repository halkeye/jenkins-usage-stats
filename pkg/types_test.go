package pkg_test

import (
	"testing"
	"time"

	"github.com/abayer/jenkins-usage-stats/pkg"
	"github.com/stretchr/testify/assert"
)

func TestTimestampFuncs(t *testing.T) {
	testCases := []struct {
		orig        string
		reorganized string
		timestamp   time.Time
		err         error
	}{
		{
			orig:        "14/Feb/2008:03:44:55 +0000",
			reorganized: "2008-02-14T03:44:55Z",
			timestamp:   time.Date(2008, time.February, 14, 3, 44, 55, 0, time.UTC),
		},
		{
			orig:        "02/Mar/2014:21:23:59 +0000",
			reorganized: "2014-03-02T21:23:59Z",
			timestamp:   time.Date(2014, time.March, 2, 21, 23, 59, 0, time.UTC),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.orig, func(t *testing.T) {
			assert.Equal(t, tc.reorganized, pkg.JSONTimestampToRFC3339(tc.orig))
			r := &pkg.JSONReport{TimestampString: tc.orig}
			ts, err := r.Timestamp()
			if tc.err != nil {
				assert.Equal(t, tc.err, err)
			} else {
				assert.Equal(t, tc.timestamp, ts)
			}
		})
	}
}
