```
$ export MUTTER_NAME=Joe
$ mutter
Joe The Bag of Bones types...
```

Mutter picks names in this order:
1. The `-name` flag.
2. The `MUTTER_NAME` environment variable.
3. The first word of git config `user.name`.
4. "Human", if all else fails.
