<h1 align=center>ASCII-Art</h1>

This program is supposed to output a received [STRING] as argument in a graphic representation using ASCII.
It might use "flag" command line as [OPTION] given as first argument (excluding the file name as argument).

<h2 align=center>FS</h2>
The main feature of this program is the FS (File System) feature, which will have the following format :

```console
go run . [STRING] [BANNER]
```
where [BANNER], given as second argument, stands for the file name to be used as template file for the Print'Art.
The program will scan this file then return a string line table. 
Additionnally, this program's functionality might be useful for its other features.

<h2 align=center>OUTPUT</h2>
This functionality will allow the program to use the File System's functionality as suggested by the following format :

```console
go run . [OPTION] [STRING] [BANNER]
```
It won't be printing the result on the terminal though. Instead, the program will write it into a new file (that it will create), or overwrite it into an existing one. The file name is given as second argument after the flag that has to follow the pattern "--" followed by the feature's name then "=" followed by the file name.

<h2 align=center>COLOR</h2>
This functionality won't use the File System's functionality. It will use the standard template with the following format :

```console
go run . [OPTION] [STRING]
```
For this functionality, the [OPTION] will be the color setting, with the pattern "--" followed by the feature's name (color) then "=" follow by a color name that will determine the color of the final output. An optional setting given as third argument might be a string of letters for these letters to be be colored in the next argument (the string to print).

<h2 align=center>JUSTIFY</h2>
This functionality will use the File System's functionality with the following format :

```console
go run . [OPTION] [STRING] [BANNER]
```
This functionality is supposed to determine where to print the result in the terminal, which implies to find out the terminal size first. So the second argument's pattern will the same as the previous ones except for the feature's name "align" then the place to align it after the "=" such as "right", "left", "center"...