def showHelp
  puts <<EOS

dockervlm:
  back up tool of docker data volume (for docker-compose)

usage:
  1. cd /path/of/target/docker-compose.yml/
  2. dockervlm mode [options]

mode:
* export
  [-f filename_format_%Y%m%d-%H%M%S.tar]
  [-d ./export/destination/path]
* import

EOS
end
