class Freeport < Formula
    desc "A simple CLI tool to kill the processes that are using a specific port"
    homepage "https://github.com/ccc159/freeport"
    url "https://github.com/ccc159/freeport/releases/download/v1.0.0/freeport-1.0.0-darwin-amd64.tar.gz"
    sha256 "YOUR_SHA256_SUM_FOR_THIS_VERSION"
    license "MIT"
    version "1.0.0"
  
    def install
      bin.install "freeport"
    end
  
    test do
      system "#{bin}/freeport", "--version"
    end
  end
  