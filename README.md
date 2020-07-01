![](./assets/Browzy.png)

![](https://github.com/piaverous/browzy/workflows/goreleaser/badge.svg?branch=master)

[![Open Source Love](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://github.com/ellerbrock/open-source-badges/) [![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php) [![star this repo](http://githubbadges.com/star.svg?user=piaverous&repo=browzy&style=flat)](https://github.com/piaverous/browzy) [![fork this repo](http://githubbadges.com/fork.svg?user=piaverous&repo=browzy&style=flat)](https://github.com/piaverous/browzy/fork) 

## A CLI utility to manage web bookmarks

### Installing

This tool is only available on MacOS for now 👨🏼‍💻

To use it, download and add it to your path like so:

```
curl -o browzy https://github.com/piaverous/browzy/releases/download/v0.1.0/browzy
mv browzy /usr/local/bin
```


### Usage

```
Browzy is a CLI for managing Web bookmarks directly 
in your terminal.

Usage:
  browzy [command]

Available Commands:
  help        Help about any command
  new         Add a new bookmark
  open        Browse your bookmarks and open them in your browser

Flags:
  -f, --file string   path to bookmarks file (default "/Users/piaverous/.browzy")
  -h, --help          help for browzy

Use "browzy [command] --help" for more information about a command.
```