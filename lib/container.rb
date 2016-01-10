require 'open3'
require_relative 'error'

class Container < Struct.new(:container_name, :id, :volumes)

  def export
      o, e, s = Open3.capture3("docker-compose ps -q #{container_name}")
      raise DockerComposeError.new(e) if s != 0 or
      return o.strip
  end

  def to_s
    return "container: #{container_name}\n" +
    "  id: #{id}\n" +
    "  volume(s):\n" +
    volumes.map{|v| "  - #{v}\n"}.reduce{|a, b| a + b}.to_s
  end

  def export(options)
    targets = volumes.reduce{|a, b| a + ' ' + b}
    Containers.export(options, self)
  end

  def targets
    return volumes.reduce{|a, b| a + ' ' + b}
  end

end

module Containers

  CONTAINER_VOLUME = '/backup'

  def self.parseConfigYaml(config)
    return  config.
      select {|name, settings| settings.key?('volumes')}.
      map { |name, settings|
        Container.new(name, getContainerId(name), settings['volumes'].
        map {|volume| volume.include?(":") ? volume.split(/(?<=:)(.*)/)[1] : volume})
      }
  end

  def self.getContainerId(name)
    o, e, s = Open3.capture3("docker-compose ps -q #{name}")
    raise DockerComposeError.new(e) if s != 0
    raise DockerComposeError.new("container of #{name} not found... no running containers?") if o.strip == ''
    return o.strip
  end

  def self.export(options, container)
    tar_file = makeTimestampFileName(options.file_name_format)
    volume = makeDataVolumePath(File.join(options.destination, container.container_name))
    o, e, s = Open3.capture3("docker run --rm --volumes-from #{container.id} -v #{volume} busybox tar cvf #{tar_file} #{container.targets}")
    raise DockerExportError.new(e) if s != 0
    return o
  end

  private

  def self.makeTimestampFileName(format)
    name = DateTime.now.strftime(format)
    return File.join(CONTAINER_VOLUME, name)
  end

  def self.makeDataVolumePath(host_destination)
    return "#{host_destination}:#{CONTAINER_VOLUME}"
  end

end
