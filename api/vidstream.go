package api

import (
	"errors"
	"fmt"
	"goku/tools"
	"log"
	"strings"

	// network
	"net/http"

	// self
	"goku/models"

	// parsers
	"github.com/PuerkitoBio/goquery"
)

type VidstreamApiClient struct {
	// Api Client
	domain         string // the vidstream domain (for replacing links)
	Url            string // the vidstream url
	DefaultQuality string // default download quality
}

func search(domain_url, name string) ([]models.Anime, error) {
	/*
		Search GogoAnime

		syntax:
			domain	- Vidstream Domain
			name		- Anime Name

		return:
			[]models.Anime, error

	*/

	// URL
	search_url := fmt.Sprintf("%s/search.html?keyword=%s", domain_url, name)

	// Make request
	rq, err := http.Get(search_url)

	// Check
	err = tools.CheckHttpErr(rq, err)
	if err != nil {
		return []models.Anime{}, err
	}

	// auto close stream
	defer rq.Body.Close()

	// load into parser
	doc, err := goquery.NewDocumentFromResponse(rq)
	if err != nil {
		return []models.Anime{}, err
	}

	// parse
	var found_anime []models.Anime
	doc.Find(".video-block").Each(func(i int, s *goquery.Selection) {
		// runs on each element
		//
		anime_name := s.Find(".name").Text()
		anime_date := s.Find(".date").Text()
		// Anime URL
		anime_url, err := s.Find("a").Attr("href")
		if err == false {
			anime_url = ""
		}
		anime_url = ResolveRelativeURL(search_url, anime_url)
		// Anime Img URL
		anime_img, err := s.Find(".img img").Attr("src")
		if err == false {
			anime_img = ""
		}
		//
		//fmt.Printf("Found Anime Named: %s\n", anime_name)
		found_anime = append(found_anime, models.Anime{Name: tools.StripWhitespace(anime_name), Date: tools.StripWhitespace(anime_date), ImageUrl: tools.StripWhitespace(anime_img), Url: tools.StripWhitespace(anime_url), SourceApi: "vidstream"})
	})

	//
	if len(found_anime) <= 0 {
		return []models.Anime{}, errors.New("no anime found.")
	}
	return found_anime, nil
}

func CreateVidstreamClientFromDomain(domain string) (VidstreamApiClient, error) {
	// Retuns a Vidstream client
	if strings.Contains(domain, ".") {
		return VidstreamApiClient{domain: domain, Url: fmt.Sprintf("https://%s", domain)}, nil
	}
	return VidstreamApiClient{}, errors.New("not a valid domain")
}

func (v *VidstreamApiClient) SearchAnime(anime_name string) ([]models.Anime, error) {
	/*
		Search for the Anime in the domain
	*/

	// checkz
	if anime_name != "" {
		return search(v.Url, anime_name)
	}
	return []models.Anime{}, nil
}

func get_anime_info_from_url(anime_url string) (models.Anime, error) {
	// Returns a Anime object
	anime := models.Anime{}

	// Get Page
	r, err := http.Get(anime_url)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	// Check
	err = tools.CheckHttpErr(r, err)
	if err != nil {
		return anime, err
	}

	// GoQuery
	doc, err := goquery.NewDocumentFromResponse(r)
	if err != nil {
		return anime, err
	}

	doc.Find(".video-info-left").Each(func(i int, selec *goquery.Selection) {
		// Iterate through the objects and
		anime.Name = tools.StripWhitespace(selec.Find(".video-details .date").Text())
		anime.Description = tools.StripWhitespace(selec.Find(".video-details .post-entry").Text())
		anime.Date = tools.StripWhitespace(selec.Find(".listing.items.lists .video-block").Last().Find(".meta .date").Text())
		anime.SourceApi = fmt.Sprintf("vidstream")
		anime.Url = anime_url
	})

	return anime, nil
}

func (v *VidstreamApiClient) GetAnimeFromUrl(anime_url string) (models.Anime, error) {

	if anime_url == "" {
		return models.Anime{}, errors.New("no anime url provided")
	} else if !strings.Contains(anime_url, "https") {
		return models.Anime{}, errors.New("url not supported")
	}

	return get_anime_info_from_url(anime_url)
}
