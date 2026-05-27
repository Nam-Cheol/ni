#!/usr/bin/env sh
set -eu

REPO="${NI_INSTALL_REPO:-Nam-Cheol/ni}"
VERSION="${NI_INSTALL_VERSION:-}"
TAG="${NI_INSTALL_TAG:-}"
BASE_URL="${NI_INSTALL_BASE_URL:-}"
DRY_RUN=0

if [ "${HOME:-}" = "" ] && [ "${BINDIR:-}" = "" ]; then
  echo "install.sh: HOME is not set; set BINDIR to choose an install directory" >&2
  exit 1
fi

BINDIR="${BINDIR:-$HOME/.local/bin}"

usage() {
  cat <<'EOF'
Install ni from GitHub Releases.

Usage:
  sh install.sh [--dry-run] [--version VERSION] [--repo OWNER/REPO]

Options:
  --dry-run          Show the selected platform, asset, and install path.
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
TARGET="$BINDIR/$BIN_NAME"

say "ni installer"
say "  repository: $REPO"
say "  platform:   $OS/$ARCH"
say "  asset:      $ASSET"
say "  checksums:  $CHECKSUMS"
say "  install to: $TARGET"

if [ "$DRY_RUN" -eq 1 ]; then
  say "  mode:       dry-run"
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
say ""
say "Next steps:"
say "  1. Ensure $BINDIR is on your PATH."
say "  2. Check the installed CLI:"
say "     $TARGET --help"
say "     $TARGET version"
say ""
say "The installer does not install model skills or run downstream work."
