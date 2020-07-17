# Message Committer

[![Build Status](https://dev.azure.com/wandersonolivs/go-msgcommitter/_apis/build/status/obiwandsilva.go-msgcommitter?branchName=master)](https://dev.azure.com/wandersonolivs/go-msgcommitter/_build/latest?definitionId=1&branchName=master)

Message Committer is a git hook that helps you create meaningful and organized commit messages. The message structure is based on the [Angular Commit Messages Convention](https://docs.google.com/document/d/1QrDFcIiPjSLDn3EL15IJygNPiHORgU1_OOAqWjiDU5Y).

## Usage

You can use on both Linux or Windows.

The easiest way to use it is moving the `prepare-commit-msg` (on Windows use `prepare-commit-msg.exe`) file located in the root of this project and move it to the `.git/hooks` directory of your desired project.

But you can also build it manually if you have `go` in your environment with `go build -o prepare-commit-msg cmd/msgcommitter/msgcommitter.go`. Then put the generated file in the `.git/hooks` directory of your desired project.

After moving the file, make your code changes, add the changes to stage and then call `git commit`. The terminal will ask for information about your changes, follow the descriptions. At the end, a file containing the generated and formatted message will be shown by git in a text editor. Examples:

```
doc(README): add instructions

Add more instructions about how to use msgcommitter.

Closes #001
```
Make your corrections to the message or quit the text editor.

To confirm the message was added to your commit history, your can use `git log`.