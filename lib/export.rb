require 'yaml'
require_relative './container'
require_relative './option_parser'

def export(args, default_config)

  puts ''
  puts "* \e[35mdockervlm\e[0m: \e[32mexport\e[0m"
  puts ''

  # parse args
  options = Parser.parseForExportOptions(args, default_config)

  # read docker-compose.yml
  begin
    config = YAML.load_file('docker-compose.yml')
    puts 'docker-compose.yml read.'
  rescue
    $stderr.puts 'error: docker-compose.yml not found; exit.'
    exit 1
  end

  # parse & find out volumes
  begin
    containers = Containers.parseConfigYaml(config)
    containers.each.with_index {|c, i| puts "\n\e[36mtarget #{i + 1}\e[0m"; puts c}
  rescue => e
    $stderr.puts e
    exit 1
  end

  # export volumes
  begin
    containers.each {|c|
      prs = c.export(options)
    }
  rescue => e
    puts "error: #{e}"
  end
end
