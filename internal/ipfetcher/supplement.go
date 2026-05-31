package ipfetcher

// SupplementalASNs are ASNs known to be Iranian that RIPE's country-resource-list
// misses. This happens when the LIR (Local Internet Registry) is registered in a
// different country but its resources are used in Iran.
//
// Entries verified via RIPE whois (country: IR on inetnum/route objects).
var SupplementalASNs = []string{
	"39368", // Mahdiar Rafiee / Serverir — LIR registered in UAE, IPs used in Iran
}

// SupplementalPrefixes are Iranian IP prefixes that RIPE's country-resource-list
// misses for the same reason: the LIR is registered outside Iran.
var SupplementalPrefixes = []string{
	"185.211.56.0/22", // AE-AIROSERVER-20170703 — country: IR, LIR in UAE
}
