let
        pkgs = import <nixpkgs> {};
in
        pkgs.mkShell {
                packages = [pkgs.go_1_27];
                shellHook = "go version";
        }
