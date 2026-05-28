# Homebrew Tap Setup

Date: 2026-05-28

Scope: Exact manual setup path for the intended external Homebrew tap. This is
distribution infrastructure only; it does not add `ni-kernel` runtime behavior.

Current Homebrew status: Planned.

Evidence: `git ls-remote https://github.com/Nam-Cheol/homebrew-tap.git`
returned "Repository not found" on 2026-05-28. Because the tap repository does
not exist, `brew install Nam-Cheol/tap/ni` has not been tested and Homebrew must
not be marked Available.

## Tap Decision

Use this tap repository:

```text
Nam-Cheol/homebrew-tap
```

Use this formula name:

```text
ni
```

Users will install from the tap only after it exists and is verified:

```bash
brew install Nam-Cheol/tap/ni
```

Do not add this command to public install instructions until the command has
actually passed.

## Manual Setup Steps

These steps must be run by the repository owner or someone with permission to
create repositories under `Nam-Cheol`.

1. Create a public GitHub repository named `homebrew-tap` under `Nam-Cheol`.
   The repository URL must be:

   ```text
   https://github.com/Nam-Cheol/homebrew-tap
   ```

2. Clone the new tap repository:

   ```bash
   git clone https://github.com/Nam-Cheol/homebrew-tap.git
   cd homebrew-tap
   ```

3. Create the Homebrew formula directory:

   ```bash
   mkdir -p Formula
   ```

4. Add `Formula/ni.rb` with the formula below.

5. Commit and push the formula:

   ```bash
   git add Formula/ni.rb
   git commit -m "Add ni formula"
   git push origin main
   ```

6. Validate from a machine with Homebrew installed:

   ```bash
   brew update
   brew audit --strict --online Formula/ni.rb
   brew install --build-from-source Formula/ni.rb
   ni --help
   ni version
   ```

7. Test the published tap path:

   ```bash
   brew uninstall ni
   brew untap Nam-Cheol/tap || true
   brew install Nam-Cheol/tap/ni
   ni --help
   ni version
   ```

8. Only after the published tap command works, update README and install docs
   from Planned to Available with the tested command output as evidence.

## Formula Sources

The first formula should point at the already verified v0.3.0 GitHub Release
archives. The checksums below come from
[v0.3.0 Release Verification](70_RELEASE_VERIFICATION_v0.3.0.md).

| Platform | Source URL | sha256 |
| --- | --- | --- |
| macOS amd64 | `https://github.com/Nam-Cheol/ni/releases/download/v0.3.0/ni_0.3.0_darwin_amd64.tar.gz` | `b6d65b177f0a58e7c9457fc562494e8d6dfdc92655aa0b1bb4aa697a8da952e0` |
| macOS arm64 | `https://github.com/Nam-Cheol/ni/releases/download/v0.3.0/ni_0.3.0_darwin_arm64.tar.gz` | `a41a45afb0e1f11779b28d70f397430773d7ad5f23252771077cc8fafefe0f33` |
| Linux amd64 | `https://github.com/Nam-Cheol/ni/releases/download/v0.3.0/ni_0.3.0_linux_amd64.tar.gz` | `7032a70dbe8e3824b10c6fa83e315507d8d135c89fe1cf0cc1597ebab19896e9` |
| Linux arm64 | `https://github.com/Nam-Cheol/ni/releases/download/v0.3.0/ni_0.3.0_linux_arm64.tar.gz` | `e7401a78465f2401c1948a05c2a4c646dfc9e6f0be834e8f0b888a466e3b20f9` |

## Formula/ni.rb

```ruby
class Ni < Formula
  desc "Project Intent Compiler for AI Agents"
  homepage "https://github.com/Nam-Cheol/ni"
  version "0.3.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/Nam-Cheol/ni/releases/download/v#{version}/ni_#{version}_darwin_arm64.tar.gz"
      sha256 "a41a45afb0e1f11779b28d70f397430773d7ad5f23252771077cc8fafefe0f33"
    end

    if Hardware::CPU.intel?
      url "https://github.com/Nam-Cheol/ni/releases/download/v#{version}/ni_#{version}_darwin_amd64.tar.gz"
      sha256 "b6d65b177f0a58e7c9457fc562494e8d6dfdc92655aa0b1bb4aa697a8da952e0"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/Nam-Cheol/ni/releases/download/v#{version}/ni_#{version}_linux_arm64.tar.gz"
      sha256 "e7401a78465f2401c1948a05c2a4c646dfc9e6f0be834e8f0b888a466e3b20f9"
    end

    if Hardware::CPU.intel?
      url "https://github.com/Nam-Cheol/ni/releases/download/v#{version}/ni_#{version}_linux_amd64.tar.gz"
      sha256 "7032a70dbe8e3824b10c6fa83e315507d8d135c89fe1cf0cc1597ebab19896e9"
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

## Availability Gate

Homebrew remains Planned until all of these are true:

1. `Nam-Cheol/homebrew-tap` exists as the intended public tap.
2. `Formula/ni.rb` is published in that tap.
3. Homebrew audit passes.
4. A local formula install passes.
5. The published tap command passes:

   ```bash
   brew install Nam-Cheol/tap/ni
   ni --help
   ni version
   ```

If any step has not passed, keep Homebrew Planned or at most Experimental. Do
not claim package-manager availability from formula text alone.
