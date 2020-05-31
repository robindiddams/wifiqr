# Wifiqr

Creates a qr code for connecting to the network you're connected to.

You just got connected to a new wifi and everyone's asking you for the deets? Try this:

```bash
go install github.com/Robindiddams/wifiqr
# it will ask to unlock your keychain
wifiqr
```
It'll pop up with a qr code that lets people join the network you're on.

Only supports MacOS, for now.

Because its annoying to type your username and password each time, it caches the qrcode image in `$HOME/.wifiqr`. If you dont want that you can run with `--no-cache`, or if you want to change the dir just set `WIFIQR_DIR` in your env to whereever you want it.
