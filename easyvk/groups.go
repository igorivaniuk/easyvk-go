package easyvk

import (
	"encoding/json"
	"strconv"
	"strings"
)

type Groups struct {
	vk *VK
}

// https://vk.com/dev/groups.getById
type GroupsByIdResponse []GroupObject

// GetById returns information about communities by their IDs.
// https://vk.com/dev/groups.getById
func (g *Groups) GetById(groupIds []int, fields []string) (GroupsByIdResponse, error) {
	params := map[string]string{
		"group_ids": strings.Join(intIdsToString(groupIds), ","),
		"fields":    strings.Join(fields, ","),
	}
	resp, err := g.vk.Request("groups.getById", params)
	if err != nil {
		return nil, err
	}
	var groups GroupsByIdResponse
	err = json.Unmarshal(resp, &groups)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

// https://vk.com/dev/groups.isMember
type IsMembersResponse []IsMember

type IsMember struct {
	UserId     int `json:"user_id"`
	Member     int `json:"member"`
	Request    int `json:"request"`
	Invitation int `json:"invitation"`
}

// https://vk.com/dev/groups.isMember
func (g *Groups) IsMembers(groupId int, userIds []int) (IsMembersResponse, error) {
	params := map[string]string{
		"group_id": strconv.Itoa(groupId),
		"user_ids": strings.Join(intIdsToString(userIds), ","),
		"extended": "1",
	}
	resp, err := g.vk.Request("groups.isMember", params)
	if err != nil {
		return nil, err
	}
	var res IsMembersResponse
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// https://vk.com/dev/groups.isMember
func (g *Groups) IsMember(groupId int, userId int) (*IsMember, error) {
	params := map[string]string{
		"group_id": strconv.Itoa(groupId),
		"user_id":  strconv.Itoa(userId),
		"extended": "1",
	}
	resp, err := g.vk.Request("groups.isMember", params)
	if err != nil {
		return nil, err
	}
	res := &IsMember{}
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type GetMembersIdsParams struct {
	GroupId int
	Sort string
	Offset int
	Count int
}

type GetMembersIdsResponse struct {
	Count int `json:"count"`
	Items []int `json:"items"`
}

// https://vk.com/dev/groups.getMembers
func (g *Groups) GetMembersIds(p GetMembersIdsParams) (*GetMembersIdsResponse, error) {
	// set default count
	count := 100
	if p.Count != 0 {
		count = p.Count
	}
	params := map[string]string{
		"group_id": strconv.Itoa(p.GroupId),
		"sort":  p.Sort,
		"offset":  strconv.Itoa(p.Offset),
		"count":  strconv.Itoa(count),
	}
	resp, err := g.vk.Request("groups.getMembers", params)
	if err != nil {
		return nil, err
	}
	res := &GetMembersIdsResponse{}
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type GetMembersInfoParams struct {
	GroupId int
	Sort string
	Offset int
	Count int
	Fields string
	Filter string
}

type GetMembersInfoResponse struct {
	Count int `json:"count"`
	Items []UserObject `json:"items"`
}

// https://vk.com/dev/groups.getMembers
func (g *Groups) GetMembersInfo(p GetMembersInfoParams) (*GetMembersInfoResponse, error) {
	// set default count
	count := 100
	if p.Count != 0 {
		count = p.Count
	}
	// set field for return user object not id
	if p.Fields == "" {
		p.Fields = "photo_50"
	}
	params := map[string]string{
		"group_id": strconv.Itoa(p.GroupId),
		"sort":  p.Sort,
		"offset":  strconv.Itoa(p.Offset),
		"count":  strconv.Itoa(count),
		"fields":  p.Fields,
		"filter":  p.Filter,
	}
	resp, err := g.vk.Request("groups.getMembers", params)
	if err != nil {
		return nil, err
	}
	res := &GetMembersInfoResponse{}
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// String with the confirmation code.
// https://vk.com/dev/groups.getCallbackConfirmationCode
type GetCallbackConfirmationCodeResponse struct {
	Code string `json:"code"`
}

// GetCallbackConfirmationCode returns Callback API confirmation code for the community.
// https://vk.com/dev/groups.getCallbackConfirmationCode
func (g *Groups) GetCallbackConfirmationCode(groupId int) (*GetCallbackConfirmationCodeResponse, error) {
	params := map[string]string{
		"group_id": strconv.Itoa(groupId),
	}
	resp, err := g.vk.Request("groups.getCallbackConfirmationCode", params)

	if err != nil {
		return nil, err
	}
	res := &GetCallbackConfirmationCodeResponse{}
	err = json.Unmarshal(resp, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// https://vk.com/dev/groups.getCallbackServers
type GetCallbackServersResponse struct {
	Count int              `json:"count"`
	Items []CallbackServer `json:"items"`
}

type CallbackServer struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	CreatorId int    `json:"creator_id"`
	Url       string `json:"url"`
	SecretKey string `json:"secret_key"`
	/*
		unconfigured — адрес не задан;
		fail — подтвердить адрес не удалось;
		wait — адрес ожидает подтверждения;
		ok — сервер подключен.
	*/
	Status string `json:"status"`
}

// https://vk.com/dev/groups.getCallbackServers
func (g *Groups) GetCallbackServers(groupId int, serverIds []int) (*GetCallbackServersResponse, error) {
	params := map[string]string{
		"group_id":   strconv.Itoa(groupId),
		"server_ids": strings.Join(intIdsToString(serverIds), ","),
	}
	resp, err := g.vk.Request("groups.getCallbackServers", params)

	if err != nil {
		return nil, err
	}
	res := &GetCallbackServersResponse{}
	err = json.Unmarshal(resp, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// https://vk.com/dev/groups.addCallbackServer
type AddCallbackServerResponse struct {
	ServerId int `json:"server_id"`
}

// https://vk.com/dev/groups.addCallbackServer
func (g *Groups) AddCallbackServer(groupId int, url, title, secretKey string) (*AddCallbackServerResponse, error) {
	params := map[string]string{
		"group_id":   strconv.Itoa(groupId),
		"url":        url,
		"title":      title,
		"secret_key": secretKey,
	}
	resp, err := g.vk.Request("groups.addCallbackServer", params)

	if err != nil {
		return nil, err
	}

	res := &AddCallbackServerResponse{}
	err = json.Unmarshal(resp, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// https://vk.com/dev/groups.editCallbackServer
type EditCallbackServerResponse int

// https://vk.com/dev/groups.editCallbackServer
func (g *Groups) EditCallbackServer(groupId, serverId int, url, title, secretKey string) (EditCallbackServerResponse, error) {
	params := map[string]string{
		"group_id":   strconv.Itoa(groupId),
		"server_id":  strconv.Itoa(serverId),
		"url":        url,
		"title":      title,
		"secret_key": secretKey,
	}
	resp, err := g.vk.Request("groups.editCallbackServer", params)

	var res EditCallbackServerResponse
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// https://vk.com/dev/groups.setCallbackSettings
type SetCallbackSettingsResponse int

var knowEvents = []string{
	"message_new",
	"message_reply",
	"message_allow",
	"message_deny",
	"photo_new",
	"audio_new",
	"video_new",
	"wall_reply_new",
	"wall_reply_edit",
	"wall_reply_delete",
	"wall_reply_restore",
	"wall_post_new",
	"wall_repost",
	"board_post_new",
	"board_post_edit",
	"board_post_restore",
	"board_post_delete",
	"photo_comment_new",
	"photo_comment_edit",
	"photo_comment_delete",
	"photo_comment_restore",
	"video_comment_new",
	"video_comment_edit",
	"video_comment_delete",
	"video_comment_restore",
	"market_comment_new",
	"market_comment_edit",
	"market_comment_delete",
	"market_comment_restore",
	"poll_vote_new",
	"group_join",
	"group_leave",
}

// https://vk.com/dev/groups.setCallbackSettings
func (g *Groups) SetCallbackSettings(groupId, serverId int, enableEvents []string) (SetCallbackSettingsResponse, error) {
	params := map[string]string{
		"group_id":  strconv.Itoa(groupId),
		"server_id": strconv.Itoa(serverId),
	}

	var enabled = func(event string) bool {
		for _, ev := range enableEvents {
			if ev == event {
				return true
			}
		}
		return false
	}

	for _, event := range knowEvents {
		if enabled(event) {
			params[event] = "1"
		} else {
			params[event] = "0"
		}
	}

	resp, err := g.vk.Request("groups.setCallbackSettings", params)
	var res SetCallbackSettingsResponse
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
