require 'yaml'

class DefaultConfig
  def initialize(config_path)
    @config = YAML.load_file(config_path)
  end
  def file_name_format
    return @config['file_name_format']
  end
  def destination
    return @config['destination']
  end
end
