package mem

import (
	"github.com/nu7hatch/gouuid"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func newVisitor(req *http.Request) (*http.Cookie, error) {
	// moved initialModel to line 18
	id, err := uuid.NewV4()
	if err != nil {
		ctx := appengine.NewContext(req)
		log.Errorf(ctx, "ERROR newVisitor uuid.NewV4: %s", err)
		return nil, err
	}
	m := initialModel(id.String()) // moved initialModel with 'id' parameter as a string
	return makeCookie(m, req) // removed 'id.String()' from parameter
}

func currentVisitor(m model, req *http.Request) (*http.Cookie, error) {
	return makeCookie(m, req) // removed 'id' from parameter
}

func initialModel(id string) model {
	m := model{
		Name:  "",
		State: false,
		Pictures: []string{
			"one.jpg",
		},
		ID: id, // added
	}
	return m
}
