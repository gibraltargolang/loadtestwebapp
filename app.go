/**
 * @author: thomasmodeneis
 */

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/tsenart/vegeta/lib"
)

func main() {

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/", fs)
	http.HandleFunc("/benchtest", benchtest)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)

}

//Load intrface
type Load struct {
	Name     string
	Site     string
	Method   string
	Body     string
	Duration string
}

//Resp response object
type Resp struct {
	Jsondata string
	Htmldata string
	Textdata string
}

func benchtest(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t Load
	err := decoder.Decode(&t)
	if err != nil {
		log.Println(err.Error())
		return
	}

	resp, err := load(t)

	js, err := json.Marshal(resp)
	if err != nil {
		return
	}

	log.Println("ok")
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)

}

//Load func for doing the loadtesting
func load(l Load) (Resp, error) {
	rate := uint64(100)                        // per second
	i, _ := strconv.Atoi(l.Duration)           //cast
	duration := time.Duration(i) * time.Second //calculate

	//configure target
	targeter := vegeta.NewStaticTargeter(&vegeta.Target{
		Method: l.Method,
		URL:    l.Site,
		Body:   []byte(l.Body),
	})
	attacker := vegeta.NewAttacker()

	// define results
	var results vegeta.Results
	for res := range attacker.Attack(targeter, rate, duration) {
		results = append(results, res)
	}

	//configure to generate a json report
	rep := vegeta.ReportJSON
	jsondata, _ := rep.Report(results)

	//configure to generate a html report
	rep = vegeta.ReportPlot
	htmldata, _ := rep.Report(results)

	//configure to generate a text report
	rep = vegeta.ReportText
	textdata, _ := rep.Report(results)

	//return response object and nil error
	r := Resp{string(jsondata), string(htmldata), string(textdata)}
	return r, nil
}
