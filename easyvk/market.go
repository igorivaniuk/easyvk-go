package easyvk

import (
	"fmt"
	"strconv"
)

// https://vk.com/dev/market
type Market struct {
	vk *VK
}

// https://vk.com/dev/market.deleteComment
func (m *Market) DeleteComment(ownerID, commentId int) (bool, error) {
	params := map[string]string{
		"owner_id":   fmt.Sprint(ownerID),
		"comment_id": fmt.Sprint(commentId),
	}

	resp, err := m.vk.Request("market.deleteComment", params)
	if err != nil {
		return false, err
	}

	ok, err := strconv.ParseUint(string(resp), 10, 8)
	if err != nil {
		return false, err
	}
	return ok == 1, nil
}
