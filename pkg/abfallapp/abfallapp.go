// SPDX-FileCopyrightText: 2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package abfallapp

const (
	UrlApi = "https://aachen-abfallapp.regioit.de/abfall-app-aachen/rest"

	UrlApiStandorte     = UrlApi + "/standorte?ort={ort_name}&ortsteil={ortsteil}&standortart={standortart_name}"
	UrlApiStandortArten = UrlApi + "/standorte/standortarten?ort={ort_name}"
	UrlApiOrte          = UrlApi + "/orte"
	UrlApiFraktionen    = UrlApi + "/fraktionen"
	UrlApiStrassen      = UrlApi + "/orte/{ort_id}/strassen"
	UrlApiOrtsteile     = UrlApi + "/standorte/standortarten/{standortart_name}/ortsteile?ort={ort_name}" // ?kat={kategorie_id} von stoffe
	UrlApiStoffe        = UrlApi + "/stoffe?ort={ort_name}"
)
