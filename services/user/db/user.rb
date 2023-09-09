require'mysql2'
class DB
    def initialize()
        @@db_conn = Mysql2::Client.new(host: "localhost", username: "root", password: '', database: 'mysql')
    end
end