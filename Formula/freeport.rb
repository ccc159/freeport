class Freeport < Formula
    desc "A simple CLI tool to kill the processes that are using a specific port"
    homepage "https://github.com/ccc159/freeport"
    url "https://github.com/ccc159/freeport/releases/download/v0.1.0/freeport-0.1.0-darwin-amd64.tar.gz"
    sha256 "8abad89d03a997ef3dfc911f413f1652995ae894294b51e4aeafca4e65796dc9"
    license "MIT"
    version "0.1.0"
  
    def install
      bin.install "freeport"
    end
  
    test do
      system "#{bin}/freeport", "--version"
    end
  end
  