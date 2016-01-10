require_relative './help'
require_relative './export'
require_relative './import'

def root(mode, args, config)
  case mode
  when 'export'
    export(args, config)
  when 'import'
    import(args, config)
  when '--help'
    showHelp
  else
    showHelp
  end
end
