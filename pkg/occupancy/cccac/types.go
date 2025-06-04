// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package cccac

type ResponseCurrentStatus struct {
	Changed Status `json:"changed"`
}

type Status struct {
	Status string `json:"status"`
	Time   uint   `json:"time"`
	Type   string `json:"type"`
}
