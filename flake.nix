{
  description = "Example kickstart Go module project.";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
    devshells.url = "git://5.161.202.250/nix-dev-shells";
    devshells.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs = inputs @ {flake-parts, ...}:
    flake-parts.lib.mkFlake {inherit inputs;} {
      systems = ["x86_64-linux" "aarch64-linux" "aarch64-darwin" "x86_64-darwin"];
      perSystem = {
        config,
        self',
        inputs',
        pkgs,
        system,
        ...
      }: let
        name = "walrossweb";
        version = "latest";
        vendorHash = null; # update whenever go.mod changes
        mkGoDevShell = inputs.devshells.lib.${system}.mkGoDevShell;
      in {
        devShells.default = mkGoDevShell {
          cmd = ''
            ( cd cmd/walross && go run main.go ) & \
            ( tailwindcss --input input.css --output cmd/walross/static/styles.css --watch ) & \
            wait
          '';
          hotReload = false;
          extraPackages = with pkgs; [
            nixpkgs-fmt
            tailwindcss
          ];
        };

        packages.default = pkgs.buildGoModule {
          inherit name vendorHash;
          src = ./.;
          subPackages = ["cmd/walross"];
        };
      };
    };
}
