package log_test

import (
	"io/ioutil"
	"testing"

	"github.com/komand/plugin-sdk-go2/log"
)

func BenchmarkPrintln(b *testing.B) {
	l := log.NewBufferedLogger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Println("Hey everybody!")
	}
}

func BenchmarkPrintf(b *testing.B) {
	l := log.NewBufferedLogger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Printf("Hey everybody! %s %d %+v", "Nick Riviera", 50003, []string{"stuff", "and", "things"})
	}
}

func BenchmarkWriteAndFlush5(b *testing.B) {
	l := log.NewBufferedLogger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%100 == 0 {
			if _, err := l.Flush(ioutil.Discard); err != nil {
				b.Fatal(err)
			}
		} else {
			l.Println("Hey everybody!")
		}
	}
}
