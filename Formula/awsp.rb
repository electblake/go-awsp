class Awsp < Formula
    desc "Allows for easy switching between AWS profiles"
    homepage "https://github.com/electblake/homebrew-awsp"
    url "https://github.com/electblake/homebrew-awsp/archive/refs/tags/v0.1.0.tar.gz"
    version "0.1.0"
    sha256 "ca9fa6fecc92823f921f6d0457e021145acce0501eaa0a6a8d7ec06d16851398"
  
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