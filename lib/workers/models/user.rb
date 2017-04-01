class User < ActiveRecord::Base
    has_many :release
    has_many :social
end