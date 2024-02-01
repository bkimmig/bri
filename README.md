# bri

A cli for my personal tasks

# Install

Clone the repo and run `go install .` from the root.

# Commands

## wordle

0 indicates the letter is a miss
1 indicates the letter is in the incorrect place
2 indicates the letter is in it's exact position

usage

```bash
bri wordle
```

which will open a prompt for you start entering what you know.

Say the word is "bulky" for the day you'd enter your first guess as

```bash
>>> stray 00002
```

It will then spit out words ordered by their entropy.

Put in your next guess

```bash
>>> bulgy 22202
```

And you'll have a small list to pick from.
