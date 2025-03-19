# Where should we write the GIF?
Output media/auth/mobitag.webm
Output media/auth/mobitag.gif

Set Margin 60
Set MarginFill "#6B50FF"
Set BorderRadius 10

# Set up a 1200x600 terminal with 46px font.
Set FontSize 14
Set Width 1200
Set Height 600

Enter
Type "# ⚙️ Build from source"
Enter
Type "go build -o bin/mobitag main.go"
Enter
Sleep 7000ms

Type "# Add the cli to your path"
Enter
Type "export PATH=$PATH:$(pwd)/bin"
Enter

Type "clear"
Enter
Type "# 🚀 Enjoy the cli"
Enter

Type "# 📝 Get help"
Enter
Type "mobitag --help"
Enter
Sleep 2000ms


Type "# 🧪 Test environment"
Enter
Type "mobitag dryRun"
Enter
Sleep 2000ms

Type "# 📲 Send sms"
Enter
Type `mobitag send --to ${MOBILIS_NUMBER} --from ${MOBILIS_NUMBER} --message "Hello World : a mobitag from Go(lang) XD"`
Enter
Sleep 5s

Type "# 📲 Send sms with a pipe"
Enter
Type `echo "Hello World : a mobitag from Go(lang) XD" | mobitag sendPipe --to ${MOBILIS_NUMBER} --from ${MOBILIS_NUMBER}`
Enter
Sleep 5s

Type "# 📲 That's all folks!"
Enter
Sleep 1s

Type "# 🔖 Stay connected"
Enter
Type "# github.com/opt-nc/mobitag-cli"
Sleep 3s
