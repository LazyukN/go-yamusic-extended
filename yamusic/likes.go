package yamusic

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type (
	LikesService struct {
		client *Client
	}

	LikeResp struct {
		InvocationInfo InvocationInfo `json:"invocationInfo"`
		Error          Error          `json:"error"`
		Result         struct {
			Revision int `json: revision`
		}
	}
)

func (s *LikesService) Like(
	ctx context.Context,
	entity string,
	ids []string,
) (*LikeResp, *http.Response, error) {
	form := url.Values{}
	form.Set(entity+"Ids", strings.Join(ids, ","))

	uri := fmt.Sprintf("users/%v/likes/%vs/add-multiple", s.client.userID, entity)

	req, err := s.client.NewRequest(http.MethodPost, uri, form)
	if err != nil {
		return nil, nil, err
	}

	like := new(LikeResp)
	resp, err := s.client.Do(ctx, req, like)
	return like, resp, err

}

func (s *LikesService) Dislike(
	ctx context.Context,
	entity string,
	ids []string,
) (*LikeResp, *http.Response, error) {
	form := url.Values{}
	form.Set(entity+"Ids", strings.Join(ids, ","))

	uri := fmt.Sprintf("users/%v/likes/%vs/remove", s.client.userID, entity)

	req, err := s.client.NewRequest(http.MethodPost, uri, form)
	if err != nil {
		return nil, nil, err
	}

	like := new(LikeResp)
	resp, err := s.client.Do(ctx, req, like)
	return like, resp, err

}
