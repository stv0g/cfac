# SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
# SPDX-License-Identifier: Apache-2.0
{
  description = "Code for Aachen";

  nixConfig = {
    extra-substituters = "https://stv0g.cachix.org";
    extra-trusted-public-keys = "stv0g.cachix.org-1:Bliox3TtWqQhKr2w6HMSbpwn9E9M2vgKmA/N7VpYOmY=";
  };

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
          inherit (pkgs) cfac;
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

        formatter = pkgs.nixfmt-rfc-style;
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
