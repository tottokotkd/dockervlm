require_relative './help'
require_relative './export'

def root(mode, args, default_config)
  case mode
  when 'export'
    export(args, default_config)
  when 'import'
    import(args, default_config)
  when '--help'
    showHelp
  else
    showHelp
  end
end
