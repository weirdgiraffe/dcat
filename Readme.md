# dcat - delayed cat

Simple program to output text with some delay before printing characters.

![](https://github.com/weirdgiraffe/dcat/blob/main/preview.gif)


## Install

```
go install github.com/weirdgiraffe/dcat
```

## Usage

```
Usage: dcat [options] [text]

Print provided [text] with random delays between characters. If [text] is not provided, then prints stdin.

options:
  -max-delay duration
        maximum delay between printing a character (default 200ms)
  -min-delay duration
        minimum delay between printing a character (default 5ms)
```
