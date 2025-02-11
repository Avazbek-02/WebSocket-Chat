# WebSocket-Chat

## Loyihaga umumiy tavsif

WebSocket-Chat - bu real vaqt rejimida foydalanuvchilar o'zaro xabar almashishi mumkin bo'lgan WebSocket asosidagi chat dasturi. WebSocket protokoli orqali xabarlarni tez va kechikishsiz uzatish imkoniyati mavjud.

## Texnologiyalar

- **Backend**: Go (Gorilla WebSocket kutubxonasi)
- **Frontend**: HTML, CSS, JavaScript

## O‘rnatish va ishga tushirish

### Talablar

- Go dasturlash tili o‘rnatilgan bo‘lishi kerak
- Internet yoki lokal tarmoqda ishlaydigan Wi-Fi

### 1. Loyihani klonlash

```sh
    git clone https://github.com/username/WebSocket-Chat.git
    cd WebSocket-Chat
```

### 2. Zarur kutubxonalarni o‘rnatish

```sh
    go mod tidy
```

### 3. Serverni ishga tushirish

```sh
    go run main.go
```

Server `http://10.10.1.189:9999/` manzilida ishga tushadi.(Buni siz uzizning ip manzilizga uzgartiras)

## Konfiguratsiya

Agar server boshqa IP yoki portda ishlashi kerak bo‘lsa, `main.go` faylida quyidagi qatorni o‘zgartiring:

```go
    http.ListenAndServe("0.0.0.0:9999", nil)
```

9999 portni kerakli portga o‘zgartirish mumkin.

## Ishlatish bo‘yicha qo‘llanma

### 1. Chatga kirish

Brauzerdan `http://10.10.1.189:9999/` manziliga kiring.

### 2. Foydalanuvchi nomini kiritish

Sayt ochilganda ism kiritish so‘raladi. Agar kiritilmasa, `Anonim` nomi beriladi.

### 3. Xabar yuborish

Matn maydoniga xabar yozing va **Yuborish** tugmasini bosing yoki `Enter` tugmasini bosing.

## Muammolarni hal qilish

Agar WebSocket ulanmasa:

- Server IP manzilini to‘g‘ri kiritganingizni tekshiring (`index.html` faylida `serverIP` o‘zgaruvchisini moslang).
- Portni firewall yoki boshqa dasturlar bloklamayotganiga ishonch hosil qiling.
- Brauzer konsolida (`F12 > Console`) xatoliklar borligini tekshiring.

## Muallif

- **Avazbek** - Backend dasturchi

## Litsenziya

Bu loyiha **MIT** litsenziyasi asosida tarqatiladi.

