package utils

import (
	"testing"
)

func TestUtils(t *testing) {
	t.Run("ToXMLTag", func(t *testing.T) {
		if ToXMLTag("tag", "val") != "<tag>val</tag>" {
			t.Error("bad ToXMLTag")
		}
	})
}