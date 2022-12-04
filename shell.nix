{ pkgs ? import <nixpkgs> {} }:
   pkgs.mkShell {
      # Environment variables
      GIT_AUTHOR_EMAIL = "jwblair@gmail.com";
      GIT_AUTHOR_NAME = "Joshua Blair";
      GIT_COMMITTER_NAME = "Joshua Blair";
      GIT_COMMITTER_EMAIL = "jwblair@gmail.com";
      # Development
      nativeBuildInputs = with pkgs; [
	 delve
         go
         gotools
         gopls
         go-outline
         gocode
         gopkgs
         gocode-gomod
         godef
         golint
      ];
   }
