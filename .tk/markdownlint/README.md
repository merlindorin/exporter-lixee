# Taskfile Markdown Linting Hook

> This repository includes a Taskfile dedicated to Markdown linting, designed to be included and reused across different
> Taskfile configurations.

<!-- TOC -->
* [Taskfile Markdown Linting Hook](#taskfile-markdown-linting-hook)
  * [Summary](#summary)
  * [Prerequisites](#prerequisites)
  * [Configuration](#configuration)
  * [Usage](#usage)
    * [Example Integration](#example-integration)
    * [Linting Markdown Files](#linting-markdown-files)
<!-- TOC -->

## Summary

The Taskfile included in this repository provides tasks for linting Markdown files, promoting consistent and
standardized documentation format across projects. It is meant to be easily integrated with any Taskfile setup, so you
can enforce style rules on your Markdown files using  [markdownlint](https://github.com/markdownlint/markdownlint), a
tool that checks for and corrects style issues in Markdown documents.

## Prerequisites

This Taskfile assumes that you have the following prerequisites installed:

* [Task](https://taskfile.dev/): A task runner that allows you to execute commands and scripts defined in
  a  `Taskfile.yaml`.
* [markdownlint-cli](https://github.com/igorshubovych/markdownlint-cli): A command-line interface for markdownlint used
  for linting Markdown files.

Make sure these prerequisites are met before using the provided Taskfile.

## Configuration

The Taskfile automatically includes the following markdownlint configuration provided by an external Taskfile from the
official source.

The external Taskfile can be imported directly as shown in the Taskfile snippet below:

## Usage

To utilize the Markdown linting tasks, include the external Taskfile in your project's  `Taskfile.yaml`. This enables
you to run linting commands as predefined tasks within your development workflow.

### Example Integration

Assuming you have the following  `Taskfile.yaml`  in your project:

```yaml
version: '3'

includes:
  markdownlint: https://raw.githubusercontent.com/vanyda-official/taskfile-markdownlint/main/tasks.yaml
```

Initially integrate the Taskfile with the above configuration.

### Linting Markdown Files

Once the Taskfile is included, you can lint your Markdown files by invoking a markdownlint task as follows:

```shell
task markdownlint:lint
```

By running this command, the markdownlint tool will be executed to analyze your Markdown files, and any style issues
will be reported accordingly. This ensures that all your Markdown documentation adheres to the predefined style guide.
