package config

const (
	noRoutesConfigBlob = `
	debug = true
	listen = "0.0.0.0:8080"
	write_timeout = "15s"
	read_timeout = "15s"
	idle_timeout = "2m"
	`

	genericRouteConfigBlob = `
	debug = true
	listen = "0.0.0.0:8080"
	write_timeout = "15s"
	read_timeout = "15s"
	idle_timeout = "2m"

	[[routes]]
	type = "generic"
	path = '/'
	status_code = 200
	headers = '{"Content-Type":"text/plain"}'
	body = 'oi'
	`

	echoRouteConfigBlob = `
	debug = true
	listen = "0.0.0.0:8080"
	write_timeout = "15s"
	read_timeout = "15s"
	idle_timeout = "2m"

	[[routes]]
	type = "echo"
	path = '/echo'
	status_code = 200
	headers = '{"Content-Type":"text/plain"}'
	`

	multipleRoutesConfigBlob = `
	debug = true
	listen = "0.0.0.0:8080"
	write_timeout = "15s"
	read_timeout = "15s"
	idle_timeout = "2m"
	
	[[routes]]
	type = "generic"
	path = '/'
	status_code = 200
	headers = '{"Content-Type":"text/plain"}'
	body = 'oi'
	
	[[routes]]
	type = "generic"
	path = '/health'
	status_code = 200
	headers = '{"Content-Type":"application/json"}'
	body = '{"status":"ok"}'
	
	[[routes]]
	type = "generic"
	path = '/panic'
	status_code = 500
	headers = '{"Content-Type":"application/json"}'
	body = '{"error":"database out"}'
	
	[[routes]]
	path = '/legacy'
	status_code = 200
	headers = '{"Content-Type":"application/xml"}'
	body = '<xml><status>error</status></xml>'
	
	[[routes]]
	type = "echo"
	path = '/echo'
	status_code = 200
	headers = '{"Content-Type":"text/plain"}'
	`
)
