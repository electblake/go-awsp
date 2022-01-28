class Awsp < Formula
    desc "Allows for easy switching between AWS profiles"
    homepage "https://github.com/electblake/go-awsp"
    url "https://github.com/electblake/go-awsp/archive/refs/tags/v0.0.1.tar.gz"
    version "0.0.1"
    sha256 "e2fc4ce2bc4ea6e551278c34d598125c67573fb1751f5c201817315c3c21df94"
  
    # depends_on "jq"
    # depends_on "grep"
  
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