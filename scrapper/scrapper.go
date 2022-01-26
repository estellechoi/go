package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const baseURL string = "https://kr.indeed.com/jobs?limit=50"

type job struct {
	id      string
	title   string
	summary string
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkRespCode(resp *http.Response) {
	if resp.StatusCode != 200 {
		log.Fatalln("no valid status code")
	}
}

func getPageCnt(searchText string) int {
	pageCnt := 0
	resp, err := http.Get(baseURL + "&q=" + searchText)

	checkErr(err)
	checkRespCode(resp)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	checkErr(err)

	doc.Find(".pagination").Each(func(item int, s *goquery.Selection) {
		pageCnt += s.Find("a").Length()
	})

	return pageCnt
}

func getPageURL(pageNum int, searchText string) string {
	// ..
	search := "&start=" + strconv.Itoa(pageNum*50) + "&q=" + searchText
	return baseURL + search
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func extractJob(card *goquery.Selection, c chan<- job) {
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".jobTitle > span").Text())
	summary := cleanString(card.Find(".job-snippet").Text())
	c <- job{id: id, title: title, summary: summary}
}

func fetchJobsByPage(url string, mainC chan<- []job) {
	var jobs []job
	c := make(chan job)
	resp, err := http.Get(url)

	checkErr(err)
	checkRespCode(resp)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	checkErr(err)

	cards := doc.Find(".mosaic-provider-jobcards > a")
	cards.Each(func(i int, s *goquery.Selection) {
		go extractJob(s, c)
	})

	for i := 0; i < cards.Length(); i += 1 {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func makeJobSlice(job job, c chan<- []string) {
	jobSlice := []string{job.id, job.title, job.summary}
	c <- jobSlice
}

func writeJobsCSV(jobs []job) {
	c := make(chan []string)
	file, err := os.Create("jobs.csv")

	checkErr(err)

	writer := csv.NewWriter(file)

	defer writer.Flush()

	headers := []string{"Id", "Title", "Summary"}
	wErr := writer.Write(headers)

	checkErr(wErr)

	for _, job := range jobs {
		go makeJobSlice(job, c)
	}

	for i := 0; i < len(jobs); i += 1 {
		jobSlice := <-c
		jwErr := writer.Write(jobSlice)
		checkErr(jwErr)
	}

}

func Scrap(searchText string) {
	c := make(chan []job)
	var totalJobs []job
	totalPagesCnt := getPageCnt(searchText)

	for i := 0; i < totalPagesCnt; i += 1 {
		url := getPageURL(i, searchText)
		go fetchJobsByPage(url, c)
	}

	for i := 0; i < totalPagesCnt; i += 1 {
		jobs := <-c
		totalJobs = append(totalJobs, jobs...) // merge two arrays(slices)
	}

	writeJobsCSV(totalJobs)
	fmt.Println("DONE !")

}
