package titleservice

// CategoryID type
type CategoryID int

// CategoryIDs
const (
	TvProgram        CategoryID = 1
	TvSegment        CategoryID = 2
	TvExtra          CategoryID = 3
	Webisode         CategoryID = 4
	WebSegment       CategoryID = 5
	WebExtra         CategoryID = 6
	WebClip          CategoryID = 7
	Simulcast        CategoryID = 8
	ChannelSimulcast CategoryID = 9
	WebLiveBroadcast CategoryID = 10
)
