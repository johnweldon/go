package go_stat

/*

/projects/<name>/

GET: Name, Count of Records, Days, Tags, BillableHours, TotalHours
PUT: Replace record 
POST: Update Project
DELETE: 

*/

import (
	"github.com/johnweldon/tcalc/timecalc"
	"net/http"
)

func logNote(note string) {
	var p = timecalc.Project{}
	p.AddTime("", "1209", "1209", "", note, "")
}

type NoteHandler struct{}
type EventHandler struct{}
type ProjectHandler struct{}

func (h ProjectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var rstr = r.URL.Path
	w.Write([]byte(rstr))
	return
}
func (h NoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var rstr = r.URL.Path
	w.Write([]byte(rstr))
	return
}

func (h EventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var rstr = r.URL.Path
	w.Write([]byte(rstr))
	return
}
