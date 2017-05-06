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

### Installation
```
./goforever.exe install
```


