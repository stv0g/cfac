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

  vendorHash = "sha256-/kGXDoAaSRLu/MPAAeJ5+h7OXzH6s/LuZ2dscxxZDYo=";

  nativeBuildInputs = [
    pkg-config
  ];

  buildInputs = [
    opencv
  ];

  doCheck = false;
}
