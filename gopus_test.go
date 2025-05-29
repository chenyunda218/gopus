package goopus

import (
	"os"
	"testing"
)

func TestOpus(t *testing.T) {
	f, _ := os.ReadFile("sources/tts_notify.opus")
	PrintOpusConfig(f)
	f, _ = os.ReadFile("sources/chinese.opus")
	PrintOpusConfig(f)
	f, _ = os.ReadFile("sources/recved.opus")
	PrintOpusConfig(f)
	f, _ = os.ReadFile("sources/output.opus")
	PrintOpusConfig(f)
}
