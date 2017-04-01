require 'rethinkdb'
include RethinkDB::Shortcuts


HOST_NAME = ENV['DB_PORT_28015_TCP_ADDR'] || '127.0.0.1'

module ReDB

  R = r 

  begin
    connection = R.connect(:host => HOST_NAME, :port => 28015)
  rescue Exception => err
    puts "Cannot connect to RethinkDB database (#{err.message})"
    Process.exit(1)
  end

  begin
    R.db_create('pelvis').run(connection)
  rescue RethinkDB::RqlRuntimeError => err
    puts "Database `pelvis` already exists."
  end

  # Setup tables
  begin
    R.db('pelvis').table_create('videos').run(connection)
  rescue RethinkDB::RqlRuntimeError => err
    puts "Table `videos` already exists."
  end

  begin
    R.db('pelvis').table_create('images').run(connection)
  rescue RethinkDB::RqlRuntimeError => err
    puts "Table `images` already exists."
  end

  begin
    R.db('pelvis').table_create('users').run(connection)
  rescue RethinkDB::RqlRuntimeError => err
    puts "Table `users` already exists."
  ensure
    connection.close
  end

  def ReDB.get_conn
    @rdb_connection = R.connect(:host => HOST_NAME, :port =>
        28015, :db => 'pelvis')
  end

  def close_conn
    @rdb_connection.close if @rdb_connection
  end

  def ReDB.get_table(table)
    R.db('pelvis').table(table)
  end

  def ReDB.empty_favorites(user)
    conn = ReDB.get_conn
    users = ReDB.get_table("users")
      .filter({"name" => "#{user}"}).run(conn).to_a
    unless users.nil? || users.length == 0
      ReDB.get_table("users")
        .filter({"name" => "#{user}"})
        .update({"favorites" => []})
        .run(conn)
    else
      ReDB.get_table("users")
        .insert({:name => "#{user}", :favorites => []})
        .run(conn)
    end
  end

  def ReDB.store_video(video_data)
    conn = ReDB.get_conn
    ReDB.get_table("videos").insert(video_data).run(conn)
  end

end