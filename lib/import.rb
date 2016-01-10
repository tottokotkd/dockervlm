require_relative './container'
require_relative './option_parser'

def import(args, config)

  puts ''
  puts "* \e[35mdockervlm\e[0m: \e[34mimport\e[0m"
  puts ''

  # parse args
  options = Parser.parseForImportOptions(args, config)

  # parse & find out volumes
  begin
    containers = Containers.parseConfigYaml(config.docker_compose)
    containers.each.with_index {|c, i| puts "\n\e[36mtarget #{i + 1}\e[0m"; puts c}
  rescue => e
    $stderr.puts e
    exit 1
  end

  # export volumes
  begin
    containers.each {|c|
      prs = c.import(options)
    }
  rescue => e
    puts "error: #{e}"
  end
end
