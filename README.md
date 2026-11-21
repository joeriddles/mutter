```
$ export MUTTER_NAME=Joe
$ mutter
Joe The Bag of Bones types...
```

Mutter picks names in this order:
1. The `-name` flag.
2. The `MUTTER_NAME` environment variable.
3. The first word of git config `user.name`.
4. The first word of the full name in the [GECOS field](https://en.wikipedia.org/wiki/Gecos_field)
   of the current user. On Windows, Mutter uses the display name.
5. The login name, which `whoami` prints.
6. "Human", if all else fails.
