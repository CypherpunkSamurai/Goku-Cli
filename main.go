package main

import (
	"fmt"
	"goku/api"
	"log"
)

func main() {
	// Main

	v, err := api.CreateVidstreamClientFromDomain("goload.one")
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	// Search for Anime
	// a, err := v.SearchAnime("Dragon Ball")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Printf("Anime: %v\n", a)

	// Episodes
	// anim := a[0]
	anim_url := "https://goload.one/videos/dragon-ball-episode-153"

	fmt.Printf("Animeurl %s\n", anim_url)

	anime, err := v.GetAnimeFromUrl(anim_url)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	fmt.Printf("Anime: %+v", anime)
}
