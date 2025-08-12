# Catppuccin Color Scheme Generator

This project generates Catppuccin color schemes for various applications based on the templates in the `templates` directory. The colors can be customized by editing the `whisker-color-overrides.json` file.

## Usage

### Prerequisites

*   [whiskers](https://github.com/catppuccin/whiskers)
*   [jq](https://stedolan.github.io/jq/)
*   [runfile](https://github.com/TekWizely/runfile)

### Generating the Color Schemes

To generate the color schemes, run the following commands:

*   `run light-theme`: Generates the `latte` variant of the color scheme.
*   `run dark-theme`: Generates the `mocha` variant of the color scheme.

The generated files will be placed in the `generated` directory.

### Installing the Color Schemes

To install the generated color schemes, run the following command:

```bash
run install
```

This will create a symbolic link from the `generated` directory to `~/.colorscheme.d`.

### Customizing the Colors

The colors can be customized by editing the `whisker-color-overrides.json` file. This file contains the color overrides for the `latte` and `mocha` variants of the color scheme. After editing this file, you will need to regenerate the color schemes for the changes to take effect.

### Color Palette

| Color Name | Latte | Mocha |
|:---|:---:|:---:|
| rosewater | <p align="center">![#dc8a78](https://place-hold.it/100x40/dc8a78/000000&text=)<br>#dc8a78</p> | <p align="center">![#f0d0c8](https://place-hold.it/100x40/f0d0c8/000000&text=)<br>#f0d0c8</p> |
| flamingo | <p align="center">![#dd7878](https://place-hold.it/100x40/dd7878/000000&text=)<br>#dd7878</p> | <p align="center">![#e8b8b8](https://place-hold.it/100x40/e8b8b8/000000&text=)<br>#e8b8b8</p> |
| pink | <p align="center">![#ea76cb](https://place-hold.it/100x40/ea76cb/000000&text=)<br>#ea76cb</p> | <p align="center">![#e8a8d0](https://place-hold.it/100x40/e8a8d0/000000&text=)<br>#e8a8d0</p> |
| mauve | <p align="center">![#8839ef](https://place-hold.it/100x40/8839ef/ffffff&text=)<br>#8839ef</p> | <p align="center">![#c098e0](https://place-hold.it/100x40/c098e0/000000&text=)<br>#c098e0</p> |
| red | <p align="center">![#d20f39](https://place-hold.it/100x40/d20f39/ffffff&text=)<br>#d20f39</p> | <p align="center">![#d87a8a](https://place-hold.it/100x40/d87a8a/000000&text=)<br>#d87a8a</p> |
| maroon | <p align="center">![#e64553](https://place-hold.it/100x40/e64553/ffffff&text=)<br>#e64553</p> | <p align="center">![#e098a8](https://place-hold.it/100x40/e098a8/000000&text=)<br>#e098a8</p> |
| peach | <p align="center">![#fe640b](https://place-hold.it/100x40/fe640b/ffffff&text=)<br>#fe640b</p> | <p align="center">![#e8b088](https://place-hold.it/100x40/e8b088/000000&text=)<br>#e8b088</p> |
| yellow | <p align="center">![#df8e1d](https://place-hold.it/100x40/df8e1d/ffffff&text=)<br>#df8e1d</p> | <p align="center">![#e8d098](https://place-hold.it/100x40/e8d098/000000&text=)<br>#e8d098</p> |
| green | <p align="center">![#4f6642](https://place-hold.it/100x40/4f6642/ffffff&text=)<br>#4f6642</p> | <p align="center">![#98d090](https://place-hold.it/100x40/98d090/000000&text=)<br>#98d090</p> |
| teal | <p align="center">![#426060](https://place-hold.it/100x40/426060/ffffff&text=)<br>#426060</p> | <p align="center">![#88d0c0](https://place-hold.it/100x40/88d0c0/000000&text=)<br>#88d0c0</p> |
| sky | <p align="center">![#3d6b85](https://place-hold.it/100x40/3d6b85/ffffff&text=)<br>#3d6b85</p> | <p align="center">![#80c8d8](https://place-hold.it/100x40/80c8d8/000000&text=)<br>#80c8d8</p> |
| sapphire | <p align="center">![#3d6275](https://place-hold.it/100x40/3d6275/ffffff&text=)<br>#3d6275</p> | <p align="center">![#70b8d0](https://place-hold.it/100x40/70b8d0/000000&text=)<br>#70b8d0</p> |
| blue | <p align="center">![#3d5085](https://place-hold.it/100x40/3d5085/ffffff&text=)<br>#3d5085</p> | <p align="center">![#80a8e8](https://place-hold.it/100x40/80a8e8/000000&text=)<br>#80a8e8</p> |
| lavender | <p align="center">![#7287fd](https://place-hold.it/100x40/7287fd/ffffff&text=)<br>#7287fd</p> | <p align="center">![#a8b0e8](https://place-hold.it/100x40/a8b0e8/000000&text=)<br>#a8b0e8</p> |
| text | <p align="center">![#4c4f69](https://place-hold.it/100x40/4c4f69/ffffff&text=)<br>#4c4f69</p> | <p align="center">![#d8e0f0](https://place-hold.it/100x40/d8e0f0/000000&text=)<br>#d8e0f0</p> |
| subtext1 | <p align="center">![#5c5f77](https://place-hold.it/100x40/5c5f77/ffffff&text=)<br>#5c5f77</p> | <p align="center">![#c0c8e0](https://place-hold.it/100x40/c0c8e0/000000&text=)<br>#c0c8e0</p> |
| subtext0 | <p align="center">![#6c6f85](https://place-hold.it/100x40/6c6f85/ffffff&text=)<br>#6c6f85</p> | <p align="center">![#a8b0d0](https://place-hold.it/100x40/a8b0d0/000000&text=)<br>#a8b0d0</p> |
| overlay2 | <p align="center">![#7c7f93](https://place-hold.it/100x40/7c7f93/ffffff&text=)<br>#7c7f93</p> | <p align="center">![#909ac0](https://place-hold.it/100x40/909ac0/000000&text=)<br>#909ac0</p> |
| overlay1 | <p align="center">![#8c8fa1](https://place-hold.it/100x40/8c8fa1/ffffff&text=)<br>#8c8fa1</p> | <p align="center">![#7880a0](https://place-hold.it/100x40/7880a0/ffffff&text=)<br>#7880a0</p> |
| overlay0 | <p align="center">![#9ca0b0](https://place-hold.it/100x40/9ca0b0/000000&text=)<br>#9ca0b0</p> | <p align="center">![#606880](https://place-hold.it/100x40/606880/ffffff&text=)<br>#606880</p> |
| surface2 | <p align="center">![#acb0be](https://place-hold.it/100x40/acb0be/000000&text=)<br>#acb0be</p> | <p align="center">![#505870](https://place-hold.it/100x40/505870/ffffff&text=)<br>#505870</p> |
| surface1 | <p align="center">![#bcc0cc](https://place-hold.it/100x40/bcc0cc/000000&text=)<br>#bcc0cc</p> | <p align="center">![#404860](https://place-hold.it/100x40/404860/ffffff&text=)<br>#404860</p> |
| surface0 | <p align="center">![#ccd0da](https://place-hold.it/100x40/ccd0da/000000&text=)<br>#ccd0da</p> | <p align="center">![#303850](https://place-hold.it/100x40/303850/ffffff&text=)<br>#303850</p> |
| base | <p align="center">![#eff1f5](https://place-hold.it/100x40/eff1f5/000000&text=)<br>#eff1f5</p> | <p align="center">![#1a1a24](https://place-hold.it/100x40/1a1a24/ffffff&text=)<br>#1a1a24</p> |
| mantle | <p align="center">![#e6e9ef](https://place-hold.it/100x40/e6e9ef/000000&text=)<br>#e6e9ef</p> | <p align="center">![#181825](https://place-hold.it/100x40/181825/ffffff&text=)<br>#181825</p> |
| crust | <p align="center">![#dce0e8](https://place-hold.it/100x40/dce0e8/000000&text=)<br>#dce0e8</p> | <p align="center">![#11111b](https://place-hold.it/100x40/11111b/ffffff&text=)<br>#11111b</p> |

