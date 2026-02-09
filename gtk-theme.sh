#! /usr/bin/env bash

gtk_theme="Orchis-Purple-Light"
color_scheme="prefer-light"
[ "$MODE" = "dark" ] && gtk_theme="Orchis-Purple-Dark" && color_scheme="prefer-dark"

# Set system color scheme (Firefox and other apps follow this for live reload)
gsettings set org.gnome.desktop.interface color-scheme "$color_scheme" 2>/dev/null || true

icon_theme="Papirus-Light"
[ "$MODE" = "dark" ] && icon_theme="Papirus-Dark"

cursor_theme="Bibata-Modern-Ice"

cat > ~/.config/gtk-3.0/settings.ini <<EOF
[Settings]
gtk-theme-name=$gtk_theme
gtk-icon-theme-name=$icon_theme
gtk-font-name=Sans 12
gtk-cursor-theme-name=$cursor_theme
EOF

cat > ~/.config/gtk-4.0/settings.ini <<EOF
[Settings]
gtk-theme-name=$gtk_theme
gtk-icon-theme-name=$icon_theme
gtk-font-name=Sans 12
gtk-cursor-theme-name=$cursor_theme
EOF
