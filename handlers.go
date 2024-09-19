package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func handler_home(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, fmt.Sprintf("%v", myLibrary.String()))
}

func handler_view(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vid := vars["vid"]

	if _, vid_exists := myLibrary.viewCountMap[vid]; vid_exists { ///1000s
		go myLibrary.incrementViewCount(vid)
		io.WriteString(w, fmt.Sprintf("Now you are watching %v", vid))
	} else {
		io.WriteString(w, "Requested video does not exist!\n")
	}
}

func handler_viewcount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vid := vars["vid"]

	if _, vid_exists := myLibrary.viewCountMap[vid]; vid_exists {
		count := myLibrary.getViewCount(vid)
		io.WriteString(w, fmt.Sprintf("count for %v is %v", vid, count))
	} else {
		io.WriteString(w, "Requested video does not exist!\n")
	}
}
