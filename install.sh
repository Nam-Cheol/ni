#!/usr/bin/env sh
set -eu

REPO="${NI_INSTALL_REPO:-Nam-Cheol/ni}"
VERSION="${NI_INSTALL_VERSION:-}"
TAG="${NI_INSTALL_TAG:-}"
BASE_URL="${NI_INSTALL_BASE_URL:-}"
DRY_RUN=0
UPDATE_PATH=0
UNINSTALL=0
PATH_BLOCK_BEGIN="# >>> ni installer >>>"
PATH_BLOCK_END="# <<< ni installer <<<"

if [ "${HOME:-}" = "" ] && [ "${BINDIR:-}" = "" ]; then
  echo "install.sh: HOME is not set; set BINDIR to choose an install directory" >&2
  exit 1
fi

BINDIR="${BINDIR:-$HOME/.local/bin}"

usage() {
  cat <<'EOF'
Install ni from GitHub Releases.

Usage:
  sh install.sh [--dry-run] [--update-path] [--version VERSION] [--repo OWNER/REPO]
  sh install.sh --uninstall

Options:
  --dry-run          Show the selected platform, asset, and install path.
  --update-path      Add a reversible ni-managed PATH block to zsh or bash
                     shell profile when the install directory is not on PATH.
  --uninstall        Remove the installed ni binary and the ni-managed PATH
                     block, if present.
  --version VERSION  Install a specific release version, such as 0.2.0.
                     Tags are resolved as vVERSION unless VERSION starts with v.
  --repo OWNER/REPO  GitHub repository to download from.
  -h, --help         Show this help.

Environment:
  BINDIR             Install directory. Default: $HOME/.local/bin
  NI_INSTALL_TAG     Exact release tag, overriding --version tag resolution.
  NI_INSTALL_VERSION Asset version used in file names.
  NI_INSTALL_BASE_URL
                     Base URL for release assets. Defaults to GitHub Releases.
  NI_INSTALL_OS      Override detected os for tests: linux, darwin, windows.
  NI_INSTALL_ARCH    Override detected arch for tests: amd64, arm64.
  NI_INSTALL_SHELL_PROFILE
                     Override the shell profile used by --update-path or
                     --uninstall.
EOF
}

say() {
  printf '%s\n' "$*"
}

warn() {
  printf 'install.sh: warning: %s\n' "$*" >&2
}

die() {
  printf 'install.sh: %s\n' "$*" >&2
  exit 1
}

need() {
  command -v "$1" >/dev/null 2>&1 || die "required command not found: $1"
}

while [ "$#" -gt 0 ]; do
  case "$1" in
    --dry-run)
      DRY_RUN=1
      shift
      ;;
    --update-path)
      UPDATE_PATH=1
      shift
      ;;
    --uninstall)
      UNINSTALL=1
      shift
      ;;
    --version)
      [ "$#" -ge 2 ] || die "--version requires a value"
      VERSION="$2"
      shift 2
      ;;
    --repo)
      [ "$#" -ge 2 ] || die "--repo requires OWNER/REPO"
      REPO="$2"
      shift 2
      ;;
    -h|--help)
      usage
      exit 0
      ;;
    *)
      die "unknown option: $1"
      ;;
  esac
done

path_has_bindir() {
  case ":${PATH:-}:" in
    *":$BINDIR:"*) return 0 ;;
    *) return 1 ;;
  esac
}

shell_profile() {
  if [ "${NI_INSTALL_SHELL_PROFILE:-}" != "" ]; then
    printf '%s\n' "$NI_INSTALL_SHELL_PROFILE"
    return
  fi

  case "$(basename "${SHELL:-}")" in
    zsh) printf '%s\n' "$HOME/.zshrc" ;;
    bash) printf '%s\n' "$HOME/.bashrc" ;;
    *) printf '%s\n' "" ;;
  esac
}

remove_path_block() {
  profile="$1"
  [ "$profile" != "" ] || return 0
  [ -f "$profile" ] || return 0

  tmp="${profile}.ni-tmp.$$"
  status=0
  awk -v begin="$PATH_BLOCK_BEGIN" -v end="$PATH_BLOCK_END" '
    $0 == begin { skip = 1; changed = 1; next }
    $0 == end { skip = 0; next }
    !skip { print }
    END { if (changed != 1) exit 2 }
  ' "$profile" >"$tmp" || status=$?
  status="${status:-0}"
  if [ "$status" -eq 0 ]; then
    mv "$tmp" "$profile"
    say "Removed ni PATH block from $profile"
  else
    rm -f "$tmp"
  fi
}

add_path_block() {
  profile="$1"
  [ "$profile" != "" ] || die "could not choose a zsh or bash profile; set NI_INSTALL_SHELL_PROFILE"

  mkdir -p "$(dirname "$profile")"
  touch "$profile"

  if grep -Fq "$PATH_BLOCK_BEGIN" "$profile"; then
    say "ni PATH block already exists in $profile"
    return
  fi

  path_expr="$BINDIR"
  case "$BINDIR" in
    "$HOME"/*)
      path_expr="\$HOME${BINDIR#"$HOME"}"
      ;;
  esac

  {
    printf '\n%s\n' "$PATH_BLOCK_BEGIN"
    printf 'export PATH="%s:$PATH"\n' "$path_expr"
    printf '%s\n' "$PATH_BLOCK_END"
  } >>"$profile"
  say "Added ni PATH block to $profile"
}

detect_os() {
  if [ "${NI_INSTALL_OS:-}" != "" ]; then
    printf '%s\n' "$NI_INSTALL_OS"
    return
  fi

  case "$(uname -s)" in
    Linux) printf '%s\n' linux ;;
    Darwin) printf '%s\n' darwin ;;
    MINGW*|MSYS*|CYGWIN*) printf '%s\n' windows ;;
    *) die "unsupported operating system: $(uname -s)" ;;
  esac
}

detect_arch() {
  if [ "${NI_INSTALL_ARCH:-}" != "" ]; then
    printf '%s\n' "$NI_INSTALL_ARCH"
    return
  fi

  case "$(uname -m)" in
    x86_64|amd64) printf '%s\n' amd64 ;;
    arm64|aarch64) printf '%s\n' arm64 ;;
    *) die "unsupported architecture: $(uname -m)" ;;
  esac
}

resolve_latest_tag() {
  need curl
  curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" \
    | sed -n 's/.*"tag_name"[[:space:]]*:[[:space:]]*"\([^"]*\)".*/\1/p' \
    | sed -n '1p'
}

sha256_file() {
  if command -v sha256sum >/dev/null 2>&1; then
    sha256sum "$1" | awk '{print $1}'
    return 0
  fi

  if command -v shasum >/dev/null 2>&1; then
    shasum -a 256 "$1" | awk '{print $1}'
    return 0
  fi

  return 1
}

download() {
  need curl
  curl -fsSL "$1" -o "$2"
}

OS="$(detect_os)"
ARCH="$(detect_arch)"

case "$OS/$ARCH" in
  linux/amd64|linux/arm64|darwin/amd64|darwin/arm64|windows/amd64) ;;
  windows/arm64) die "windows arm64 release asset is not configured" ;;
  *) die "unsupported platform: $OS/$ARCH" ;;
esac

EXT="tar.gz"
BIN_NAME="ni"
if [ "$OS" = "windows" ]; then
  EXT="zip"
  BIN_NAME="ni.exe"
fi

TARGET="$BINDIR/$BIN_NAME"

if [ "$UNINSTALL" -eq 1 ]; then
  profile="$(shell_profile)"
  if [ -f "$TARGET" ]; then
    rm -f "$TARGET"
    say "Removed $TARGET"
  else
    say "No installed ni binary found at $TARGET"
  fi
  rmdir "$BINDIR" 2>/dev/null || true
  remove_path_block "$profile"
  say "Uninstall complete."
  say "Open a new shell, then verify:"
  say "  command -v ni"
  say "The command should not find the ni install removed by this installer."
  exit 0
fi

if [ "$TAG" = "" ]; then
  if [ "$VERSION" != "" ]; then
    case "$VERSION" in
      v*) TAG="$VERSION" ;;
      *) TAG="v$VERSION" ;;
    esac
  elif [ "$DRY_RUN" -eq 1 ]; then
    TAG="<latest>"
  else
    TAG="$(resolve_latest_tag)"
    [ "$TAG" != "" ] || die "could not resolve latest release tag for $REPO"
  fi
fi

if [ "$VERSION" = "" ]; then
  VERSION="${TAG#v}"
else
  VERSION="${VERSION#v}"
fi

ASSET="ni_${VERSION}_${OS}_${ARCH}.${EXT}"
CHECKSUMS="ni_${VERSION}_checksums.txt"

if [ "$BASE_URL" = "" ]; then
  BASE_URL="https://github.com/$REPO/releases/download/$TAG"
fi

ASSET_URL="${BASE_URL%/}/$ASSET"
CHECKSUM_URL="${BASE_URL%/}/$CHECKSUMS"

say "ni installer"
say "  repository: $REPO"
say "  platform:   $OS/$ARCH"
say "  asset:      $ASSET"
say "  checksums:  $CHECKSUMS"
say "  install to: $TARGET"
if path_has_bindir; then
  say "  PATH:       $BINDIR is already on PATH"
else
  say "  PATH:       $BINDIR is not currently on PATH"
fi

if [ "$DRY_RUN" -eq 1 ]; then
  say "  mode:       dry-run"
  if [ "$UPDATE_PATH" -eq 1 ]; then
    profile="$(shell_profile)"
    say "  path file:  ${profile:-<unknown>}"
    say "  path mode:  would add a ni-managed PATH block if needed"
  fi
  say ""
  say "Would download:"
  say "  $ASSET_URL"
  say "  $CHECKSUM_URL"
  exit 0
fi

TMPDIR_ROOT="${TMPDIR:-/tmp}"
WORKDIR="$(mktemp -d "$TMPDIR_ROOT/ni-install.XXXXXX")"
cleanup() {
  rm -rf "$WORKDIR"
}
trap cleanup EXIT INT TERM

ARCHIVE_PATH="$WORKDIR/$ASSET"
CHECKSUM_PATH="$WORKDIR/$CHECKSUMS"
EXTRACT_DIR="$WORKDIR/extract"

say "Downloading $ASSET"
download "$ASSET_URL" "$ARCHIVE_PATH" || die "download failed: $ASSET_URL"

CHECKSUM_AVAILABLE=0
if download "$CHECKSUM_URL" "$CHECKSUM_PATH"; then
  CHECKSUM_AVAILABLE=1
else
  warn "checksum file was not available; continuing without checksum verification"
fi

if [ "$CHECKSUM_AVAILABLE" -eq 1 ]; then
  EXPECTED_SHA="$(awk -v asset="$ASSET" '
    {
      name = $2
      sub(/^\*/, "", name)
      sub(/^\.\//, "", name)
      if (name == asset) {
        print $1
        exit
      }
    }
  ' "$CHECKSUM_PATH")"

  [ "$EXPECTED_SHA" != "" ] || die "checksum file does not contain $ASSET"

  if ACTUAL_SHA="$(sha256_file "$ARCHIVE_PATH")"; then
    [ "$EXPECTED_SHA" = "$ACTUAL_SHA" ] || die "checksum mismatch for $ASSET"
    say "Verified checksum for $ASSET"
  else
    warn "no sha256 verifier found; continuing without local checksum verification"
  fi
fi

mkdir -p "$EXTRACT_DIR"
case "$EXT" in
  tar.gz)
    need tar
    tar -xzf "$ARCHIVE_PATH" -C "$EXTRACT_DIR"
    ;;
  zip)
    need unzip
    unzip -q "$ARCHIVE_PATH" -d "$EXTRACT_DIR"
    ;;
  *)
    die "unsupported archive extension: $EXT"
    ;;
esac

FOUND_BIN="$(find "$EXTRACT_DIR" -type f -name "$BIN_NAME" | sed -n '1p')"
[ "$FOUND_BIN" != "" ] || die "archive did not contain $BIN_NAME"

mkdir -p "$BINDIR"
cp "$FOUND_BIN" "$TARGET"
chmod 0755 "$TARGET"

say "Installed ni to $TARGET"

if path_has_bindir; then
  PATH_READY=1
elif [ "$UPDATE_PATH" -eq 1 ]; then
  profile="$(shell_profile)"
  remove_path_block "$profile"
  add_path_block "$profile"
  PATH_READY=0
else
  PATH_READY=0
fi

say ""
say "Next steps:"
if [ "$PATH_READY" -eq 1 ]; then
  say "  1. Open a new shell and check the global command:"
else
  say "  1. Open a new shell after adding $BINDIR to PATH."
  if [ "$UPDATE_PATH" -eq 0 ]; then
    say "     To let this installer add a reversible zsh/bash profile block, rerun with --update-path."
  fi
  say "  2. Check the global command:"
fi
say "     ni --help"
say "     ni version"
say ""
say "Uninstall:"
say "  sh install.sh --uninstall"
say ""
say "The installer does not install model skills or run downstream work."
