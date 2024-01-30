# -*- mode: Python -*-

# For more on Extensions, see: https://docs.tilt.dev/extensions.html
load('ext://restart_process', 'docker_build_with_restart')

# Records the current time, then kicks off a server update.
# Normally, you would let Tilt do deploys automatically, but this
# shows you how to set up a custom workflow that measures it.
local_resource(
    'deploy',
    'date')

compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/example-admission-controller ./'
if os.name == 'nt':
  compile_cmd = 'build.bat'

local_resource(
  'example-admission-controller-compile',
  compile_cmd,
  deps=['./server.go'],
  resource_deps=['deploy'])

docker_build_with_restart(
  'example-admission-controller-image',
  '.',
  entrypoint=['/app/build/example-admission-controller'],
  dockerfile='deployments/Dockerfile',
  only=[
    './build',
  ],
  live_update=[
    sync('./build', '/app/build'),
  ],
)

k8s_yaml('deployments/kubernetes.yaml')
k8s_resource('example-admission-controller', port_forwards=8000,
             resource_deps=['deploy', 'example-admission-controller-compile'])