require 'yaml'
require_relative './container'
require_relative './option_parser'

def export(args, config)

  puts ''
  puts "* \e[35mdockervlm\e[0m: \e[32mexport\e[0m"
  puts ''

  # parse args
  options = Parser.parseForExportOptions(args, config)

  # parse & find out volumes
  begin
    puts ''
    puts 'export targets ------'
    containers = Containers.parseConfigYaml(config.docker_compose)
    containers.each.with_index {|c, i| puts "\n\e[36mtarget #{i + 1}\e[0m"; puts c}
  rescue => e
    $stderr.puts e
    exit 1
  end

  # export volumes
  begin
    containers.each {|c|
      puts ''
      puts "exporting #{c.container_name}"
      prs = c.export(options)
    }
  rescue => e
    $stderr.puts "error: #{e}"
  end
end
