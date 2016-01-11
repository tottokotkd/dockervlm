#!/usr/bin/env ruby

require 'test/unit'
require 'open3'

class ExportTest < Test::Unit::TestCase
  def test_export
    system("echo hoge;sleep 1;echo hoge 1>&2;echo hoge 1>&2;echo hoge; sleep 1;echo hoge 1>&2;")
  end
end
