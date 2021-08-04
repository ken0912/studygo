package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"math"
	"math/rand"
	"time"
	"unicode/utf8"

	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tts "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tts/v20190823"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randString(length int32) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Int63()%int64(len(charset))]
	}
	return string(b)
}

type VoiceConfig struct {
	Endpoint  string
	Region    string
	VoiceCode int64
}

var (
	DefaultBoyVoice = VoiceConfig{
		Endpoint:  "tts.tencentcloudapi.com",
		Region:    "ap-shanghai",
		VoiceCode: 101015,
	}
	DefaultGirlVoice = VoiceConfig{
		Endpoint:  "tts.tencentcloudapi.com",
		Region:    "ap-shanghai",
		VoiceCode: 101016,
	}
)

func GetVoiceFromTencentCloud(SecretID string, SecretKey string, voice VoiceConfig, message string) (string, error) {
	credential := common.NewCredential(
		SecretID,
		SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = voice.Endpoint
	client, _ := tts.NewClient(credential, voice.Region, cpf)

	request := tts.NewTextToVoiceRequest()
	request.SessionId = common.StringPtr(randString(68))
	request.Text = common.StringPtr(message)
	request.ModelType = common.Int64Ptr(1)
	request.VoiceType = common.Int64Ptr(voice.VoiceCode)
	request.Volume = common.Float64Ptr(10)
	request.Speed = common.Float64Ptr(math.Min(2.0, float64(utf8.RuneCountInString(message)/8)))
	response, err := client.TextToVoice(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return "", err
	}

	//base64编码的wav/mp3音频数据
	return *response.Response.Audio, nil
}
func main() {

	SecretID := "AKIDQghRv96nFSO3RZQ0KOBiyxf5F9m8D1at"
	SecretKey := "xhBOKzNVMwGJ6csra2EAqGOGsDI6ewyj"
	message := `鲁迅听罢，愤然地说人疯了，天也疯了。正在这时，表弟阮久荪闯进来，嚷嚷着外面有人要杀他。后来经人解释才知道，阮久荪在来的路上，看到了很多饿死的人，遭受刺激导致神经错乱。鲁迅目睹这一景象，想到杨开铭的失心发疯，还有菜市口行刑场面那蘸着人血的馒头……便凝神构思，随即提笔撰稿，创作出他首篇白话文小说《狂人日记》。`
	encodeVoice, err := GetVoiceFromTencentCloud(SecretID, SecretKey, DefaultGirlVoice, message)
	if err != nil {
		fmt.Println("GetVoiceFromTencentCloud error:", err)
	}

	data, err := base64.StdEncoding.DecodeString(encodeVoice)

	streamer, format, err := wav.Decode(bytes.NewReader(data))
	fmt.Println("streamer:", streamer)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/30))
	speaker.Play(streamer)

	for {
	}
}
