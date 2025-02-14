package apis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"codecosta.com/palico/app/models"
	"codecosta.com/palico/app/utils"
)

var CostaAPIBotToken string
var CostaAPIPort string

func CostaAPIGetRequest(endpoint string, discordUsername string) []byte {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://api:%s/palico%s", CostaAPIPort, endpoint), nil)
	req.Header = http.Header{
		"Authorization":      {"Bot " + CostaAPIBotToken},
		"X-Discord-Username": {discordUsername},
	}

	res, err := Client.Do(req)
	if err != nil {
		utils.LogAPIError(http.MethodGet, endpoint, err.Error())
		return []byte("ERROR")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		utils.LogSystemError("CostaAPIGetRequest", err.Error())
		return []byte("ERROR")
	}
	res.Body.Close()

	return body
}

// ---

/**
    These functions are used to interact with the Costa API.
	It is up to the caller to check for res.Error before using the response.
**/

func GetAPIMonsterInfo(monsterName string, discordUsername string) (models.MonsterResponse, error) {
	endpoint := fmt.Sprintf("/monster/%s", monsterName)

	var res models.MonsterResponse
	err := json.Unmarshal(CostaAPIGetRequest(endpoint, discordUsername), &res)
	if err != nil {
		utils.LogSystemError("GetMonsterInfo.Unmarshal", err.Error())
		return res, err
	}

	return res, nil
}
