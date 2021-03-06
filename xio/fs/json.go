package fs

import (
	"bufio"
	"encoding/json"
	"io"
	"os"

	"github.com/jholder85638/toolbox/errs"
	"github.com/jholder85638/toolbox/xio"
	"github.com/jholder85638/toolbox/xio/fs/safe"
)

// LoadJSON data from the specified path.
func LoadJSON(path string, data interface{}) error {
	f, err := os.Open(path)
	if err != nil {
		return errs.Wrap(err)
	}
	defer xio.CloseIgnoringErrors(f)
	if err = json.NewDecoder(bufio.NewReader(f)).Decode(data); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

// SaveJSON data to the specified path.
func SaveJSON(path string, data interface{}, format bool) error {
	return safe.WriteFile(path, func(w io.Writer) error {
		encoder := json.NewEncoder(w)
		if format {
			encoder.SetIndent("", "  ")
		}
		return errs.Wrap(encoder.Encode(data))
	})
}
