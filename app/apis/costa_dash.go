package apis

import (
	"fmt"
	"io"
	"net/http"

	"codecosta.com/palico/app/utils"
)

var CostaAPIBotToken string
var CostaAPIPort string

func CostaAPIGetRequest(endpoint string, discordUsername string) []byte {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://api:%s/bot%s", CostaAPIPort, endpoint), nil)
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

func GetMonsterInfo(monsterName string) []byte {
	return CostaAPIGetRequest(fmt.Sprintf("/monster/%s", monsterName), "")
}
