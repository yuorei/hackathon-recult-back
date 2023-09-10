require'mysql2'
class DB
    def initialize()
        @@db_conn = Mysql2::Client.new(host: "localhost", username: "root", password: '', database: 'USER_INFORMATION')
    end

    def fetch_user_from_database_by_id(id)
        begin
            result=@@db_conn.query("SELECT * FROM USERS WHERE id = '#{id}'")
        rescue => e
            puts "エラーが発生しました: #{e.message}"
            return e.message
          end
        return result
    end

    def fetch_user_from_database_by_email(email)
        begin
            result=@@db_conn.query("SELECT * FROM USERS WHERE email = '#{email}'")
        rescue => e
            puts "エラーが発生しました: #{e.message}"
            return e.message
          end
        return result
    end

    def create_user_in_database(request)
        begin
            @@db_conn.query("INSERT INTO USERS(name, email, hashed_password, gender, affiliation) VALUES ('#{request.name}', '#{request.email}', '#{request.password}', '#{request.gender}', '#{request.affiliation}')")
          rescue => e
            puts "エラーが発生しました：#{e.message}"
            return e.message
          end
    end
end