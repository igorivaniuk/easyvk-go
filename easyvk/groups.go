package easyvk

import (
	"encoding/json"
	"strings"
)

type Groups struct {
	vk *VK
}

// https://vk.com/dev/groups.getById
type GroupsByIdResponse []GroupObject

// GetById returns information about communities by their IDs.
// https://vk.com/dev/groups.getById
func (g *Groups) GetById(groupIds []string, fields string) (GroupsByIdResponse, error) {
	params := map[string]string{
		"group_ids": strings.Join(groupIds, ","),
		"fields":   fields,
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
func (g *Groups) GetCallbackConfirmationCode(groupId string) (GetCallbackConfirmationCodeResponse, error) {
	params := map[string]string{
		"group_id": groupId,
	}
	resp, err := g.vk.Request("groups.getCallbackConfirmationCode", params)
	if err != nil {
		return GetCallbackConfirmationCodeResponse{}, err
	}
	var codeRes GetCallbackConfirmationCodeResponse
	err = json.Unmarshal(resp, &codeRes)
	if err != nil {
		return GetCallbackConfirmationCodeResponse{}, err
	}

	return codeRes, nil
}
