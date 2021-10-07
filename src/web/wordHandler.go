package web

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"story_writer/src/constant"
	"story_writer/src/model"

	"github.com/gorilla/mux"
)

func (web *Web) AddWord(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()

	var wordInput *model.Word

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("Request Error request read error")
	}

	if err = json.Unmarshal(body, &wordInput); err != nil {
		return nil, errors.New("Request body error unmarshal error")
	}

	if len(strings.Fields(wordInput.Word)) > 1 {
		return nil, errors.New("multiple words sent")
	}

	if len(wordInput.Word) == 0 {
		return nil, errors.New("No word sent")
	}

	response, err := web.managerModule.AddWord(ctx, wordInput.Word)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (web *Web) GetStories(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()

	var limit, offset int
	var sortBy, order string

	limit, err := strconv.Atoi(r.FormValue("limit"))
	if err != nil || limit > constant.LIMIT || limit < 1 {
		limit = constant.LIMIT
	}

	offset, err = strconv.Atoi(r.FormValue("offset"))
	if err != nil {
		offset = 0
	}

	sortBy = r.FormValue("sort")
	if sortBy == "" {
		sortBy = constant.DEFAULT_SORT_BY
	}

	order = r.FormValue("order")
	if order == "" {
		order = constant.DEFAULT_ORDER
	}

	params := model.Params{Limit: limit, Offset: offset, SortBy: sortBy, Order: order}
	response, err := web.managerModule.GetStories(ctx, params)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (web *Web) GetStoryById(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()

	vars := mux.Vars(r)

	storyId, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Println("[GetStoryById] story id missing")
		return nil, errors.New("story id is missing")
	}

	storyResp, err := web.managerModule.GetStoryById(ctx, storyId)
	if err != nil {
		return storyResp, err
	}

	return storyResp, nil
}
