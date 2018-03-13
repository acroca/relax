# Relax

This is a go app that plays ambient sounds (rain, wind, ocean weaves, etc.) at the same time changing their volume over time to random values, resulting in a relaxing and not repetitive background sound.

There are many similar products like this, like [A soft murmur](https://asoftmurmur.com/) and [Noisly](https://www.noisli.com/).

The main reason I built my own (and less fancy) is to control it with the media keys in the keyboard without any external plugin.

# Features

- Generates a single binary with the OGG audio tracks embeded.
- Subscribes to the play/pause media key on OSX

# Building it

`make build` should do the job. It'll end up creating a `relax` binary in the root folder, feel free to move it to wherever you prefer, for example `mv relax /usr/local/bin`.

# Current Tracks

Copied from [https://github.com/Muges/ambientsounds](https://github.com/Muges/ambientsounds).
