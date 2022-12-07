# Choba is Love, Choba is Life

![Chobanyan-nyan](https://fumetteria.moe/wp-content/uploads/2022/12/choba.gif)

### Choba is GoLang library created as a "joke"

### Contributors :
* Bogdan Budei
* Giuseppe Difilippo
* Simone Coccè

## Features

### 1. ChobaTypes

|   ChobaType   |  Go Type  |
| ------------- | --------- |
| ChobaInt      |   int64   |
| ChobaUInt     |   uint64  |
| ChobaFloat    |   float64 |
| ChobaString   |   string  |
| ChobaRune     |   rune    |
| ChoByte       |   byte    |

### 2. ChobaFunctions

|     Function      |   Parameters  |   Returns    |                        Description                     |               Notes               |
| ----------------- | ------------- | ------------ | ------------------------------------------------------ | --------------------------------- |
| Choba()           | n ChobaInt    |      N/A     | Prints 'Choba' n times                                 | None                              |
| GetChoba()        | N/A           |      N/A     | Prints an ASCII art of The Choba                       | 80 columns needed                 |
| GetChobaColor()   | N/A           |      N/A     | Prints a color ANSI art of The Choba                   | Only works on Linux               |
| TerminalChoba()   | N/A           |      int     | Returns terminal width int value                       | None                              |
| ChobaCheck()      | w int, rw int |      bool    | Returns true if w (width) is >= rw (required width)    | To be used with TerminalChoba()   |


### © ChobaSoft 2022
### Made with ❤️ from UniMi