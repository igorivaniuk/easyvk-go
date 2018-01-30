package easyvk

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// A Wall describes a set of methods
// to work with wall.
// https://vk.com/dev/wall
type Wall struct {
	vk *VK
}

// WallPostParams provides structure for post method.
// https://vk.com/dev/wall.post
type WallPostParams struct {
	OwnerID            int
	FriendsOnly        bool
	FromGroup          bool
	Signed             bool
	MarkAsAds          bool
	AdsPromotedStealth bool
	Message            string
	Attachments        string
	Services           string
	GUID               string
	PublishDate        uint
	PlaceID            uint
	PostID             uint
	Lat                float64
	Long               float64
}

// Post adds a new post on a user wall or community wall.
// Can also be used to publish suggested or scheduled posts.
// Returns id of created post.
// https://vk.com/dev/wall.post
func (w *Wall) Post(p WallPostParams) (int, error) {

	params := map[string]string{
		"owner_id":             fmt.Sprint(p.OwnerID),
		"message":              p.Message,
		"attachments":          p.Attachments,
		"services":             p.Services,
		"guid":                 p.GUID,
		"publish_date":         fmt.Sprint(p.PublishDate),
		"place_id":             fmt.Sprint(p.PlaceID),
		"post_id":              fmt.Sprint(p.PostID),
		"lat":                  fmt.Sprint(p.Lat),
		"long":                 fmt.Sprint(p.Long),
		"friends_only":         boolConverter(p.FriendsOnly),
		"from_group":           boolConverter(p.FromGroup),
		"signed":               boolConverter(p.Signed),
		"mark_as_ads":          boolConverter(p.MarkAsAds),
		"ads_promoted_stealth": boolConverter(p.AdsPromotedStealth),
	}

	resp, err := w.vk.Request("wall.post", params)
	if err != nil {
		return 0, err
	}
	var info struct {
		PostID int `json:"post_id"`
	}

	err = json.Unmarshal(resp, &info)
	if err != nil {
		return 0, err
	}
	return info.PostID, nil
}

// https://vk.com/dev/wall.deleteComment
func (w *Wall) DeleteComment(ownerID, commentId int) (bool, error) {
	params := map[string]string{
		"owner_id":   fmt.Sprint(ownerID),
		"comment_id": fmt.Sprint(commentId),
	}

	resp, err := w.vk.Request("wall.deleteComment", params)
	if err != nil {
		return false, err
	}

	ok, err := strconv.ParseUint(string(resp), 10, 8)
	if err != nil {
		return false, err
	}
	return ok == 1, nil
}

type CreateCommentParams struct {
	OwnerID        int
	PostID         int
	FromGroup      int
	Message        string
	ReplyToComment int
	Attachments    string
	StickerID      int
	GUID           string
}

// https://vk.com/dev/wall.createComment
func (w *Wall) CreateComment(p CreateCommentParams) (int, error) {

	params := map[string]string{
		"owner_id":         fmt.Sprint(p.OwnerID),
		"message":          p.Message,
		"attachments":      p.Attachments,
		"post_id":          fmt.Sprint(p.PostID),
		"guid":             p.GUID,
		"from_group":       fmt.Sprint(p.FromGroup),
		"sticker_id":       fmt.Sprint(p.StickerID),
		"reply_to_comment": fmt.Sprint(p.ReplyToComment),
	}

	resp, err := w.vk.Request("wall.createComment", params)
	if err != nil {
		return 0, err
	}

	postId, err := strconv.Atoi(string(resp))
	if err != nil {
		return 0, err
	}
	return postId, nil
}
