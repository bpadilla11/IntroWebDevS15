package mem

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

// removed makeCookie function

func makeCookie(m model, req *http.Request) (*http.Cookie, error) {
	ctx := appengine.NewContext(req)

	// DATASTORE
	err := storeDstore(m, req) // removed id from parameter
	if err != nil {
		log.Errorf(ctx, "ERROR makeCookie storeDstore: %s", err)
		return nil, err
	}

	// MEMCACHE
	err = storeMemc(m, req) // removed id from parameter
	if err != nil {
		log.Errorf(ctx, "ERROR makeCookie storeMemc: %s", err)
		return nil, err
	}

	// COOKIE
	cookie := &http.Cookie{
		Name:  "session-id",
		Value: m.ID, // changed 'id' to 'm.ID'
		// Secure: true,
		HttpOnly: true,
	}
	return cookie, nil
}
