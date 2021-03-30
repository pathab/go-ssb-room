package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ssb-ngi-pointer/go-ssb-room/roomdb"
	"github.com/ssb-ngi-pointer/go-ssb-room/web"
	weberrors "github.com/ssb-ngi-pointer/go-ssb-room/web/errors"
	"github.com/ssb-ngi-pointer/go-ssb-room/web/router"
	"github.com/ssb-ngi-pointer/go-ssb-room/web/webassert"
	refs "go.mindeco.de/ssb-refs"
)

func TestInviteShowAcceptForm(t *testing.T) {
	ts := setup(t)

	urlTo := web.NewURLTo(ts.Router)

	t.Run("token doesnt exist", func(t *testing.T) {
		a, r := assert.New(t), require.New(t)

		testToken := "nonexistant-test-token"
		acceptURL404 := urlTo(router.CompleteInviteFacade, "token", testToken)
		r.NotNil(acceptURL404)

		// prep the mocked db for http:404
		ts.InvitesDB.GetByTokenReturns(roomdb.Invite{}, roomdb.ErrNotFound)

		// request the form
		acceptForm := acceptURL404.String()
		t.Log(acceptForm)
		doc, resp := ts.Client.GetHTML(acceptForm)
		// 500 until https://github.com/ssb-ngi-pointer/go-ssb-room/issues/66 is fixed
		a.Equal(http.StatusInternalServerError, resp.Code)

		// check database calls
		r.EqualValues(1, ts.InvitesDB.GetByTokenCallCount())
		_, tokenFromArg := ts.InvitesDB.GetByTokenArgsForCall(0)
		a.Equal(testToken, tokenFromArg)

		// fix #66
		// assertLocalized(t, doc, []localizedElement{
		// 	{"#welcome", "AuthFallbackWelcome"},
		// 	{"title", "AuthFallbackTitle"},
		// })
		gotErr := doc.Find("#errBody").Text()
		wantErr := weberrors.ErrNotFound{What: "invite"}
		a.EqualError(wantErr, gotErr)
	})

	wantNewMemberPlaceholder := "@                                            .ed25519"

	t.Run("token DOES exist", func(t *testing.T) {
		a, r := assert.New(t), require.New(t)

		testToken := "existing-test-token"
		validAcceptURL := urlTo(router.CompleteInviteFacade, "token", testToken)
		r.NotNil(validAcceptURL)

		// prep the mocked db for http:200
		fakeExistingInvite := roomdb.Invite{ID: 1234}
		ts.InvitesDB.GetByTokenReturns(fakeExistingInvite, nil)

		// request the form
		validAcceptForm := validAcceptURL.String()
		t.Log(validAcceptForm)
		doc, resp := ts.Client.GetHTML(validAcceptForm)
		a.Equal(http.StatusOK, resp.Code)

		// check database calls
		r.EqualValues(2, ts.InvitesDB.GetByTokenCallCount())
		_, tokenFromArg := ts.InvitesDB.GetByTokenArgsForCall(1)
		a.Equal(testToken, tokenFromArg)

		webassert.Localized(t, doc, []webassert.LocalizedElement{
			{"#welcome", "InviteFacadeWelcome"},
			{"title", "InviteFacadeTitle"},
		})

		form := doc.Find("form#consume")
		r.Equal(1, form.Length())

		webassert.CSRFTokenPresent(t, form)

		webassert.ElementsInForm(t, form, []webassert.FormElement{
			{Name: "invite", Type: "hidden", Value: testToken},
			{Name: "id", Type: "text", Placeholder: wantNewMemberPlaceholder},
		})
	})

	t.Run("token DOES exist but has no suggested alias", func(t *testing.T) {
		a, r := assert.New(t), require.New(t)

		testToken := "existing-test-token-2"
		validAcceptURL := urlTo(router.CompleteInviteFacade, "token", testToken)
		r.NotNil(validAcceptURL)

		testInvite := roomdb.Invite{ID: 4321}
		ts.InvitesDB.GetByTokenReturns(testInvite, nil)

		// request the form
		validAcceptForm := validAcceptURL.String()
		t.Log(validAcceptForm)
		doc, resp := ts.Client.GetHTML(validAcceptForm)
		a.Equal(http.StatusOK, resp.Code)

		// check database calls
		r.EqualValues(3, ts.InvitesDB.GetByTokenCallCount())
		_, tokenFromArg := ts.InvitesDB.GetByTokenArgsForCall(2)
		a.Equal(testToken, tokenFromArg)

		webassert.Localized(t, doc, []webassert.LocalizedElement{
			{"#welcome", "InviteFacadeWelcome"},
			{"title", "InviteFacadeTitle"},
		})

		form := doc.Find("form#consume")
		r.Equal(1, form.Length())

		webassert.CSRFTokenPresent(t, form)
		webassert.ElementsInForm(t, form, []webassert.FormElement{
			{Name: "invite", Type: "hidden", Value: testToken},
			{Name: "id", Type: "text", Placeholder: wantNewMemberPlaceholder},
		})
	})
}

func TestInviteConsumeInviteHTTP(t *testing.T) {
	ts := setup(t)
	a, r := assert.New(t), require.New(t)
	urlTo := web.NewURLTo(ts.Router)

	testToken := "existing-test-token-2"
	validAcceptURL := urlTo(router.CompleteInviteFacade, "token", testToken)
	r.NotNil(validAcceptURL)
	validAcceptURL.Host = "localhost"
	validAcceptURL.Scheme = "https"

	testInvite := roomdb.Invite{ID: 4321}
	ts.InvitesDB.GetByTokenReturns(testInvite, nil)

	// request the form (for a valid csrf token)
	validAcceptForm := validAcceptURL.String()
	t.Log(validAcceptForm)
	doc, resp := ts.Client.GetHTML(validAcceptForm)
	a.Equal(http.StatusOK, resp.Code)

	// we need a functional jar to unpack the Set-Cookie response for the csrf token
	jar, err := cookiejar.New(nil)
	r.NoError(err)

	// update the jar
	csrfCookie := resp.Result().Cookies()
	a.Len(csrfCookie, 1, "should have one cookie for CSRF protection validation")
	jar.SetCookies(validAcceptURL, csrfCookie)

	// get the corresponding token from the page
	csrfTokenElem := doc.Find("input[name='gorilla.csrf.Token']")
	a.Equal(1, csrfTokenElem.Length())
	csrfName, has := csrfTokenElem.Attr("name")
	a.True(has, "should have a name attribute")
	csrfValue, has := csrfTokenElem.Attr("value")
	a.True(has, "should have value attribute")

	// create the consume request
	testNewMember := refs.FeedRef{
		ID:   bytes.Repeat([]byte{1}, 32),
		Algo: refs.RefAlgoFeedSSB1,
	}
	consumeVals := url.Values{
		"invite": []string{testToken},
		"id":     []string{testNewMember.Ref()},

		csrfName: []string{csrfValue},
	}

	// construct the consume endpoint url
	consumeInviteURL, err := ts.Router.Get(router.CompleteInviteConsume).URL()
	r.Nil(err)
	consumeInviteURL.Host = "localhost"
	consumeInviteURL.Scheme = "https"

	// construct the header with Referer and Cookie
	var csrfCookieHeader = http.Header(map[string][]string{})
	csrfCookieHeader.Set("Referer", "https://localhost")
	cs := jar.Cookies(consumeInviteURL)
	r.Len(cs, 1, "expecting one cookie for csrf")
	theCookie := cs[0].String()
	a.NotEqual("", theCookie, "should have a new cookie")
	csrfCookieHeader.Set("Cookie", theCookie)
	ts.Client.SetHeaders(csrfCookieHeader)

	// prepare the mock
	ts.InvitesDB.ConsumeReturns(testInvite, nil)

	// send the POST
	resp = ts.Client.PostForm(consumeInviteURL.String(), consumeVals)
	a.Equal(http.StatusOK, resp.Code, "wrong HTTP status code for sign in")

	// check how consume was called
	r.EqualValues(1, ts.InvitesDB.ConsumeCallCount())
	_, tokenFromArg, newMemberRef := ts.InvitesDB.ConsumeArgsForCall(0)
	a.Equal(testToken, tokenFromArg)
	a.True(newMemberRef.Equal(&testNewMember))

	consumedDoc, err := goquery.NewDocumentFromReader(resp.Body)
	r.NoError(err)

	joinHref, ok := consumedDoc.Find("#join-link").Attr("href")
	a.True(ok)

	// validate ssb-uri
	joinURI, err := url.Parse(joinHref)
	r.NoError(err)

	a.Equal("ssb", joinURI.Scheme)
	a.Equal("experimental", joinURI.Opaque)

	params := joinURI.Query()
	a.Equal("join-room", params.Get("action"))

	gotRA := params.Get("multiserverAddress")

	a.True(strings.HasPrefix(gotRA, "net:localhost:8008~shs:"), "not for the test host: %s", gotRA)
	a.True(strings.HasSuffix(gotRA, base64.StdEncoding.EncodeToString(ts.NetworkInfo.RoomID.PubKey())), "public key missing? %s", gotRA)
}

func TestInviteConsumeInviteJSON(t *testing.T) {
	ts := setup(t)
	a, r := assert.New(t), require.New(t)
	urlTo := web.NewURLTo(ts.Router)

	testToken := "existing-test-token-2"
	validAcceptURL := urlTo(router.CompleteInviteFacade, "token", testToken)
	r.NotNil(validAcceptURL)

	testInvite := roomdb.Invite{ID: 4321}
	ts.InvitesDB.GetByTokenReturns(testInvite, nil)

	// create the consume request
	testNewMember := refs.FeedRef{
		ID:   bytes.Repeat([]byte{1}, 32),
		Algo: refs.RefAlgoFeedSSB1,
	}

	var consume inviteConsumePayload
	consume.Invite = testToken
	consume.ID = testNewMember

	// construct the consume endpoint url
	consumeInviteURL := urlTo(router.CompleteInviteConsume)
	r.NotNil(consumeInviteURL)

	// prepare the mock
	ts.InvitesDB.ConsumeReturns(testInvite, nil)

	// send the POST
	resp := ts.Client.SendJSON(consumeInviteURL.String(), consume)
	a.Equal(http.StatusOK, resp.Code, "wrong HTTP status code for sign in")

	// check how consume was called
	r.EqualValues(1, ts.InvitesDB.ConsumeCallCount())
	_, tokenFromArg, newMemberRef := ts.InvitesDB.ConsumeArgsForCall(0)
	a.Equal(testToken, tokenFromArg)
	a.True(newMemberRef.Equal(&testNewMember))

	var jsonConsumeResp inviteConsumeJSONResponse
	err := json.NewDecoder(resp.Body).Decode(&jsonConsumeResp)
	r.NoError(err)

	a.Equal("successful", jsonConsumeResp.Status)

	gotRA := jsonConsumeResp.RoomAddress
	a.True(strings.HasPrefix(gotRA, "net:localhost:8008~shs:"), "not for the test host: %s", gotRA)
	a.True(strings.HasSuffix(gotRA, base64.StdEncoding.EncodeToString(ts.NetworkInfo.RoomID.PubKey())), "public key missing? %s", gotRA)
}
