package titleservice

import "testing"

func TestValidCategoryID(t *testing.T) {
	for _, tt := range []struct {
		id   CategoryID
		want bool
	}{
		{TvProgram, true},
		{TvSegment, true},
		{TvExtra, true},
		{Webisode, true},
		{WebSegment, true},
		{WebExtra, true},
		{WebClip, true},
		{Simulcast, true},
		{ChannelSimulcast, true},
		{WebLiveBroadcast, true},
		{CategoryID(-1), false},
		{CategoryID(0), false},
		{CategoryID(11), false},
	} {
		if got := validCategoryID(tt.id); got != tt.want {
			t.Fatalf("validCategoryID(%d) = %v, want %v", tt.id, got, tt.want)
		}
	}
}

func TestLookupLiveChannelID(t *testing.T) {
	for _, tt := range []struct {
		name string
		ok   bool
		id   LiveChannelID
	}{
		{"", false, 0},
		{"svt1", true, SVT1},
		{"TV4 Film", true, TV4Film},
		{"Kanal 11", true, Kanal11},
		{"UNKNOWN", false, 0},
	} {
		id, ok := LookupLiveChannelID(tt.name)

		if ok != tt.ok {
			t.Fatalf("ok = %v, want %v", ok, tt.ok)
		}

		if id != tt.id {
			t.Fatalf("id = %d, want %d", id, tt.id)
		}
	}
}

func TestValidLiveChannelID(t *testing.T) {
	for _, tt := range []struct {
		id   LiveChannelID
		want bool
	}{
		{SVT1, true},
		{TV4, true},
		{LiveChannelID(1211), true},
		{LiveChannelID(0), false},
		{LiveChannelID(9999), false},
	} {
		if got := validLiveChannelID(tt.id); got != tt.want {
			t.Fatalf("validLiveChannelID(%d) = %v, want %v", tt.id, got, tt.want)
		}
	}
}

func TestAllChannels(t *testing.T) {
	for id, name := range allMMSChannels {
		if !validLiveChannelID(id) {
			t.Fatalf("validLiveChannelID(%d) = false, want true", id)
		}

		if got, ok := LookupLiveChannelID(name); !ok || got != id {
			t.Fatalf("LookupLiveChannelID(%q) = %q true, want %q true", name, got, id)
		}
	}
}

var allMMSChannels = map[LiveChannelID]string{
	1001: "svt1",
	1002: "svt2",
	1005: "Discovery",
	1015: "Kanal5",
	1028: "TV3",
	1029: "TV4",
	1043: "Kanal9",
	1047: "Discovery World",
	1048: "Discovery Science",
	1050: "svtB",
	1052: "TV3 Sport HD",
	1055: "Esports TV",
	1061: "Sjuan",
	1070: "TV4 Film",
	1074: "TV6",
	1089: "TV4 Sport",
	1103: "TV4 Fakta",
	1153: "TV4 Guld",
	1155: "TV4 Komedi",
	1197: "TV8",
	1199: "Animal Planet",
	1211: "svt24",
	1424: "Discovery HD Showcase",
	1508: "TV12",
	1671: "Kunskapskanalen",
	1683: "TV10",
	1733: "TLC",
	1736: "Investigation Discovery",
	1796: "TV4 Fakta XL",
	1906: "Eurosport 1",
	2025: "Kanal 11",
	2047: "Eurosport 2 Sweden",
}
