{
  buildGo124Module,
  pkg-config,
  opencv,
}:
buildGo124Module {
  name = "cfac";
  src = ./.;
  vendorHash = "sha256-TBIq7amF2pf1lWsGSXwxcCmj36mhfDs7tbB/pnyJ52s=";

  nativeBuildInputs = [
    pkg-config
  ];

  buildInputs = [
    opencv
  ];

  doCheck = false;
}
