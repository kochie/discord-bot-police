## DiscordGo Police Bot

A discord bot to do random things on the server.

- Detects anime and gets angry
- Detects furries and bans them.
- Plays AOE2 sounds

**Join [Discord Gophers](https://discord.gg/0f1SbxBZjYoCtNPP)
Discord chat channel for support.**

### Build

This assumes you already have a working Go environment setup and that
DiscordGo is correctly installed on your system.


From within the pingpong example folder, run the below command to compile the
example.

```sh
go build
```

### Usage

This example uses bot tokens for authentication only. While user/password is 
supported by DiscordGo, it is not recommended for bots.

```
./pingpong --help
Usage of ./pingpong:
  -t string
        Bot Token
```

The below example shows how to start the bot

```sh
./pingpong -t YOUR_BOT_TOKEN
Bot is now running.  Press CTRL-C to exit.
```
