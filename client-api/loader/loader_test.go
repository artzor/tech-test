package loader

import (
	"clientapi/entity"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoader_NextRow(t *testing.T) {

	tt := []struct {
		sample   string
		wantRows []entity.PortDetails
		wantErr  string
	}{
		{
			sample: sampleGood,
			wantRows: []entity.PortDetails{
				{
					ID:       "AAAAA",
					Name:     "AAA",
					City:     "aaaa",
					Country:  "aaaaa",
					Alias:    []string{},
					Coords:   []float64{55.5136433, 25.4052165},
					Province: "AAAAA",
					Timezone: "Asia/Dubai",
					Unlocs:   []string(nil),
					Code:     "232323",
					Regions:  []string{},
				},
				{
					ID:       "BBBBB",
					Name:     "bbbb",
					City:     "BBBBB",
					Country:  "BBBB",
					Alias:    []string{},
					Coords:   []float64{54.37, 24.47},
					Province: "BBB",
					Timezone: "Asia/Dubai",
					Unlocs:   []string{"BBB"},
					Code:     "52001",
					Regions:  []string{}},
			},
			wantErr: "",
		},
		{
			sample:   sampleBadJson,
			wantRows: nil,
			wantErr:  `port id: invalid character '['`,
		},
		{
			sample:   sampleEmptyID,
			wantRows: nil,
			wantErr:  "empty port ID",
		},
	}

	for _, test := range tt {
		l, err := New(strings.NewReader(test.sample))

		require.NoError(t, err)
		var haveRows []entity.PortDetails

		for {
			row, err := l.NextRow()
			if errors.Is(err, ErrEOF) {
				assert.Equal(t, test.wantRows, haveRows)
				break
			}

			if err != nil {
				assert.EqualError(t, err, test.wantErr)
				break
			}
			haveRows = append(haveRows, row)
		}

		assert.Equal(t, test.wantRows, haveRows)
	}
}

const sampleGood = `{
  "AAAAA": {
    "name": "AAA",
    "city": "aaaa",
    "country": "aaaaa",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "AAAAA",
    "timezone": "Asia/Dubai",
    "code": "232323"
  },
  "BBBBB": {
    "name": "bbbb",
    "coordinates": [
      54.37,
      24.47
    ],
    "city": "BBBBB",
    "province": "BBB",
    "country": "BBBB",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "BBB"
    ],
    "code": "52001"
  }
}`

const sampleBadJson = `{
[
  "AAAAA":
    "name": "AAA",
    "city": "aaaa",
    "country": "aaaaa",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "AAAAA",
    "timezone": "Asia/Dubai",
    "code": "232323"
  }
]
}`

const sampleEmptyID = `{
  "": {
    "name": "AAA",
    "city": "aaaa",
    "country": "aaaaa",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "AAAAA",
    "timezone": "Asia/Dubai",
    "code": "232323"
  }
}
`
