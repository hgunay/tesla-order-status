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

Bu noktada cron job icerisinde interaktif olarak token alamayacaginiz icin uygulamayi bir kereligine terminal icerisinde calistirip url'i buradan ekleyin.
```console
go run cmd/main.go
```

Uygulamayi calistiginda access token bilgisine sahip olmadigi icin tarayici uzerinden yonlendirme adresini isteyecek.
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

Bilgiler asagidaki gibi gozukecek.
```console
---------------------------------------------
               GÜNCELLEMELER
---------------------------------------------
- Kilometre: 0.8
+ Kilometre: 0.9173260799999999

---------------------------------------------
             ARAÇ BİLGİLERİ
---------------------------------------------
- Model Kodu: MY
- Araç Tipi: Yeni
- VIN: XP........
- Sipariş Durumu: Rezerve Edildi
- Sipariş Alt Durumu: _Z
- Seri: General Production
- Trim Kodu: $....
- Araç Konum Kodu: 4....
- Konfigürasyon ID: 1..>...
- Üretim Yılı: 2025
- Sipariş Tarihi: 09.05.2025 18:45
- Kilometre: 0.92 KM

---------------------------------------------
           MÜŞTERİ BİLGİLERİ
---------------------------------------------
- İsim: ......
- Soyisim: ......
- E-posta: ......
- Telefon: ......

Adres Bilgileri:
- Adres: ......
- Şehir: ......
- İl: ......
- Posta Kodu: ......

---------------------------------------------
             ÖDEME BİLGİLERİ
---------------------------------------------
- Toplam Tutar: ...... TL
- Ödenen Tutar: ...... TL
- Ön Ödeme Tutarı: ...... TL

---------------------------------------------
                KAYIT DURUMU
---------------------------------------------
- Kayıt Durumu: (bilgi yok)
- Kayıt Tarihi: 09.05.2025 18:59

---------------------------------------------
                  TESLİMAT
---------------------------------------------
- Tahmini Teslimat Aralığı: 28 May - 11 June
- Tahmini Varış Tarihi: 28.05.2025 03:00
- Teslimat Randevusu: -
- Teslimat Yeri: -

---------------------------------------------
                GÖREV DURUMU
---------------------------------------------
✗ Son Ödeme
✓ Kayıt
✓ Teslimat
✗ Finansman
✓ Servis Ziyareti
✗ Sigorta
✗ Zamanlama
✗ Sipariş Onayı
✓ Fabrika Teslim
✗ Hazır Ürün
✗ Son Fatura
✗ Hazırlık
✗ Takas
✓ Gecikme Engeli

---------------------------------------------
             SİGORTA BİLGİSİ
---------------------------------------------
- Sigorta Durumu: IGNORE
- Sigorta Şirketi: 
---------------------------------------------


> Sending message to Telegram bot...
📦 Tesla sipariş bilgilerinde güncelleme var:

➖ Kilometre: 0.8
➕ Kilometre: 0.9173260799999999
```

Telegram mesaj ornegi:<br>
![image](https://github.com/user-attachments/assets/9d60e29f-bead-43c9-8c20-0b0e4aa667ce)

