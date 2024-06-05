class Freeport < Formula
    desc "A simple CLI tool to kill the processes that are using a specific port"
    homepage "https://github.com/ccc159/freeport"
    url "https://github.com/ccc159/freeport/releases/download/v0.1.0/freeport-0.1.0-darwin-amd64.tar.gz"
    sha256 "790dd329fe85d7f03c61274ad97feedeb2905af1e6dec41d4e3b2d35aada8496"
    license "MIT"
    version "0.1.0"
  
    def install
      bin.install "freeport"
    end
  
    test do
      system "#{bin}/freeport", "--version"
    end
  end
  