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

  vendorHash = "sha256-cHVM5BtvW4004T3wMzrIKhpofyLLCqAEkMFLVkIyo4A=";

  nativeBuildInputs = [
    pkg-config
  ];

  buildInputs = [
    opencv
  ];

  doCheck = false;
}
