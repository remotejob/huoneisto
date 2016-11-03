package bookgen

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/remotejob/kaukotyoeu_utils/dbhandler"
	"github.com/shogo82148/go-shuffle"
	"gopkg.in/mgo.v2"
)

//Create create file
func Create(session mgo.Session, themes string, locale string, filename string) {

	log.Println("Start Create file", filename, themes, locale)

	// articles := dbhandler.GetAllUseful(session, strings.TrimSpace(themes), strings.TrimSpace(locale))
	articles := dbhandler.GetAllUseful(session, "realestate", "fi_FI")

	log.Println("GET keywords", len(articles))

	var numberstoshuffle []int

	for num := range articles {

		numberstoshuffle = append(numberstoshuffle, num)

	}
	rand.Seed(time.Now().UTC().UnixNano())

	shuffle.Ints(numberstoshuffle)

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {

		log.Println(err.Error())
	}

	// defer f.Close()

	for _, i := range numberstoshuffle {

		paragraph := articles[i].Phrase + "\n"

		if _, err = f.WriteString(paragraph); err != nil {

			log.Println(err.Error())
		}
	}
	err = f.Close()
	if err != nil {
		log.Println(err.Error())
	}
}
