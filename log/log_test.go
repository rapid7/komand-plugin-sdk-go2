package log_test

import (
	"io/ioutil"
	"testing"

	"github.com/komand/plugin-sdk-go2/log"
)

func BenchmarkError(b *testing.B) {
	l := log.NewBufferedLogger(log.Info)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Error("Hey everybody!")
	}
}

func BenchmarkErrorf(b *testing.B) {
	l := log.NewBufferedLogger(log.Info)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Errorf("Hey everybody! %s %d %+v", "Nick Riviera", 50003, []string{"stuff", "and", "things"})
	}
}

func BenchmarkWriteAndFlush5(b *testing.B) {
	l := log.NewBufferedLogger(log.Info)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%100 == 0 {
			if _, err := l.Flush(ioutil.Discard); err != nil {
				b.Fatal(err)
			}
		} else {
			l.Error("Hey everybody!")
		}
	}
}
