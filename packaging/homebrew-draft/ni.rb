# Draft-only formula stored for audit and local review.
# This file is not published in a Homebrew tap and does not make Homebrew Available.

class Ni < Formula
  desc "Project Intent Compiler for AI Agents"
  homepage "https://github.com/Nam-Cheol/ni"
  url "https://github.com/Nam-Cheol/ni/archive/refs/tags/v0.5.0.tar.gz"
  sha256 "67a694ff9e9e076b2cfc731c96575604e18abea03b1bb1f818e95b9aee54bb02"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", "-trimpath",
           "-ldflags", "-s -w -X ni/internal/version.Version=#{version}",
           "-o", bin/"ni", "./cmd/ni"
  end

  test do
    assert_match "ni is a project intent compiler", shell_output("#{bin}/ni --help")
    assert_match version.to_s, shell_output("#{bin}/ni version")
  end
end
