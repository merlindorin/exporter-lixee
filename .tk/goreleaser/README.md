# Gorelease Taskfile

This repository contains a reusable Taskfile that provides common goreleaser features for various Golang projects.

<!-- TOC -->
* [Gorelease Taskfile](#gorelease-taskfile)
  * [Usage](#usage)
<!-- TOC -->

## Usage

To use this `Taskfile`, include it in your project's `Taskfile`. For example:

```yaml
version: '3'

includes:
  goreleaser: https://raw.githubusercontent.com/vanyda-official/taskfile-goreleaser/main/tasks.yaml
```
