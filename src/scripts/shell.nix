{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = with pkgs; [
    stdenv.cc.cc.lib
    gnumake
    poetry
  ];

  # libstdc++.so.6 happens to be here; needed for locust
  LD_LIBRARY_PATH = "/nix/var/nix/profiles/per-user/human/profile/lib/julia";
  # I will probably switch back to Arch this summer; fed up with nixos
}
