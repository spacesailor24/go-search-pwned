# Go Search Pwned

This utility uses [haveibeenpwned's](https://haveibeenpwned.com) API to search their database of leaked passwords, to see if your password has been included in any of their known leaks

There's an awesome video about how it works and why it's secure [Here](https://www.youtube.com/watch?v=hhUb5iknVJs)

## How to Use

I've compiled three Go binaries for OSX, Windows, and Linux all for the `amd64` architecture

If you need a binary for something other than `amd64`, please let me know by creating an issue, or compiling the source code yourself

Once you have the Go binary on your system, execute it with:

On Linux:

```bash
./go-search-pwned-v1.0 your-password
```

On OSX:

```bash
./go-search-pwned-v1.0-osx your-password
```

On Windows:

I'm not actually sure how this works, because I don't have a Windows machine to test it on, sorry :)

Please create an issue or a PR if you know what the proper command is, thank you!

```bash
./go-search-pwned-v1.0-win your-password
```
