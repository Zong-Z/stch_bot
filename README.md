# STCHB - Simple chat-bot (Status - `in development`)

[![CodeFactor](https://www.codefactor.io/repository/github/zong-z/stch_bot/badge)](https://www.codefactor.io/repository/github/zong-z/stch_bot)

**STCHB** - is a simple Telegram chat-bot written in Go (Golang) for a fun, using "[Go Telegram Bot API](https://github.com/go-telegram-bot-api/telegram-bot-api)"

![Gopher image](https://golang.org/doc/gopher/fiveyears.jpg)
*Gopher image by [Renee French][rf], licensed under [Creative Commons 3.0 Attributions license][cc3-by].*

## Downloads

Go to [releases page](https://github.com/Zong-Z/stch_bot/releases) for details.

### Run

1) You need to install [Docker](https://docs.docker.com/get-docker) and [Docker Compose](https://docs.docker.com/compose/install).
2) `cd <folder with the project>/stch_bot`.
3) Setting up configs [`configs.toml`](https://github.com/Zong-Z/stch_bot/blob/master/configs/configs.toml).
4) `docker-compose up -d --build --remove-orphans`.

### Development

#### Prerequisites

- Recommended IDEs
  - [JetBrains GoLand IDE](https://www.jetbrains.com/go) (2020.2.2 and above)
  - [Visual Studio Code](https://code.visualstudio.com) (1.48 and above)
- [Go (Golang)](https://golang.org/dl)

#### Dependencies

- [Golang Telegram Bot API](https://github.com/go-telegram-bot-api/telegram-bot-api)
  - `go get -u github.com/go-telegram-bot-api/telegram-bot-api`
- [Redis client for Golang](https://github.com/go-redis/redis)
  - `go get github.com/go-redis/redis/v8`
- [Go package for UUIDs](https://github.com/google/uuid)
  - `go get github.com/google/uuid`
- [TOML parser and encoder for Go with reflection](https://github.com/BurntSushi/toml)
  - `go get github.com/BurntSushi/toml`

## TODO

-[x] Opportunity to choose a city and age.

-[x] Ability to send photos, videos and documents to chat.

## Contributions

If you have **questions**, **ideas** or you find a **bug**, you can create an [issue](https://github.com/Zong-Z/stch_bot/issues) and it will be reviewed. If you want to contribute to the source code, fork this repository (`master`), realize your ideas and then create a new pull request. **Feel free!**

## License

Developed by **Zong-Z (Nazar)** as open source software under the [MIT License](https://github.com/Zong-Z/stch_bot/blob/master/LICENSE).
