package spf

import "testing"

func TestParseSPFMechanismQualifiers(t *testing.T) {
	mechanisms := Parse("v=spf1 ip4:192.0.2.0/24 ~all")
	if len(mechanisms) != 2 {
		t.Fatalf("Parse() returned %d mechanisms, want 2", len(mechanisms))
	}
	if mechanisms[0].Name != "ip4" || mechanisms[0].Arg != "192.0.2.0/24" || mechanisms[0].Qualifier != '+' {
		t.Fatalf("first mechanism = %#v", mechanisms[0])
	}
	if mechanisms[1].Name != "all" || mechanisms[1].Qualifier != '~' {
		t.Fatalf("second mechanism = %#v", mechanisms[1])
	}
}
