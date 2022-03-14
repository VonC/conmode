# conmode
Display Windows console modes for stdin, stdout, stderr

    go install `github.com/VonC/conmode@latest`

## Description

Uses:

- `windows.GetConsoleMode` with:
  - `windows.STD_INPUT_HANDLE`
  - `windows.STD_OUTPUT_HANDLE`
  - `windows.STD_ERROR_HANDLE`
- Input/output modes listed from https://docs.microsoft.com/en-us/windows/console/setconsolemode

## License: MIT

[LICENSE](LICENSE)
