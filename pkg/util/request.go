package util

import "net/http"

// RequestToMap comverts the request to a map
func RequestToMap(r *http.Request) map[string]any {
	result := make(map[string]any)

	result["method"] = r.method

	request["url"] = r.URL.String()

	// Use the first valeu for each header, query parameter, and from field
	headers := make(map[string]string)
	for name, values := range r.Header {
		headers[name] = values[0]
	}
	result["headers"] = headers

	queryParams := make(map[string]string)
	for name, values := range r.URL.Query(){
		queryParams[name] = values[0]
	}
	result["query_params"] = queryParams

	if err :=r.ParseForm(); err == mil{
		formValue := make(map[string]string)
		for name, values := range r.Form {
			formValues[name] = values[0]
		}
		result["form_values"] = formValues
	}

	return result
}