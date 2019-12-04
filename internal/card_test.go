package server_test

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"iKnowThisWord/internal/model"
	. "iKnowThisWord/internal/store"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestServer_HandleCardFind(t *testing.T) {
	err := seedCards(3)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodGet, "/api/cards", nil)
	if err != nil {
		t.Error(err)
	}
	rr := httptest.NewRecorder()
	handler := s.HandleCardFind()
	handler.ServeHTTP(rr, req)

	var cards []model.Card
	err = json.Unmarshal([]byte(rr.Body.String()), &cards)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(cards), 3)
}

func TestServer_HandleCardSave(t *testing.T) {
	err := refreshCardTable()
	if err != nil {
		t.Fatal(err)
	}

	type saveCardTest struct {
		inputJSON string
		code      int
	}

	saveCardTests := []*saveCardTest{
		{`{ "word": "consider", "meaning": "deem to be"}`, 201},
		{`{ "word": "consider", "meaning": ""}`, 400},
		{`{ "word": "", "meaning": "deem to be"}`, 400},
		{`{}`, 400},
	}

	for _, tc := range saveCardTests {
		req, err := http.NewRequest(http.MethodPost, "/api/cards", bytes.NewBufferString(tc.inputJSON))
		if err != nil {
			t.Error(err)
		}

		rr := httptest.NewRecorder()
		handler := s.HandleCardSave()
		handler.ServeHTTP(rr, req)

		assert.Equal(t, tc.code, rr.Code)

		if tc.code == 201 {
			idStr := strings.TrimSuffix(rr.Body.String(), "\n")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				t.Fatal(err)
			}
			assert.Greater(t, id, 0)
		}

		if tc.code == 400 {
			decodedMap := make(map[string]interface{})
			err = json.Unmarshal([]byte(rr.Body.String()), &decodedMap)
			if err != nil {
				t.Fatal(err)
			}

			assert.NotEmpty(t, decodedMap["error"])
		}
	}
}

func TestServer_HandleCardDelete(t *testing.T) {
	id, err := seedCard()
	if err != nil {
		t.Fatal(err)
	}

	type deleteCardTest struct {
		id   int
		code int
	}

	deleteCardTests := []*deleteCardTest{
		{id, http.StatusOK},
		{id - 1, http.StatusNotFound},
	}

	for _, tc := range deleteCardTests {
		req, err := http.NewRequest(http.MethodDelete, "/api/cards", nil)
		if err != nil {
			t.Fatal(err)
		}

		req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(id)})
		rr := httptest.NewRecorder()

		handler := s.HandleCardDelete()
		handler.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, tc.code)
	}

}

func TestServer_HandleCardFindByID(t *testing.T) {
	err := refreshCardTable()
	if err != nil {
		t.Fatal(err)
	}

	c := testCard(1)

	_, err = store.Card().FindByID(strconv.Itoa(c.ID))
	assert.EqualError(t, err, ErrRecordNotFound.Error())

	err = store.Card().Save(c)
	if err != nil {
		t.Fatal(err)
	}

	result, err := store.Card().FindByID(strconv.Itoa(c.ID))
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, result.Word, c.Word)
}
