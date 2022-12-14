# golang text to image discord bot with stable diffusion

[![GitHub version](https://img.shields.io/github/release/jcksnvllxr80/go-txt2img-discord-bot.svg)](lib-release)
[![GitHub download](https://img.shields.io/github/downloads/jcksnvllxr80/go-txt2img-discord-bot/total.svg)](lib-release)
[![GitHub stars](https://img.shields.io/github/stars/jcksnvllxr80/go-txt2img-discord-bot.svg)](lib-stars)
[![GitHub issues](https://img.shields.io/github/issues/jcksnvllxr80/go-txt2img-discord-bot.svg)](lib-issues)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](lib-licence)

## Image of project

![ai_image discord bot](./ai_image.png "Image of project")

## high-level

1. send request to listening bot in discord with description of image
2. golang bot kicks off the creation based on string given in discord
3. image is created using stable diffusion
4. bot returns image to discord to complete the process

## install stable diffusion on a windows system with nvidia graphics card and WSL

[tutorial](https://www.assemblyai.com/blog/how-to-run-stable-diffusion-locally-to-generate-images/) to learn how to install on your windows host

## install go on WSL2

[tutorial](https://dev.to/deadwin19/how-to-install-golang-on-wslwsl2-2880) to learn to install go on WSL2

## configure a discord bot

[tutorial](https://dev.to/aurelievache/learning-go-by-examples-part-4-create-a-bot-for-discord-in-go-43cf) to learn how to create a golang discord bot

## run the bot 

change dir to where `main.go` was compiled

```bash
./main
```

## call the bot into action from discord

```text
@ai_image rick james with a technicolor dreamcoat in space
```
