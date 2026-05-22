#!/usr/bin/env bash

set -e

BINARY_URL="https://github.com/Edwar-Sanz/alias-manager/releases/download/v1.0.0/am-linux-amd64"
INSTALL_DIR="$HOME/.local/bin"
AMF_SOURCE_LINE='[ -f ~/.config/am/.amf ] && source ~/.config/am/.amf'
WRAPPER='
# am — alias manager shell integration
am() {
    command am "$@"
    local code=$?
    [ $code -eq 0 ] && source ~/.config/am/.amf
    return $code
}
# end am — alias manager shell integration'

GREEN='\033[0;32m'
RED='\033[0;31m'
GRAY='\033[0;90m'
RESET='\033[0m'

ok()   { echo -e "${GREEN}✓${RESET} $1"; }
fail() { echo -e "${RED}✗${RESET} $1"; exit 1; }
info() { echo -e "${GRAY}→${RESET} $1"; }

detect_rc() {
    case "$(basename "$SHELL")" in
        zsh)  echo "$HOME/.zshrc" ;;
        bash) echo "$HOME/.bashrc" ;;
        *)    echo "" ;;
    esac
}

echo ""
echo "  Installing am — Alias Manager"
echo ""

# download binary
info "Downloading binary..."
mkdir -p "$INSTALL_DIR"
curl -fsSL "$BINARY_URL" -o "$INSTALL_DIR/am" || fail "Download failed. Check your internet connection."
chmod +x "$INSTALL_DIR/am"
ok "Binary installed at $INSTALL_DIR/am"

# check PATH
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo ""
    echo -e "  ${RED}Warning:${RESET} $INSTALL_DIR is not in your PATH."
    echo "  Add this to your rc file:"
    echo ""
    echo "    export PATH=\"\$HOME/.local/bin:\$PATH\""
    echo ""
fi

# detect rc file
RC_FILE=$(detect_rc)
if [ -z "$RC_FILE" ]; then
    echo ""
    echo "  Shell not recognized. Add manually to your rc file:"
    echo ""
    echo "$WRAPPER"
    echo ""
    echo "  $AMF_SOURCE_LINE"
    exit 0
fi

# add source line
if grep -qF "$AMF_SOURCE_LINE" "$RC_FILE" 2>/dev/null; then
    ok "Source line already present in $RC_FILE"
else
    echo "" >> "$RC_FILE"
    echo "$AMF_SOURCE_LINE" >> "$RC_FILE"
    ok "Source line added to $RC_FILE"
fi

# add shell wrapper
if grep -q "am — alias manager shell integration" "$RC_FILE" 2>/dev/null; then
    ok "Shell wrapper already present in $RC_FILE"
else
    echo "$WRAPPER" >> "$RC_FILE"
    ok "Shell wrapper added to $RC_FILE"
fi

echo ""
echo "  Done. Reload your shell to start using am:"
echo ""
echo "    source $RC_FILE"
echo ""
