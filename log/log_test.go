package log_test

import (
	"io/ioutil"
	"testing"

	"github.com/komand/plugin-sdk-go2/log"
)

func BenchmarkInfo(b *testing.B) {
	l := log.NewLogger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Info("Hey everybody!")
	}
}

func BenchmarkInfof(b *testing.B) {
	l := log.NewLogger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Infof("Hey everybody! %s %d %+v", "Nick Riviera", 50003, []string{"stuff", "and", "things"})
	}
}

func BenchmarkWriteAndFlush5(b *testing.B) {
	l := log.NewLogger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%100 == 0 {
			if _, err := l.Flush(ioutil.Discard); err != nil {
				b.Fatal(err)
			}
		} else {
			l.Info("Hey everybody!")
		}
	}
}
