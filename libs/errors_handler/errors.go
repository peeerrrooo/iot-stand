package errors_handler

var Handler map[int]string = map[int]string{
	// Common errors.
	10: "Internal error",
	11: "Error in JSON parse",
	12: "Error in JSON ARRAY parse",
	13: "Error in JSON encode",
	14: "Error in JSON ARRAY encode",
	15: "Cannot prepare phone number for standard",
	16: "Phone is not valid",

	// Game-service errors.
	1000: "Bad json format for input message",
	1001: "Cannot find field 'method'",
	1002: "Cannot find field 'req_id'",
	1003: "Unknown method",
	1004: "Example test error code",
	1005: "Miss param",
}
