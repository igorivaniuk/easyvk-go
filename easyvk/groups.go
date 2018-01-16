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
func (g *Groups) GetById(groupIds []int, fields string) (GroupsByIdResponse, error) {
	var groupIdsStr []string
	for _, gid := range groupIds {
		text := strconv.Itoa(gid)
		groupIdsStr = append(groupIdsStr, text)
	}
	params := map[string]string{
		"group_ids": strings.Join(groupIdsStr, ","),
		"fields":    fields,
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
	var serverIdsStr []string
	for _, gid := range serverIds {
		text := strconv.Itoa(gid)
		serverIdsStr = append(serverIdsStr, text)
	}
	params := map[string]string{
		"group_id":   strconv.Itoa(groupId),
		"server_ids": strings.Join(serverIdsStr, ","),
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
