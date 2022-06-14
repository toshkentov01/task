package posts

import (
	"encoding/json"
	"log"
	"strconv"
	"sync"

	"github.com/go-resty/resty/v2"
)

var results []Meta = []Meta{}

func getFirstPartOfPosts(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 25; i++ {
		strI := strconv.Itoa(i)
		var URL = "https://gorest.co.in/public/v1/posts?page=" + strI + ""

		client := resty.New()

		response, err := client.R().
			SetHeader("Content-Type", "application/json").
			Get(URL)

		if err != nil {
			log.Println("Error while getting posts: ", err.Error())
			return
		}

		datas := Meta{}

		err = json.Unmarshal(response.Body(), &datas)

		if err != nil {
			log.Println("Error while unmarshalling body: ", err.Error())
			return
		}

		results = append(results, datas)
	}
}

func getSecondPartOfPosts(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 25; i < 50; i++ {
		strI := strconv.Itoa(i)
		var URL = "https://gorest.co.in/public/v1/posts?page=" + strI + ""

		client := resty.New()

		response, err := client.R().
			SetHeader("Content-Type", "application/json").
			Get(URL)

		if err != nil {
			log.Println("Error while getting posts: ", err.Error())
			return
		}

		datas := Meta{}

		err = json.Unmarshal(response.Body(), &datas)

		if err != nil {
			log.Println("Error while unmarshalling body: ", err.Error())
			return
		}

		results = append(results, datas)
	}
}

// GetPosts ...
func GetPosts() []Meta {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go getFirstPartOfPosts(&wg)
	go getSecondPartOfPosts(&wg)

	wg.Wait()

	return results
}
