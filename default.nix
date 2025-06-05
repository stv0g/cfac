# SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
# SPDX-License-Identifier: Apache-2.0

{
  buildGo124Module,
  pkg-config,
  opencv,
}:
buildGo124Module {
  pname = "cfac";
  version = "0.1.0";

  src = ./.;

  vendorHash = "sha256-FX0gv8W5wZYelkN6gjwRaMNx7vD+QmPBI7AUWQJf5Vs=";

  nativeBuildInputs = [
    pkg-config
  ];

  buildInputs = [
    opencv
  ];

  doCheck = false;
}
