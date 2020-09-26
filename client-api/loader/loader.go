// Package loader provides functionality to load json file
package loader

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/artzor/tech-test/client-api/entity"
)

// Loader provides an API to read and parse file
type Loader struct {
	decoder *json.Decoder
}

// ErrEOF is return when file reading is finished
var ErrEOF = errors.New("end of file")

// New returns new Loader instance
func New(reader io.Reader) (*Loader, error) {
	decoder := json.NewDecoder(reader)

	if _, err := decoder.Token(); err != nil {
		return nil, fmt.Errorf("json parse: %v", err)
	}

	return &Loader{
		decoder: decoder,
	}, nil
}

// NextRow will read next object from file and deserialize it
// it will return ErrEOF if reading finished or other error if parsing failed
func (l *Loader) NextRow() (entity.PortDetails, error) {
	tkn, err := l.decoder.Token()
	if err != nil {
		return entity.PortDetails{}, fmt.Errorf("port id: %v", err)
	}

	_, ok := tkn.(json.Delim)
	if ok {
		return entity.PortDetails{}, ErrEOF
	}

	portID, ok := tkn.(string)
	if !ok {
		return entity.PortDetails{}, fmt.Errorf("failed to parse token: %v", tkn)
	}

	if portID == "" {
		return entity.PortDetails{}, fmt.Errorf("empty port ID")
	}

	if !l.decoder.More() {
		return entity.PortDetails{}, ErrEOF
	}

	portDetails := entity.PortDetails{}
	if err := l.decoder.Decode(&portDetails); err != nil {
		return entity.PortDetails{}, err
	}

	portDetails.ID = portID
	return portDetails, nil
}
