require_relative '../../lib/user_services_pb'
require_relative '../../db/user'

class UserServiceImpl < User::UserService::Service

    def initialize()
      @@db = DB.new()
    end

    # TODO: 未完成なので、実装する
    def get_user(request, _call)
      # ここに実際のユーザー取得ロジックを実装する
      user = fetch_user_from_database(request.id)

      response = User::GetUserResponse.new(
        id: user.id,
        name: user.name,
        mail: user.mail,
        gender: user.gender,
        affiliation: user.affiliation,
        # TODO: 対応させる
        # groups: user.groups.map { |group| User::Group.new(id: group.id, name: group.name, description: group.description) },
        # skills: user.skills.map { |skill| User::Skill.new(id: skill.id, name: skill.name, description: skill.description, level: skill.level) }
      )

      return response
    end

    def create_user(request, _call)
      err_create_user = @@db.create_user_in_database(request)
      user =  @@db.fetch_user_from_database_by_email(request.email)

      user_id = nil
      user_name = nil
      user_email = nil
      user_gender = nil
      user_affiliation = nil

      user.each do |row|
        user_id = row['id']
        user_name = row['name']
        user_email = row['email']
        user_hashed_password = row['hashed_password']
        user_gender = row['gender']
        user_affiliation = row['affiliation']
      end

      if err_create_user != nil
        user_id = 0
        user_email = err_create_user
      end

      response = User::CreateUserResponse.new(
        id: user_id,
        name: user_name,
        email: user_email,
        gender: user_gender,
        affiliation: user_affiliation,
        # groups: user.groups.map { |group| User::Group.new(id: group.id, name: group.name, description: group.description) },
        # skills: user.skills.map { |skill| User::Skill.new(id: skill.id, name: skill.name, description: skill.description, level: skill.level) }
      )

      return response
    end
end
