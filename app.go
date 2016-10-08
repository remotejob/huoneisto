package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"time"

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

func init() {

	themes = os.Getenv("THEMES")
	locale = os.Getenv("LOCALE")
	addrs = []string{os.Getenv("ADDRS")}
	dbadmin = os.Getenv("DBADMIN")
	username = os.Getenv("USERNAME")
	password = os.Getenv("PASSWORD")
	mechanism = os.Getenv("MECHANISM")
	sites = []string{os.Getenv("SITES")}
}

func main() {

	log.Println(themes)
	log.Println(locale)
	log.Println(addrs[0])
	log.Println("dbadmin", dbadmin)
	log.Println(username)
	log.Println(password)
	log.Println(sites[0])

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:     addrs,
		Timeout:   60 * time.Second,
		Database:  dbadmin,
		Username:  username,
		Password:  password,
		Mechanism: mechanism,
	}

	dbsession, err := mgo.DialWithInfo(mongoDBDialInfo)

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
		// fmt.Println(newArticle.Modarticle.Author)
		// fmt.Println(newArticle.Modarticle.Site)
		// fmt.Println(newArticle.Modarticle.Title)
		// fmt.Println(newArticle.Modarticle.Stitle)
		// fmt.Println(newArticle.Modarticle.Mcontents)

		// fmt.Println(newArticle.Modarticle.Contents)

	} else {
		fmt.Println("Creates stitle EXIST!! but it possible", stitle)
	}

}
