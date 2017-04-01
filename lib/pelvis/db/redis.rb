require 'redis'
require 'json'

REDIS_HOST_NAME = ENV['REDIS_PORT_6379_TCP_ADDR'] || '127.0.0.1'
redis = Redis.new(:host => REDIS_HOST_NAME, :port => 6379)

module RedisDB
	def add_task(id, metadata)
		redis.set(id, metadata.to_json)
		redis.lpush('render_queue', id)
	end

	def get_favorites(key_name)
		JSON.parse(redis.get(key_name))
	end

	def set_favorites(key_name, favorites_list)
		redis.set(key_name, favorites_list)
	end

	def get_task(id)
		JSON.parse(redis.get(id))
	end

	def get_task_id_from_queue()
		redis.rpop('render_queue')
	end

	def add_to_queue(queue, id)
		redis.lpush(queue, id)
	end
end