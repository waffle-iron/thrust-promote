require_relative '../tools/helpers'


class AudioRender
    attr_reader :audio_file

    def initialize(audio_file, file_ext)
        @audio_file = audio_file
        basename = File.basename audio_file, File.extname(audio_file)
        @transcoded_file = basename + '.' + file_ext
    end

    def run
        puts "Running command: ffmpeg -i #{@audio_file} #{@transcoded_file}"
        cmd = `ffmpeg -i #{@audio_file} #{@transcoded_file}`
        @transcoded_file
    end

    def self.run(audio_file, file_ext)
        AudioRender.new(audio_file, file_ext).run
    end
end


