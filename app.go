package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/remotejob/huoneisto_utils/bookgen"
	"github.com/remotejob/huoneisto_utils/entryHandler"
	"github.com/remotejob/kaukotyoeu/dbhandler"
	mgo "gopkg.in/mgo.v2"
)

var themes string
var locale string
var addrs []string
var dbadmin string
var username string
var password string
var mechanism string
var sites []string
var mongoDBDialInfo mgo.DialInfo

// var tick int

func init() {

	themes = os.Getenv("THEMES")
	locale = os.Getenv("LOCALE")
	addrs = []string{os.Getenv("ADDRS")}
	dbadmin = os.Getenv("DBADMIN")
	username = os.Getenv("USERNAME")
	password = os.Getenv("PASSWORD")
	mechanism = os.Getenv("MECHANISM")
	sites = []string{os.Getenv("SITES")}

	// ktick, _ = strconv.Atoi(os.Getenv("TICK"))

	mongoDBDialInfo = mgo.DialInfo{
		Addrs:     addrs,
		Timeout:   60 * time.Second,
		Database:  dbadmin,
		Username:  username,
		Password:  password,
		Mechanism: mechanism,
	}

}

func main() {

	gocron.Every(1).Minutes().Do(Run)
	// gocron.Every(30).Second().Do(Run)

	<-gocron.Start()

	// go func() {
	// 	c := time.Tick(time.Duration(tick) * time.Second)
	// 	for range c {
	// 		// Note this purposfully runs the function
	// 		// in the same goroutine so we make sure there is
	// 		// only ever one. If it might take a long time and
	// 		// it's safe to have several running just add "go" here.
	// 		Run()
	// 	}
	// }()

	// Other processing or the rest of your program here.
	//time.Sleep(5 * time.Second)

	// Or to block forever:
	// select {}
}

//Run runner for utils
func Run() {
	// log.Println(themes)
	// log.Println(locale)
	// log.Println(addrs[0])
	// log.Println("dbadmin", dbadmin)
	// log.Println(username)
	// log.Println(password)
	// log.Println(sites[0])
	// log.Println("tick", tick)

	// pauseint := rand.Perm(tick)[0]
	// log.Println("sleeppause", pauseint)

	// time.Sleep(time.Duration(pauseint) * time.Second)

	// log.Println("end pause startdb", pauseint)

	dbsession, err := mgo.DialWithInfo(&mongoDBDialInfo)

	if err != nil {
		panic(err)
	}
	defer dbsession.Close()
	bookgen.Create(*dbsession, themes, locale, "/blog.txt")

	buf := bytes.NewBuffer(nil)

	f, err := os.Open("/blog.txt")
	if err != nil {

		log.Println(err.Error())
	}

	_, err = io.Copy(buf, f)
	if err != nil {

		log.Println(err.Error())
	}

	err = f.Close()
	if err != nil {

		log.Println(err.Error())
	}

	allsitemaplinks := dbhandler.GetAllSitemaplinks(*dbsession, sites[0])

	uniqLinks := make(map[string]struct{})

	for _, sitemaplink := range allsitemaplinks {
		uniqLinks[sitemaplink.Stitle] = struct{}{}

	}

	newArticle := entryHandler.NewEntryarticle()

	stitle := newArticle.AddTitleStitleMcontents(buf.Bytes(), sites, uniqLinks)

	if _, ok := uniqLinks[stitle]; !ok {

		uniqLinks[stitle] = struct{}{}

		newArticle.AddAuthor()
		newArticle.InsertIntoDB(*dbsession)

	} else {
		fmt.Println("Creates stitle EXIST!! but it possible", stitle)
	}

	dbsession.Close()

	log.Println("END close DB")

}
