{
  description = "Go tools for some-project";

  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
    self,
    nixpkgs,
    flake-utils,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        version = "0.0.1";
        mkApp =
          day:
          {
            program = "${self.packages.${system}.default}/bin/${day}";
          }
          // {
            type = "app";
          };
      in
        {
        packages = {
          default = pkgs.buildGoModule {
            pname = "advent-of-code-2024";
            inherit version;
            src = ./.;
            pwd = ./.;

            vendorHash = null;
            subPackages = [
              "day1"
              # "day2"
            ];
          };
        };
        apps = {
          day1 = mkApp "day1";
          # day2 = mkApp "day2";
        };

        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gofumpt
            gotools
            reftools
            golines
            gomodifytags
            gotests
            iferr
            impl
            delve
            ginkgo
            gotestsum
            govulncheck
          ];
        };
      }
    );
}
