class Release < ActiveRecord::Base
    has_many :track
    belongs_to :user
end