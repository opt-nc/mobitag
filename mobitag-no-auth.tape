# Where should we write the GIF?
Output media/no-auth/mobitag.webm

Set Margin 60
Set MarginFill "#6B50FF"
Set BorderRadius 10

# Set up a 1200x600 terminal with 46px font.
Set FontSize 14
Set Width 1200
Set Height 600

# Type a command in the terminal.
Type "# 🙌 Welcome to mobitag cli demo"

Sleep 500ms
# Run the command by pressing enter.
Enter
Enter
Sleep 1s

# Type a command in the terminal
# Type "# 📜 Get source code"
# Enter
# Type "gh repo clone https://github.com/opt-nc/mobitag-cli.git"
# Enter
# Sleep 4000ms
# Type "cd mobitag-cli"
# Enter
# Sleep 2000ms
Type "ls"
Sleep 200ms
Enter
Sleep 1s
Sleep 200ms
# Run the command by pressing enter.

Enter
Type "# ⚙️ Build the cli"
Enter
Type "task build"
Enter
Sleep 3000ms

Type "# Add the cli to your path"
Enter
Type "export PATH=$PATH:$(pwd)/bin"
Enter
Sleep 2000ms

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
Type `mobitag send --to ${MOBILIS_NUMBER} --message "Hello World : a mobitag from Go(lang) XD"`
Enter
Sleep 5s

Type "# 📲 That's all folks!"
Enter
Sleep 1s

Type "# 🔖 Stay connected"
Enter
Type "# github.com/opt-nc/mobitag"
Sleep 3s
