# 👩‍💻 WebF-CLI 👨‍💻

Official WebF CLI

## Usage

#### Convert json file in webf file
```bash
./WebF-CLI.exe parseFile ./data.json ./data.webf
```

#### Convert webf file in json file
```bash
./WebF-CLI.exe decodeFile ./data.webf ./data.json
```

#### Convert json data in webf file
```bash
./WebF-CLI.exe parse {"title":"hello world !","meta": ["key":"desc","value"],"content":"..."} ./data.json
```

#### Convert webf data in json file
```bash
./WebF-CLI.exe decode [webf data] ./data.webf
```

#### Convert webf data in json data and get them
```bash
./WebF-CLI.exe decode [webf data]
```

#### Convert json data in webf data and get them
```bash
./WebF-CLI.exe parse {"title":"hello world !","meta": ["key":"desc","value"],"content":"..."}
```

#### Convert webf file in json data and get them
```bash
./WebF-CLI.exe decodeFile ./data.webf
```

#### Convert json file in webf data and get them
```bash
./WebF-CLI.exe parseFile ./data.json
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MPL](https://choosealicense.com/licenses/mpl-2.0/)