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
