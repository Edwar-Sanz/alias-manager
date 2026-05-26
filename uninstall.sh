#!/usr/bin/env bash

set -e

INSTALL_DIR="$HOME/.local/bin"
BINARY_PATH="$INSTALL_DIR/am"
CONFIG_DIR="$HOME/.config/am"
AMF_SOURCE_LINE='[ -f ~/.config/am/.amf ] && source ~/.config/am/.amf'
WRAPPER_BEGIN='# am — alias manager shell integration'
WRAPPER_END='# end am — alias manager shell integration'

GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[0;33m'
GRAY='\033[0;90m'
RESET='\033[0m'

ok()   { echo -e "${GREEN}✓${RESET} $1"; }
warn() { echo -e "${YELLOW}!${RESET} $1"; }
fail() { echo -e "${RED}✗${RESET} $1"; exit 1; }
info() { echo -e "${GRAY}→${RESET} $1"; }

detect_rc() {
    case "$(basename "$SHELL")" in
        zsh)  echo "$HOME/.zshrc" ;;
        bash) echo "$HOME/.bashrc" ;;
        *)    echo "" ;;
    esac
}

confirm() {
    local prompt="$1"
    local answer
    read -r -p "$prompt [y/N] " answer
    [[ "$answer" =~ ^[Yy]$ ]]
}

echo ""
echo "  Uninstalling am — Alias Manager"
echo ""

# remove binary
if [ -f "$BINARY_PATH" ]; then
    rm -f "$BINARY_PATH" || fail "Could not remove $BINARY_PATH"
    ok "Binary removed from $BINARY_PATH"
else
    info "No binary found at $BINARY_PATH"
fi

# detect rc file
RC_FILE=$(detect_rc)
if [ -z "$RC_FILE" ] || [ ! -f "$RC_FILE" ]; then
    warn "Shell rc file not found. Remove these manually if present:"
    echo ""
    echo "    $AMF_SOURCE_LINE"
    echo ""
    echo "    (and the wrapper block between '$WRAPPER_BEGIN'"
    echo "     and '$WRAPPER_END')"
    echo ""
else
    # backup before modifying
    BACKUP="${RC_FILE}.am-uninstall.bak"
    cp "$RC_FILE" "$BACKUP" || fail "Could not back up $RC_FILE"
    info "Backed up $RC_FILE → $BACKUP"

    CHANGED=0

    # remove wrapper block (begin marker through end marker, inclusive)
    if grep -qF "$WRAPPER_BEGIN" "$RC_FILE"; then
        sed -i "/$WRAPPER_BEGIN/,/$WRAPPER_END/d" "$RC_FILE"
        ok "Shell wrapper removed from $RC_FILE"
        CHANGED=1
    else
        info "No shell wrapper found in $RC_FILE"
    fi

    # remove source line (exact match)
    if grep -qF "$AMF_SOURCE_LINE" "$RC_FILE"; then
        ESCAPED=$(printf '%s\n' "$AMF_SOURCE_LINE" | sed 's/[][\/.^$*]/\\&/g')
        sed -i "/^${ESCAPED}$/d" "$RC_FILE"
        ok "Source line removed from $RC_FILE"
        CHANGED=1
    else
        info "No source line found in $RC_FILE"
    fi

    if [ "$CHANGED" -eq 0 ]; then
        rm -f "$BACKUP"
        info "No changes made to $RC_FILE (backup discarded)"
    fi
fi

# config directory (contains user aliases — ask first)
if [ -d "$CONFIG_DIR" ]; then
    echo ""
    warn "$CONFIG_DIR still exists and contains your aliases."
    if confirm "  Delete it?"; then
        rm -rf "$CONFIG_DIR" || fail "Could not remove $CONFIG_DIR"
        ok "Removed $CONFIG_DIR"
    else
        info "Kept $CONFIG_DIR"
    fi
fi

echo ""
echo "  Done. Reload your shell to finish:"
echo ""
if [ -n "$RC_FILE" ] && [ -f "$RC_FILE" ]; then
    echo "    source $RC_FILE"
else
    echo "    (restart your shell)"
fi
echo ""
