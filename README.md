# conmode
Display Windows console modes for stdin, stdout, stderr.  
Or decrypt a console mode integer.

    go install `github.com/VonC/conmode@latest`

## Description

### Display console mode

`conmode.exe` or `conmode.exe display` (the default command), will print:

- `windows.GetConsoleMode` with:
  - `windows.STD_INPUT_HANDLE`
  - `windows.STD_OUTPUT_HANDLE`
  - `windows.STD_ERROR_HANDLE`
- Input/output modes listed from https://docs.microsoft.com/en-us/windows/console/setconsolemode

### Decrypt console mode integer

```bash
console.exe 503
conmode 503: 'ENABLE_ECHO_INPUT|ENABLE_INSERT_MODE|ENABLE_LINE_INPUT|ENABLE_MOUSE_INPUT|ENABLE_PROCESSED_INPUT|ENABLE_QUICK_EDIT_MODE'
```

## License: MIT

[LICENSE](LICENSE)
