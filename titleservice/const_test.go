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
