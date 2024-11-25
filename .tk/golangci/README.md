# Golang CI Taskfile

This repository contains a reusable Taskfile that provides common linting features for various Golang projects. The
tasks included in this `Taskfile` use the [golangci-lint](https://github.com/golangci/golangci-lint) tool.

<!-- TOC -->
* [Golang CI Taskfile](#golang-ci-taskfile)
  * [Summary](#summary)
  * [Prerequisites](#prerequisites)
  * [Configuration](#configuration)
  * [Available Tasks](#available-tasks)
  * [Usage](#usage)
<!-- TOC -->

## Summary

This `Taskfile` provides a set tasks that help you lint your Golang project. These tasks are configured to use widely
accepted best practices and tools in the Go community, with golangci-lint at the center.

## Prerequisites

* [Task](https://taskfile.dev/): A task runner / simpler Make alternative written in Go.
* [golangci-lint](https://github.com/golangci/golangci-lint): A fast Go linters runner.

Install these tools before proceeding.

## Configuration

The `Taskfile` comes with these environmental parameters:

* `DEFAULT_GOLANGCI_LINT_FLAGS`: Default flags used by golangci-lint.
* `DEFAULT_GOLANGCI_LINT_TASKS`: Default tasks triggered by golangci-lint.

You can override these parameters by redefining them in your environment or in your `Taskfile`.

## Available Tasks

Here are the tasks that the `Taskfile` provides:

* `lint`: This task runs golangci-lint on your Go code with the flags and tasks defined in `DEFAULT_GOLANGCI_LINT_FLAGS`
  and `DEFAULT_GOLANGCI_LINT_TASKS`.
* `fix`: This task will try to automatically fix any issues found by the `lint` task.
* `boilerplate`: This task helps setting up the structure of a new Go project.

## Usage

To use this `Taskfile`, include it in your project's `Taskfile`. For example:q:

```yaml
version: '3'

includes:
  golangci: https://raw.githubusercontent.com/vanyda-official/taskfile-golangci/main/tasks.yaml
```
