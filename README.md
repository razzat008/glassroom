        a go program that fetches new updates/announcements in the 
        courses I'm enrolled in and sends it elsewhere

# Prerequisites
- install go
```bash
sudo pacman -S go
```
- get oauth key for yourself
 
- prepare yourself for dependencies
```bash
go mod tidy
```

# Usage
> [!WARNING]
> create a src/token.json to store:
> groupchat id
> api key for your telegram bot  
> and  
> credentials.json in root of this project containing oauth credentials

```bash
go run src/tel_bot.go # authenticate from the displayed url
```
After signing in, you'll get redirected to localhost(probably)
copy the token from the url and paste it in the terminal

# Contributing
All contributors are welcome.
- [Fork](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/working-with-forks/fork-a-repo) the repo
- Make changes  
- Open a PR ; specify the changes you've made

# Things to implement
- Sending updates to telegram automatically
- queuing updates or smth idk 

# References
- Official [quickstart](https://developers.google.com/workspace/classroom/quickstart/go) guide from [developers.google.com](https://developers.google.com/)
