## Go Forever

### Requirements

This service requires environment variable

1. GOFOREVER_CONFIG = c:/path/to/config.json

```javascript
{
	"Run":{
		"StartCommand":"VBoxManage startvm ubuntu --type HeadLess",
		"EndCommand":"VBoxManage controlvm ubuntu poweroff soft"
	},
	"LogsFile":"./logs"
}

```

2. Please make sure virtualbox is also installed as administrator
3. For windows nssm service is required  for more information please goto (https://nssm.cc/)

### Installation
```
./goforever.exe install
```


