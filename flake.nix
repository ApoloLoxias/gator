{

  inputs = {nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";};

outputs = {self, nixpkgs, ...}:
let
        pkgs = nixpkgs.legacyPackages.x86_64-linux;
in {
        devShells.x86_64-linux.default = pkgs.mkShell {
                packages = [pkgs.go_1_27 pkgs.bootdev-cli];
                shellHook = "go version";
        };
};

}
