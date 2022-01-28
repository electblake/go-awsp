class Awsp < Formula
    desc "Allows for easy switching between AWS profiles"
    homepage "https://github.com/electblake/homebrew-awsp"
    url "https://github.com/electblake/homebrew-awsp/archive/refs/tags/v0.0.2.tar.gz"
    version "0.0.2"
    sha256 "b867fc5cbf9e4be389373664c495074cfab612f50aa2cbd0385eb894db614e2c"
  
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