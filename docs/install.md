# Installation

!!! note
    Currently supported operating systems and chip architectures:

    1. Darwin/arm64
    2. Darwin/amd64
    3. Linux/amd64

## 1 Install dtm binary with script

In your working directory, run:

```shell
sh -c "$(curl -fsSL https://download.devstream.io/download.sh)"
```

This will download the corresponding `dtm` binary to your working directory according to your OS and chip architecture, and grant the binary execution permission.

!!! note "Optional"

    Move `dtm` to a place which is in your PATH. For example: `mv dtm /usr/local/bin/`.

## 2 Download manually from the GitHub Release page

You could find the latest version of `dtm` on the [Release](https://github.com/devstream-io/devstream/releases/) page and click Download.

If your browser isn't working properly, "curl" is also a good option:

```shell
# Version v0.10.3, OS type linux and CPU arch amd64 need to be modified as needed
$ curl -o dtm https://download.devstream.io/v0.10.3/dtm-linux-amd64
```

Note that there are multiple versions of `dtm` available, so you will need to choose the correct version for your operating system and chip architecture. Once downloaded locally, you can choose to rename it, move it to the directory containing `$PATH` and give it executable permissions, for example, on Linux you can do this by running the following command.

```shell
mv dtm-linux-amd64 /usr/local/bin/dtm
chmod +x dtm
```

Then you can verify that the permissions and version of dtm are correct with the following command.

```shell
$ dtm version
0.10.3
```

## 3 Install dtm with [asdf](https://asdf-vm.com/)

```shell
# Plugin
asdf plugin add dtm
# Show all installable versions
asdf list-all dtm
# Install specific version
asdf install dtm latest
# Set a version globally (on your ~/.tool-versions file)
asdf global dtm latest
# Now dtm commands are available
dtm help
```
