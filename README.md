# Tesla Siparis Takibi

Telegram uzerinden guncelleme bildirimlerinin gelmesi icin bot olusturulmalidir.<br>
`.env` dosyasi icerisine eklenmelidir.
```
TELEGRAM_BOT_TOKEN=
TELEGRAM_CHAT_ID=
```

Uygulamayi isterseniz cron job olarak calistirabilirsiniz.

✅ setup_cron.sh
```bash
#!/bin/bash

export PATH="/opt/homebrew/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin"

# Proje yolu (gerekirse değiştir)
PROJECT_DIR=".../tesla-order-status"
EXECUTABLE="$PROJECT_DIR/tesla-order-status"
LOG_FILE="$PROJECT_DIR/logs/cron.log"

# 30dk'da bir calisacak
CRON_LINE="*/30 * * * * cd $PROJECT_DIR && $EXECUTABLE >> $LOG_FILE 2>&1"

mkdir -p "$PROJECT_DIR/logs"

echo "> Building Go binary..."
cd "$PROJECT_DIR" || exit 1
go build -o tesla-order-status ./cmd/main.go

echo "> Installing cron job..."
(crontab -l 2>/dev/null | grep -v "$EXECUTABLE"; echo "$CRON_LINE") | crontab -

echo "> Cron job added:"
echo "$CRON_LINE"
```

Sonra proje dizininde terminal'den asagidaki komutlari sirayla calistirin.
```bash
chmod +x setup_cron.sh

./setup_cron.sh
```

Uygulamayi ilk calistirdiginizda access token bilgisine sahip olmadigi icin tarayici uzerinden yonlendirme adresini isteyecek.
```console
> Retrieving Tesla order information...

Tarayıcıda oturum açın ve yönlendirme URL'sini buraya yapıştırın:
https://auth.tesla.com/oauth2/v3/authorize?client_id=ownerapi&code_challenge=...

Yönlendirme URL'sini girin: https://auth.tesla.com/void/callback?code=...
```

Yonlendirme adresini girdikten sonra token bilgilerini girdikten sonra kaydedip isteyip istemediginizi soracak.
```console
Save current order data for future diff? (y/n): 
```

Buna "Y" cevabini verdikten sonra siparisinize ait bilgileri cekecektir.

