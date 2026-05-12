# Доработки авторизации

## Проблема

Авторизация работала через хардкоженный код `12345` — любой пользователь мог войти, зная этот код. Код проверялся в `pkg/code/none/none_verify_code.go`, реально сгенерированный 5-значный код игнорировался.

## Решение

Переключили верификацию на HTTP-ретранслятор (`me` провайдер) и написали сервис `otp-tg-sender`, который пересылает OTP-коды в Telegram.

## Как работает

```
клиент        → teamgram (bff)  → HTTP GET :8181/code?phone=...&code=...
                                         ↓
                                  otp-tg-sender  → Telegram Bot API  → админ в Telegram
```

1. Пользователь запрашивает вход (`auth.sendCode`) — bff генерирует случайный 5-значный код
2. bff делает HTTP GET на `http://127.0.0.1:8181/code?phone=...&code=...`
3. `otp-tg-sender` получает код и отправляет его сообщением в Telegram указанному пользователю
4. Админ диктует/пересылает код пользователю
5. Пользователь вводит код — bff проверяет через `me` провайдер (сравнение с сохранённым extraData)

## Настройка

### 1. Создать Telegram-бота

Написать [@BotFather](https://t.me/BotFather), команда `/newbot`, получить токен.

Написать боту любое сообщение (чтобы он мог отправлять ответы).

### 2. Прописать токен

В `.env`:
```
TG_BOT_TOKEN=123456:abcdef...
```

### 3. Пересобрать и запустить

```bash
docker compose up -d --build
```

## Изменённые файлы

| Файл | Что сделано |
|------|-------------|
| `pkg/code/send_sms_helper.go` | Добавлен `case "me"` в фабрику `NewVerifyCode` |
| `pkg/code/me/me_verify_code.go` | Починено: `SendSmsVerifyCode` возвращает код, `VerifySmsCode` возвращает `mtproto.ErrPhoneCodeInvalid` |
| `teamgramd/etc/bff.yaml` | `Name: "me"`, `SendCodeUrl: "http://127.0.0.1:8181/code"` |
| `teamgramd/etc2/bff.yaml` | То же (используется в Docker) |
| `app/bff/bff/etc/bff.yaml` | То же |
| `app/service/otp-tg-sender/main.go` | Новый сервис: HTTP-сервер на `:8181`, принимает `?phone=...&code=...`, шлёт в Telegram |
| `build.sh` | Добавлена сборка `otp-tg-sender` |
| `teamgramd/bin/runall-docker.sh` | Добавлен запуск `otp-tg-sender` первым процессом |
| `docker-compose.yaml` | Добавлен `TG_BOT_TOKEN` в `environment` |
| `.env` | Добавлен `TG_BOT_TOKEN` |

## Альтернативные варианты (не реализованы)

- **Реальный SMS-шлюз** — создать новый провайдер в `pkg/code/` (Twilio, Aliyun SMS и т.д.)
- **Админское управление пользователями** — таблица `predefined_users` уже есть в БД, можно разблокировать `auth.toggleBan` или сделать отдельный gRPC/HTTP эндпоинт
- **Быстрый фикс без SMS** — просто заменить `"12345"` на сравнение с реально сгенерированным кодом в `none_verify_code.go`
