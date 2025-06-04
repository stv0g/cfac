# SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
# SPDX-License-Identifier: Apache-2.0

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
