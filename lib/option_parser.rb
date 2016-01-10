require 'optparse'
require_relative './dockervlm_config'

class ExportOptions < Struct.new(:destination, :file_name_format)
  def initialize(destination, file_name_format)
   super(File.expand_path(destination), file_name_format)
  end
end

module Parser
  def self.parseForExportOptions(args, config)
    destination = config.destination
    file_name_format = config.file_name_format
    options = ExportOptions.new(destination, file_name_format)
    OptionParser.new{|opt|
      opt.on('-d VAL') {|v| destination = v }
      opt.on('-f VAL') {|v| file_name_format = v}
      opt.parse!(args)
    }
    return ExportOptions.new(destination, file_name_format)
  end
end
