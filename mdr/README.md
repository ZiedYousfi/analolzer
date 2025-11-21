# MDR (a Rust crate to parse and give out usable information about .rofl files)

MDR is a Rust crate that allows you to parse .rofl files (League of Legends replay files) and extract useful information from them. It provides functionality to read the binary structure of .rofl files and retrieve metadata such as match details, player information, and game events.

The files you give must be unmodified from the file name to the data it contains. At least that's what MDR expects to work properly.

## Versioning

The versioning of the crate follows the patch versioning scheme of League of Legends.

For example :

- MDR version 25.23.*
- League of Legends version 15.23.*
- ROFL files V15.23.*

- MDR version 25.23.0
- League of Legends version 15.23
- ROFL files V15.23.726.9074

This means that MDR 25.23.0 is compatible with League of Legends 15.23 and ROFL files V15.23.726.9074.

## Philosophy

This is totally made by reverse engineering. I know that there is a lot of this kind of software around but it always feels like they are not up to date and doesn't really give usable information by other software. So I decided to make my own implementation in Rust. I used a lot of different tools to reverse engineer the .rofl files, including a hex editor, a disassembler, and a debugger. I also looked at existing open source projects that deal with .rofl files to get an idea of how they work. But none of them were really up to date/easy to use.

This isn't being built for cheating reasons as I assume that once a match is played, there is nothing you really can do to cheat with it. This is more for data analysis, match history tracking, and other similar use cases.
