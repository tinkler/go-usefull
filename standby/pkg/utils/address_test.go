package utils

import (
	"strings"
	"testing"
)

func TestToLower(t *testing.T) {
	println(strings.ToLower("0x4d70d97a2a7335D601e4cC914cb19DFF812d1F5C"))
	println(strings.ToLower("0x1B4C0bc8703D3af59322849bE01559fdb920c904"))
	println(strings.ToLower("0x2BF1004D9e80ca087BD1e089d75bc8c471995aC1"))
	// println(strings.ToLower("0xcDa802a5BFFaa02b842651266969A5Bba0c66D3e"))
	// println(strings.ToLower("0x2416092f143378750bb29b79eD961ab195CcEea5"))
	// println(strings.ToLower("0x04C0599Ae5A44757c0af6F9eC3b93da8976c150A"))

}
