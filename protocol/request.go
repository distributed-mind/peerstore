package protocol

import (
	"encoding/gob"

	"github.com/pkg/errors"
)

func init() {
	gob.Register(Request{})
}

// RequestMethod - this is the indication of what request is to
// be performed.
type RequestMethod uint64

var RequestMethodToString = map[RequestMethod]string{
	GetFileMethod:        "GetFile",
	PostFileMethod:       "PostFile",
	DeleteFileMethod:     "DeleteFile",
	GetSuccessorMethod:   "GetSuccessor",
	SetPredecessorMethod: "SetPredecessor",
	GetPredecessorMethod: "GetPredecessor",
	GetFingerTableMethod: "GetFingerTable",
}

const (
	// GetFileMethod - Get File Method to be used when getting files
	GetFileMethod RequestMethod = 1 << iota
	// PostFileMethod - Post File Method to be used when inserting or updating
	PostFileMethod
	// DeleteFileMethod - Delete File Method to be used when deleting files
	DeleteFileMethod
	// GetSuccessorMethod - Chord Method to get the successor
	GetSuccessorMethod
	// SetPredecessorMethod - Chord Method to set the predecessor
	SetPredecessorMethod
	// GetPredecessorMethod - Chord Method to get the predecessor
	GetPredecessorMethod
	// GetFingerTableMethod - Chord Method to get the finger table
	GetFingerTableMethod
)

// Request - the standard request, includes a header,
// method and data.  The resource is defined in the header
// and the data length is defined in the header as well.
type Request struct {
	Header Header
	Method RequestMethod
	Data   []byte
}

// Validate - implementation of Validatable, makes sure the request is
// a valid request
func (r *Request) Validate() error {
	if err := r.Header.Validate(); err != nil {
		return errors.Wrap(err, "failed to validate request header: ")
	}

	if _, ok := RequestMethodToString[r.Method]; !ok {
		return errors.New("failed to validate request method")
	}
	return nil
}
