require "http/server"

port = 8080
host = "0.0.0.0"
mime = "text/html"

server = HTTP::Server.new do |context|
  req = context.request
  if req.method == "GET" && req.path == "/"
    filename = "./index.html"
    context.response.content_type = "text/html"
    context.response.content_length = File.size(filename)
    File.open(filename) do |file|
      IO.copy(file, context.response)
    end
    next
  end

  context.response.content_type = mime
end

server.bind_tcp(host, 8080)
puts "Listening on http://0.0.0.0:8080"
server.listen
