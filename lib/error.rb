class DockerVlmError < StandardError; end

class DockerComposeError < DockerVlmError; end
class DockerExportError < DockerVlmError; end
