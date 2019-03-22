# Go Search Pwned

This utility uses [haveibeenpwned's](https://haveibeenpwned.com) API to search their database of leaked passwords, to see if your password has been included in any of their known leaks

There's an awesome video about how it works and why it's secure [Here](https://www.youtube.com/watch?v=hhUb5iknVJs)

## TL;DR Of Why It's Secure

[haveibeenpwned's](https://haveibeenpwned.com) API only uses the first 5 characters of the **SHA1 hash** of your password and returns all known hashes the share the first 5 characters

The only part of this process which has access to your whole password is the Go script I've written

Why should you trust me? **Don't**, in fact, **PLEASE DON'T**

I encourage you to review and compile the source code yourself, or even build your own implementation of this tool

However, after reviewing the source code, you'll see that the script is just hashing the password for you, chopping of the first 5 characters, `POST`ing them to [haveibeenpwned's](https://haveibeenpwned.com) API, searching the response for the full hash of your password, and telling whether or not the full hash of your password was found in their leaked password database

No where am I sending off or storing your plain text password

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
