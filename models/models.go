// models project models.go
package models

/*
Anime:	This is a Anime object.

It contains information about Anime,

Name		- Name of the Anime
Date		- Date of publish of Anime
Url		- Url of the Anime
Episodes	- Array (slice) of Episodes
ImageUrl	- Url of the Anime Cover

*/
type Anime struct {
	Name        string
	Description string
	Date        string
	Url         string
	ImageUrl    string
	SourceApi   string
	// Episodes []Episode
}

// Episode
type Episode struct {
	Num         int
	Url         string
	EpisodeName string
	EpisodeDate string
}
