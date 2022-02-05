# WTColorSchemesSort
Sorts the color schemes jsonfield-entries in your Windows Terminal settings.json.

## Requirements
Go 1.17+ (I have used this, but actual required version to run could be lower)
settings.json from Windows Terminal

## Usage
Pass everything in bracket after "schemes" key including brackets as string.
End with tilde ~

#### Example String
```json
[
  {
    "name": "Aurelia",
    "foreground": "#EA549F",
    "background": "#1A1A1A",
    "cursorColor": "#FFFFFF",
    "selectionBackground": "#FFFFFF",
    "black": "#000000",
    "red": "#E92888",
    "green": "#4EC9B0",
    "yellow": "#CE9178",
    "blue": "#579BD5",
    "purple": "#714896",
    "cyan": "#00B6D6",
    "white": "#EAEAEA",
    "brightBlack": "#797979",
    "brightRed": "#EB2A88",
    "brightGreen": "#1AD69C",
    "brightYellow": "#E9AD95",
    "brightBlue": "#9CDCFE",
    "brightPurple": "#975EAB",
    "brightCyan": "#2BC4E2",
    "brightWhite": "#EAEAEA"
  },
  {
    "name": "UbuntuLegit",
    "foreground": "#EEEEEE",
    "background": "#2C001E",
    "cursorColor": "#FFFFFF",
    "selectionBackground": "#FFFFFF",
    "black": "#4E9A06",
    "red": "#CC0000",
    "green": "#300A24",
    "yellow": "#C4A000",
    "blue": "#3465A4",
    "purple": "#75507B",
    "cyan": "#06989A",
    "white": "#D3D7CF",
    "brightBlack": "#555753",
    "brightRed": "#EF2929",
    "brightGreen": "#8AE234",
    "brightYellow": "#FCE94F",
    "brightBlue": "#729FCF",
    "brightPurple": "#AD7FA8",
    "brightCyan": "#34E2E2",
    "brightWhite": "#EEEEEE"
  }
]~ //Notice this tilde ~, this is important
```

You will get a sorted array of schemes

## Code implementation
```golang
// unsortedJSON is the schemes array including brackets as string
sortedJSON := WTColorSchemesSort.Sort(unsortedJSON)
```