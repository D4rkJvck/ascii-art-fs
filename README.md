<h1 align=center>ASCII-Art</h1>

This program is designed to produce a graphical representation of a given [STRING] as an argument using ASCII characters. It can use the "flag" command line as the [OPTION] given as the first argument (excluding the file name as an argument).

<h2 align=center>FS (File System)</h2>

The main feature of this program is the FS (File System) feature, which will have the following format:

```console
go run . [STRING] [BANNER]
```

where [BANNER], given as the second argument, represents the filename to be used as a template for the ASCII-Art. The program will parse this file and then return a table of string lines. Additionally, this program's functionality might be useful for its other features.

<h2 align=center>OUTPUT</h2>

This functionality allows the program to use the File System's functionality as suggested by the following format:

```console
go run . [OPTION] [STRING] [BANNER]
```

However, the result won't be printed to the terminal. Instead, the program will write it to a new file (which it will create) or overwrite it in an existing one. The filename is given as the second argument after the flag, which must follow the pattern "--" followed by the feature's name and "=" followed by the filename.

<h2 align=center>COLOR</h2>

For the COLOR feature, there are two possibilities: ANSI and RGB.

- For ANSI color, use the option `--color=color` with the color specified in lowercase English. For example:

```console
go run . --color=red hello
```

- For RGB color, use the option `--color=rgb(r;g;b)` where r, g, and b are integers between 0 and 255. The option must be enclosed in double or single quotes. For example:

```console
go run . '--color=rgb(255;0;0)' hello
```

<h2 align=center>JUSTIFY</h2>

This functionality will use the File System's functionality with the following format:

```console
go run . [OPTION] [STRING] [BANNER]
```

This functionality is supposed to determine where to print the result in the terminal, which implies finding out the terminal size first. So, the pattern of the second argument will be the same as the previous ones, except for the feature's name "align," then the place to align it after the "=" such as "right," "left," "center," etc.

## Dependencies

- Go (1.15+ recommended)
- Standard Go packages
- `github.com/stretchr/testify` (for testing)

```
```
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
