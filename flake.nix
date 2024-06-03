{
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  outputs = { self, nixpkgs }:
  let 
    system = "x86_64-linux";
    pkgs = import nixpkgs { system = system; config.allowUnfree = true; };
  in
  {
    devShells.${system}.default = pkgs.mkShell {
      buildInputs = with pkgs; [ 
        go 
        gopls 
        air
        wayland 
        libxkbcommon 
        libGL 
        xorg.libX11 
        xorg.libXcursor 
        xorg.libXrandr 
        xorg.libXinerama 
        xorg.libXi.dev 
      ];
    }; 
  };
}
