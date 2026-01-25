package protocol_test

import (
	"testing"

	"beam-fs/internal/protocol"
	//"github.com/google/go-cmp/cmp"
)

func TestProtocol_ParseArgs(t *testing.T) {
	t.Parallel()
	type TestCase struct {
		got  string
		want protocol.Command
	}
}
