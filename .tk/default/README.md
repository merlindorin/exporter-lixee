# Taskfile Common Default

> The provided Taskfile is intended to be a reusable component that houses common features for various repositories.

## Summary

This `Taskfile` contains tasks that facilitate running default commands by pattern matching task names.

## Prerequisites

To get started with using this Taskfile, ensure you have the following prerequisites installed:

- [Task](https://taskfile.dev/): A task runner / simpler Make alternative written in Go.
- [jq](https://stedolan.github.io/jq/): A lightweight and flexible command-line JSON processor.

These tools need to be installed and available in your system's PATH to execute the tasks defined in the Taskfile
properly.

## Installation

This `Taskfile` should be included using `default` alias. For example:

```yaml
version: '3'

includes:
  default: ./tasks.yaml

dotenv:
  - .env
  - .env.default
```

In this way, you can run all default task by simply running `task` and not only one task called "default".

## Configuration

The Taskfile uses variables that can be customized for different scenarios. By default, it introduces the following
variable:

- `DEFAULT_DEFAULT_PATTERN`: This pattern defines the regular expression used to match default task names. The default
  value is  `.*:default`.

## Usage

The main feature of this Taskfile is the  `default`  task. This task runs all tasks in your Taskfile that match a
specific pattern.

### Variables

The  `default`  task utilizes the following variables:

- `PATTERN`: Specifies the task name pattern to match against when running the default tasks. The default value
  is  `{{.DEFAULT_DEFAULT_PATTERN}}`, which can be overridden at runtime.

### Summary

- **Description**: Runs a set of default commands based on the provided pattern.
- **Usage**:
  - `task`: Runs all tasks that have names matching the default pattern.

## Examples

Check folder [`./examples`](examples).