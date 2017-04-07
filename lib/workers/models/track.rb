require 'active_record'

class Track < ActiveRecord::Base
    belongs_to :release
end