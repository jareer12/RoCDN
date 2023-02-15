package rblx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rblx/structs"
)

func GetHeadshot(user_id int, size int, format string, circle bool) (structs.Image, error) {
	if req, err := http.NewRequest("GET", fmt.Sprintf("https://thumbnails.roblox.com/v1/users/avatar-headshot?userIds=%v&size=%vx%v&format=%v&isCircular=%v", user_id, size, size, format, circle), nil); err != nil {
		return structs.Image{}, err
	} else {
		if res, err := http.DefaultClient.Do(req); err != nil {
			return structs.Image{}, err
		} else {
			if body, err := ioutil.ReadAll(res.Body); err != nil {
				return structs.Image{}, err
			} else {
				var i structs.RobloxResponse
				fmt.Println(string(body))

				if err := json.Unmarshal(body, &i); err != nil {
					return structs.Image{}, err
				} else {
					if len(i.Data) > 0 {
						return i.Data[0], nil
					} else {
						return structs.Image{}, err
					}
				}
			}
		}
	}
}
