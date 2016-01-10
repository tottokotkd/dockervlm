require 'optparse'
require_relative './dockervlm_config'

class ExportOptions < Struct.new(:destination, :file_name_format)
  def initialize(destination, file_name_format)
   super(File.expand_path(destination), file_name_format)
  end
end

class ImportOptions < Struct.new(:source)
  def initialize(source)
   super(File.expand_path(source))
  end
end

module Parser
  def self.parseForExportOptions(args, config)
    destination = config.destination
    file_name_format = config.file_name_format
    OptionParser.new{|opt|
      opt.on('-d VAL') {|v| destination = v }
      opt.on('-f VAL') {|v| file_name_format = v}
      opt.parse!(args)
    }
    return ExportOptions.new(destination, file_name_format)
  end

  def self.parseForImportOptions(args, config)
    source = config.destination
    OptionParser.new{|opt|
      opt.on('-s VAL') {|v| source = v }
      opt.parse!(args)
    }
    return ImportOptions.new(source)
  end

end
