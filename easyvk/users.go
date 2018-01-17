package easyvk

import (
	"encoding/json"
	"strings"
)

type Users struct {
	vk *VK
}

type UsersGetResponse []UserObject

// https://vk.com/dev/users.get
/*
nameCase: Case for declension of user name and surname:
nom — nominative (default)
gen — genitive
dat — dative
acc — accusative
ins — instrumental
abl — prepositional
*/
func (u *Users) Get(userIds []int, fields []string, nameCase string) (UsersGetResponse, error) {
	params := map[string]string{}
	if len(userIds) > 0 {
		params["user_ids"] = strings.Join(intIdsToString(userIds), ",")
	}
	if len(fields) > 0 {
		params["fields"] = strings.Join(fields, ",")
	}
	if nameCase != "" {
		params["name_case"] = nameCase
	}
	resp, err := u.vk.Request("users.get", params)
	if err != nil {
		return nil, err
	}
	var users UsersGetResponse
	err = json.Unmarshal(resp, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
