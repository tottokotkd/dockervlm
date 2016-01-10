require 'yaml'

class DockervlmConfig
  def initialize(default_config_path, docker_compose_path)

    # read config
    begin
      @config = YAML.load_file(default_config_path)
    rescue
      $stderr.puts 'error: config file not found; exit.'
      exit 1
    end

    # read docker-compose.yml
    begin
      @docker_compose = YAML.load_file(docker_compose_path)
      puts 'docker-compose.yml read.'
    rescue
      $stderr.puts 'error: docker-compose.yml not found; exit.'
      exit 1
    end
  end

  def docker_compose
    return @docker_compose
  end

  def file_name_format
    return @config['file_name_format']
  end

  def destination
    return @config['destination']
  end

end
