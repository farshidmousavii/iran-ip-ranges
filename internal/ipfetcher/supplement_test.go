package ipfetcher

import (
	"strings"
	"testing"
)

func TestSupplementalASNsNotEmpty(t *testing.T) {
	if len(SupplementalASNs) == 0 {
		t.Error("SupplementalASNs should not be empty")
	}
}

func TestSupplementalPrefixesNotEmpty(t *testing.T) {
	if len(SupplementalPrefixes) == 0 {
		t.Error("SupplementalPrefixes should not be empty")
	}
}

func TestSupplementalPrefixesAreCIDR(t *testing.T) {
	for _, p := range SupplementalPrefixes {
		if !strings.Contains(p, "/") {
			t.Errorf("SupplementalPrefix %q is not a CIDR notation", p)
		}
	}
}

func TestSupplementalASNsAreNumeric(t *testing.T) {
	for _, asn := range SupplementalASNs {
		if asn == "" {
			t.Error("SupplementalASN entry is empty")
		}
		for _, c := range asn {
			if c < '0' || c > '9' {
				t.Errorf("SupplementalASN %q is not numeric", asn)
				break
			}
		}
	}
}
