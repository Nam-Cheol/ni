# Homebrew Formula Draft

This is a doc-only draft for a future external tap formula. It is not a
published Homebrew formula, not a repository-local install path, and not a claim
that `brew install` works today.

Current Homebrew status: Planned.

The intended tap path is `Nam-Cheol/homebrew-tap`, but that repository must
exist and be owner-confirmed before this draft is copied into `Formula/ni.rb`.

## Draft Formula

Replace every `REPLACE_WITH_*_SHA256` value with the checksum for the matching
archive from the published GitHub Release before validation.

```ruby
class Ni < Formula
  desc "Project Intent Compiler for AI Agents"
  homepage "https://github.com/Nam-Cheol/ni"
  version "0.3.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/Nam-Cheol/ni/releases/download/v#{version}/ni_#{version}_darwin_arm64.tar.gz"
      sha256 "REPLACE_WITH_DARWIN_ARM64_SHA256"
    end

    if Hardware::CPU.intel?
      url "https://github.com/Nam-Cheol/ni/releases/download/v#{version}/ni_#{version}_darwin_amd64.tar.gz"
      sha256 "REPLACE_WITH_DARWIN_AMD64_SHA256"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/Nam-Cheol/ni/releases/download/v#{version}/ni_#{version}_linux_arm64.tar.gz"
      sha256 "REPLACE_WITH_LINUX_ARM64_SHA256"
    end

    if Hardware::CPU.intel?
      url "https://github.com/Nam-Cheol/ni/releases/download/v#{version}/ni_#{version}_linux_amd64.tar.gz"
      sha256 "REPLACE_WITH_LINUX_AMD64_SHA256"
    end
  end

  def install
    bin.install "ni"
  end

  test do
    system "#{bin}/ni", "--help"
    system "#{bin}/ni", "version"
  end
end
```

## Validation

After the tap exists and the draft has real checksums, validate from the tap
checkout:

```bash
brew audit --strict --online Formula/ni.rb
brew install --build-from-source Formula/ni.rb
ni --help
ni version
```

Do not publish automatically from `ni-kernel`. If GoReleaser later manages the
tap, that configuration should be added only after the tap exists, the owner
confirms the approach, and `goreleaser check` passes.
