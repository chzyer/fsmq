package api

import (
	"strings"
	"testing"

	"github.com/chzyer/muxque/prot"
)

func BenchmarkApiPing(b *testing.B) {
	que, ln := runServer(b)
	defer closeServer(que, ln)
	payload := prot.NewString(strings.Repeat("h", 40000))

	api, err := New(addr)
	if err != nil {
		b.Fatal(err)
	}
	defer api.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		api.Ping(payload)
	}
	b.StopTimer()

}
