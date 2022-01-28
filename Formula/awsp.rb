class Awsp < Formula
    desc "Allows for easy switching between AWS profiles"
    homepage "https://github.com/electblake/homebrew-awsp"
    url "https://github.com/electblake/homebrew-awsp/archive/refs/tags/v0.1.4.tar.gz"
    version "0.1.4"
    sha256 "bb9879069efd71126249921dc13af0319c23b1f3471bb38574c506042a81a777"
  
    def install
      bin.install Dir["bin/_awsp"]
      bin.install Dir["bin/_awsp_prompt"]
    end
  
    def caveats; <<-EOS
      Add the following to your bash_profile or zshrc to complete the install:
  
        alias awsp=". #{HOMEBREW_PREFIX}/bin/_awsp"
  
      and source the file to pick up the change.
  
      that's it awsp with take it from there!
      You can now switch AWS profiles simply by typing
  
        awsp
  
      EOS
    end
  end