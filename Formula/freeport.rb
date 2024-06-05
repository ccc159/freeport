class Freeport < Formula
    desc "A simple CLI tool to kill the processes that are using a specific port"
    homepage "https://github.com/ccc159/freeport"
    url "https://github.com/ccc159/freeport/releases/download/v0.1.0/freeport-0.1.0-darwin-amd64.tar.gz"
    sha256 "c7124460e59f6e171aa79d614f0de6ec5d53616d2a37f063432fcdda59b25d3f"
    license "MIT"
    version "0.1.0"
  
    def install
      bin.install "freeport"
    end
  
    test do
      system "#{bin}/freeport", "--version"
    end
  end
  