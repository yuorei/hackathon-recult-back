require 'grpc'
require_relative '../service/user'

def main
    server = GRPC::RpcServer.new
    server.add_http2_port('0.0.0.0:50051', :this_port_is_insecure)
    server.handle(UserServiceImpl.new)
    server.run_till_terminated
end

main
