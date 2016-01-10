#!/usr/bin/env ruby

require 'test/unit'
require_relative '../lib/root'
require_relative '../lib/default_config'

class ExportTest < Test::Unit::TestCase
  def test_export
    docker_compose_path = File.join(File.dirname(__FILE__), './docker-compose.yml')
    config_path = File.join(File.dirname(__FILE__), '../config.yml')
    default_config = DefaultConfig.new(config_path, docker_compose_path)
    root('export', [], default_config)
  end
end
