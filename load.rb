require 'open-uri'

request_uri = 'http://api-dev.politician.world/p'
request_query = ''
url = "#{request_uri}#{request_query}"
for i in 1..100 do
	buffer = open(url).read
	puts i
end
