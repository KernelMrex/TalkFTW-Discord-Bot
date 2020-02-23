# Talk FTW announcer discord bot

To add bot on your channel [click here](https://discordapp.com/api/oauth2/authorize?client_id=680475756467453967&permissions=3213312&scope=bot)

## Helpful

To create audio

```bash
git clone https://github.com/bwmarrin/dca
ffmpeg -i welcome.mp3 -f s16le -ar 48000 -ac 2 pipe:1 | ./dca > audio.dca 
```