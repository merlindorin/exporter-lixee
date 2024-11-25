# Taskfile Git Pre-Commit Hook

> The provided Taskfile is intended to be a reusable component that houses common features for various repositories.

<!-- TOC -->
* [Taskfile Git Pre-Commit Hook](#taskfile-git-pre-commit-hook)
  * [Summary](#summary)
  * [Prerequisites](#prerequisites)
  * [Configuration](#configuration)
  * [Usage](#usage)
    * [Examples](#examples)
    * [Default Hook Behavior](#default-hook-behavior)
<!-- TOC -->

## Summary

This `Taskfile` provides a set of git tasks. As a first iteration, it only includes a task to set up a Git pre-commit
hook that automatically executes tasks ending with a specified pattern whenever a commit is attempted. This enables
you to enforce code and commit quality standards by running tests, linters, formatters, or any other actions
configured in the rest of your `Taskfile`. This is particularly useful when including multiple `Taskfile` with
pre-commit tasks.

## Prerequisites

* [Task](https://taskfile.dev/): A task runner / simpler Make alternative written in Go.
* [jq](https://stedolan.github.io/jq/): A command-line JSON processor.

Make sure these are installed and configured in your system before proceeding.

## Configuration

The supplied Taskfile utilizes the following variable to define the pattern that triggers tasks on pre-commit:

* `DEFAULT_GIT_PRE_COMMIT_PATTERN`: Regex pattern for the name of tasks to execute as pre-commit hooks (default
  is  `.*\pre-commit$`).

While it is possible to update the configuration, this is not recommended. The default pattern `.*\pre-commit$` is
usually used in included Taskfile.

## Usage

The primary task provided in this Taskfile is the  `install`  task, which installs the Git pre-commit hook.

### Examples

Regarding this `Taskfile.yaml`:

```yaml
version: '3'

includes:
  git: ./tasks.yaml
  
tasks:
  pre-commit:
    cmd:
      - echo "lorem"
```

To install the pre-commit hook using the default pattern you will have to run:

```shell
task git:install
```

After the installation of the hook, any task that match the `PATTERN` will be run before each commit. In this example,
the command `echo "lorem"` will be executed because the default pattern is `.*\pre-commit$`.

### Default Hook Behavior

1. Stashes untracked and unstaged changes.
2. Runs predetermined tasks based on the defined  `PATTERN`.
3. If all tasks pass, adds all changes to the current commit.
4. If any task fails, restores the repository to its original state without unstaged changes and exits with an error
   code.
5. Restores previously stashed changes after the tasks have run.
