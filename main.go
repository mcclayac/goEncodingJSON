package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"poetry"
	"strconv"
)

/*
Anthonys-MacBook-Pro:go mcclayac$ godoc fmt Fprintf | more
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
Fprintf formats according to a format specifier and writes to w. It
returns the number of bytes written and any write error encountered.
*/

type poemWithTitle struct {
	Title     string
	Body      poetry.Poem
	WordCount string
	TheCount  int
}

type config struct {
	Route       string
	BindAddress string   `json:"addr"`
	ValidPoems  []string `json:"valid"`
	//{"doggie.txt","cat.txt","letterA.txt"}

}

var c config

func poemHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	poemName := r.Form["name"][0]
	//fileName := "doggie.txt"
	p, err := poetry.LoadPoem(poemName)

	if err != nil {
		http.Error(w, "File Not Found", http.StatusInternalServerError)
		fmt.Printf("An Error occured reading file %s \n", poemName)
		//os.Exit(-1)
		return
	} else {

		valid := false
		for _, pp := range c.ValidPoems {
			if pp == poemName {
				valid = true
				break
			}
		}

		if valid {
			//sort.Sort(p[0])
			/*			p.SortPoem()
						pwt := poemWithTitle{poemName, p,
										strconv.FormatInt(int64(p.NumWords()),10),
										p.NumThe()}
						enc := json.NewEncoder(w)
						//enc := json.NewEncoder(w)
						//enc.Encode(p)
						enc.Encode(pwt)*/

			p.ShufflePoem()
			pwt2 := poemWithTitle{poemName, p,
				strconv.FormatInt(int64(p.NumWords()), 10),
				p.NumThe()}
			enc2 := json.NewEncoder(w)

			enc2.Encode(pwt2)

			if err != nil {
				fmt.Printf("An Error occured reading file %s \n", poemName)
				os.Exit(-1)
			}
		} else {
			fmt.Printf("Not a valid file name %s \n", poemName)
			os.Exit(-1)
		}
		// _, err = fmt.Fprintf(w, "Poem Name: %s \n\n%s\n\n",poemName, p)
	}

}

func main() {

	f, err := os.Open("config")
	if err != nil {
		os.Exit(-1)
	}
	defer f.Close()

	dec := json.NewDecoder(f)

	err = dec.Decode(&c)
	if err != nil {
		os.Exit(-1)
	}

	/*fileName := "doggie.txt"
	p, err := poetry.LoadPoem(fileName)

	if err != nil {
		fmt.Printf("An Error occured reading file %s \n", fileName)
		os.Exit(-1)
	}

	fmt.Printf("%s\n", p)
	*/

	/*	http.HandleFunc("/poem", poemHandler )
		http.ListenAndServe(":8088", nil)
	*/

	fmt.Printf("%v\n\n", c)

	http.HandleFunc(c.Route, poemHandler)
	http.ListenAndServe(c.BindAddress, nil)

}

/*
Anthonys-MacBook-Pro:go mcclayac$ curl http://127.0.0.1:8088/get\?name=doggie.txt | json_pp

*/

/*
type Values map[string][]string
Values maps a string key to a list of values. It is typically used for
query parameters and form values. Unlike in the http.Header map, the
keys in a Values map are case-sensitive.

func ParseQuery(query string) (Values, error)
ParseQuery parses the URL-encoded query string and returns a map listing
the values specified for each key. ParseQuery always returns a non-nil
map containing all the valid query parameters found; err describes the
first decoding error encountered, if any.

Query is expected to be a list of key=value settings separated by
ampersands or semicolons. A setting without an equals sign is
interpreted as a key set to an empty value.

func (v Values) Add(key, value string)
Add adds the value to key. It appends to any existing values associated
with key.

func (v Values) Del(key string)
Del deletes the values associated with key.

func (v Values) Encode() string
Encode encodes the values into ``URL encoded'' form ("bar=baz&foo=quux")
sorted by key.

func (v Values) Get(key string) string
Get gets the first value associated with the given key. If there are no
values associated with the key, Get returns the empty string. To access
multiple values, use the map directly.

func (v Values) Set(key, value string)
Set sets the key to value. It replaces any existing values.

*/
