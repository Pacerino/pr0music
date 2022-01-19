package pr0gramm

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
)

type Session struct {
	client http.Client
}

func NewSession(client http.Client) *Session {
	client.Jar, _ = cookiejar.New(nil)
	return &Session{client: client}
}

func (sess *Session) Login(username, password string) (*LoginResponse, error) {
	body := make(url.Values)
	body.Set("name", username)
	body.Set("password", password)

	uri := "https://pr0gramm.com/api/user/login"
	resp, err := sess.client.PostForm(uri, body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response LoginResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	return &response, err
}

func (sess *Session) apiGET(path string, query url.Values, target interface{}) error {
	uri := "https://pr0gramm.com/api" + path

	if query != nil {
		uri += "?" + query.Encode()
	}

	response, err := sess.client.Get(uri)
	if err != nil {
		return err
	}

	defer func() {
		_, _ = io.Copy(ioutil.Discard, response.Body)
		_ = response.Body.Close()
	}()

	if response.StatusCode/100 != 2 {
		return fmt.Errorf("error %d", response.StatusCode)
	}
	return json.NewDecoder(response.Body).Decode(target)
}

func (sess *Session) GetUserComments(user string, flags int, after int) (CommentResponse, error) {
	query := make(url.Values)
	query.Set("name", user)
	query.Set("flags", strconv.Itoa(flags))
	if after != 0 {
		query.Set("after", strconv.Itoa(after))
	}

	var response CommentResponse
	err := sess.apiGET("/profile/comments", query, &response)
	return response, err
}

func (sess *Session) GetComments() (MessagesResponse, error) {
	var response MessagesResponse
	err := sess.apiGET("/inbox/comments", nil, &response)
	return response, err
}
