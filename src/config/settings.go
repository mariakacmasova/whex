package config

import (
	"fmt"
	waProto "go.mau.fi/whatsmeow/binary/proto"
)

var (
	AppVersion             = "v4.0.0"
	AppPort                = "4001"
	AppDebug               = true
	AppOs                  = fmt.Sprintf("AldinoKemal")
	AppPlatform            = waProto.DeviceProps_PlatformType(1)
	AppBasicAuthCredential string

	PathQrCode    = "statics/qrcode"
	PathSendItems = "statics/senditems"
	PathStorages  = "storages"

	DBName = "whatsapp.db"

	WhatsappLogLevel            = "ERROR"
	WhatsappAutoReplyMessage    string
	WhatsappAutoReplyWebhook    = "https://n8n-production-e722.up.railway.app/webhook-test/whauto"
	WhatsappSettingMaxFileSize  int64 = 50000000  // 50MB
	WhatsappSettingMaxVideoSize int64 = 100000000 // 100MB
)
