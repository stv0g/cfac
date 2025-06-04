# SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
# SPDX-License-Identifier: Apache-2.0
{
  description = "Code for Aachen";

  inputs = {
    flake-utils.url = "github:numtide/flake-utils";
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
  };

  outputs =
    {
      self,
      flake-utils,
      nixpkgs,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [
            self.overlays.cfac
          ];
        };
      in
      {
        packages = {
          default = pkgs.cfac;
        };

        devShell = pkgs.mkShell {
          packages = with pkgs; [
            golangci-lint
            reuse
          ];

          inputsFrom = with pkgs; [
            cfac
          ];
        };

        formatter = nixpkgs.nixfmt-rfc-style;
      }
    )
    // {
      overlays = {
        default = self.overlays.cfac;
        cfac = final: prev: {
          cfac = final.callPackage ./default.nix { };
        };
      };
    };
}
