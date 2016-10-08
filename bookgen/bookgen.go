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

	// if _, err := os.Stat(filename); !os.IsNotExist(err) {

	// 	err := os.Remove(filename)
	// 	if err != nil {

	// 		log.Println(err.Error())
	// 		return
	// 	}

	// }

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {

		log.Println(err.Error())
	}

	// defer f.Close()

	articles := dbhandler.GetAllUseful(session, themes, locale)

	var numberstoshuffle []int

	for num := range articles {

		numberstoshuffle = append(numberstoshuffle, num)

	}
	rand.Seed(time.Now().UTC().UnixNano())

	shuffle.Ints(numberstoshuffle)

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
