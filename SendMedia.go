package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

// SendMedia sends a media file to a recipient
func SendMedia(botToken string, chatId string, path string) (*Message, error) {
	mtype, err := GetFileContentType(path)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	endPoint := ""
	fileType := ""

	if strings.Contains(mtype, "image") {
		endPoint = "sendPhoto"
		fileType = "photo"
	} else if strings.Contains(mtype, "audio") {
		endPoint = "sendAudio"
		fileType = "audio"
	} else if strings.Contains(mtype, "video") {
		endPoint = "sendVideo"
		fileType = "video"
	} else {
		endPoint = "sendDocument"
		fileType = "document"
	}

	values := map[string]io.Reader{
		fileType:  file,
		"chat_id": strings.NewReader(chatId),
	}

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return nil, err
			}
		} else {
			if fw, err = w.CreateFormField(key); err != nil {
				return nil, err
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return nil, err
		}
	}
	w.Close()

	response, err := http.Post("https://api.telegram.org/"+botToken+"/"+endPoint, w.FormDataContentType(), &b)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		j, err := json.Marshal(response.Body)
		if err != nil {
			return nil, err
		}
		log.Println(string(j))
		return nil, errors.New(endPoint + " request returned: " + response.Status)
	}
	var result *ResponseBody
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, errors.New("could not convert response to *ResponseBody")
	}
	return result.Result, nil
}
