{
  description = "Development environment";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
          config.allowUnfree = true;
          config = {
            permittedInsecurePackages = [
            ];
          };
        };
        formatter = pkgs.nixpkgs-fmt;
        # Check in https://search.nixos.org/packages to find new packages.
        # Use `nix --extra-experimental-features nix-command --extra-experimental-features flakes flake update`
        # to update the lock file if packages are out-of-date.

        # From https://nixos.wiki/wiki/Google_Cloud_SDK


        devShellPackages = with pkgs; [
          antlr
        ];
      in
      {
        defaultPackage = formatter; # or replace it with your desired default package.
        devShell = pkgs.mkShell { buildInputs = devShellPackages; };
      }
    );
}