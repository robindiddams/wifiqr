# Wifiqr

a simple tool for sharing wifi

```bash
go install github.com/Robindiddams/wifiqr
# it will ask to unlock your keychain
wifiqr
```

Only supports MacOS, for now.

Because its annoying to type your username and password each time, it caches the qrcode image in `$HOME/.wifiqr`. If you dont want that you can run with `--no-cache`, or if you want to change the dir just set `WIFIQR_DIR` in your env to whereever you want it.