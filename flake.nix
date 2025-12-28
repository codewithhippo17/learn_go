{
  description = "Go project dev flake";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-24.05";
  };

  outputs = { self, nixpkgs }: let
      systems = [ "x86_64-linux" "aarch64-linux" ];
      forAllSystems = f: nixpkgs.lib.genAttrs systems (system: f system);
    in {
      devShells = forAllSystems (system:
        let pkgs = import nixpkgs { inherit system; };
        in {
          default = pkgs.mkShell {
            name = "go-dev-shell";
            packages = with pkgs; [
              go_1_22 gopls go-tools golangci-lint delve
            ];
            shellHook = ''
              echo "Go dev shell ready"
            '';
          };
        }
      );
    };
}
