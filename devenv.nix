{ pkgs, lib, config, inputs, ... }:

{
  languages.go.enable = true;
  languages.go.lsp.enable = true;
  packages = [
    pkgs.bluez
    pkgs.goreleaser
  ];
}
