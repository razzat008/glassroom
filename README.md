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
- groupchat id
- api key for your telegram bot  
and  
- credentials.json in root of this project containing oauth credentials
# Things to implement
- Sending updates to telegram automatically
- queuing updates or smth idk 

# References
- Official [quickstart](https://developers.google.com/workspace/classroom/quickstart/go) guide from [developers.google.com](https://developers.google.com/)
